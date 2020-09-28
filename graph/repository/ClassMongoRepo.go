package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"github.com/neilnmartin/dnd5e-graphql-api/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SubClassMongo describes the subclass of the Class
type SubClassMongo struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description []string      `json:"desc" bson:"desc"`
	Flavor      string        `json:"subclass_flavor" bson:"subclass_flavor"`
}

//ProficiencyChoiceMongo describe the number and choices for Class skill proficiencies
type ProficiencyChoiceMongo struct {
	Choose int                `json:"choose" bson:"choose"`
	Type   string             `json:"type" bson:"type"`
	From   []ProficiencyMongo `json:"from" bson:"from"`
}

//ProficiencyMongo describes a base Class proficiency
type ProficiencyMongo struct {
	Name string `json:"name" bson:"name"`
}

//SkillMongo describes the skill as part of a Class proficiency or proficiency choice
type SkillMongo struct {
	Name string `json:"name" bson:"name"`
}

//ClassMongo is the mongodb type for ckass documents
type ClassMongo struct {
	ID                 bson.ObjectId            `json:"_id" bson:"_id"`
	Name               string                   `json:"name" bson:"name"`
	HitDie             int                      `json:"hit_die" bson:"hit_die"`
	ProficiencyChoices []ProficiencyChoiceMongo `json:"proficiency_choices" bson:"proficiency_choices"`
	SubClasses         []SubClassMongo          `json:"subclasses" bson:"subclasses"`
	Proficiencies      []ProficiencyMongo       `json:"proficiencies" bson:"proficiencies"`
}

type classMongoRepo struct {
	session *mgo.Session
}

func mapClassToDomain(cm *ClassMongo) *domain.Class {
	dsc := []domain.SubClass{}
	for _, sc := range cm.SubClasses {
		dsc = append(dsc, domain.SubClass{
			Name: sc.Name,
			// Description: sc.Description[0],
		})
	}
	dp := []domain.Proficiency{}
	for _, p := range cm.Proficiencies {
		dp = append(dp, domain.Proficiency{
			Name: p.Name,
		})
	}
	dpc := []domain.ProficiencyChoice{}
	for _, pc := range cm.ProficiencyChoices {
		pcf := []domain.Proficiency{}
		for _, fp := range pc.From {
			pcf = append(pcf, domain.Proficiency{
				Name: fp.Name,
			})
		}
		dpc = append(dpc, domain.ProficiencyChoice{
			Choose: pc.Choose,
			Type:   pc.Type,
			From:   pcf,
		})
	}
	rdc := &domain.Class{
		ID:                 cm.ID.Hex(),
		Name:               cm.Name,
		HitDie:             cm.HitDie,
		ProficiencyChoices: dpc,
		Proficiencies:      dp,
		SubClasses:         dsc,
	}
	log.Printf("mapped class return %+v", utils.PrettyPrint(rdc))
	return rdc
}

// GetClassByID receives a domain Class and gets a database Class matching its id
func (c classMongoRepo) GetClassByID(id string) (*domain.Class, error) {
	fmt.Printf("ID %v", id)
	i := bson.IsObjectIdHex(id)
	if !i {
		fmt.Println("not a valid hex")
		return nil, errors.New("Not a valid id")
	}
	rm := ClassMongo{
		ID: bson.ObjectIdHex(id),
	}
	err := c.session.DB("dnd5e").C("Classs").FindId(rm.ID).One(&rm)
	if err != nil {
		return nil, err
	}
	return mapClassToDomain(&rm), nil
}

// GetClassByName receives a domain Class and gets a database Class matching its name
func (c classMongoRepo) GetClassByName(name string) (*domain.Class, error) {
	rm := ClassMongo{
		Name: name,
	}
	err := c.session.DB("dnd5e").C("classes").With(c.session.Copy()).Find(bson.M{"name": name}).One(&rm)
	if err != nil {
		return nil, err
	}
	log.Printf("\n fetched class mongo: %+v", rm)
	return mapClassToDomain(&rm), nil
}

func (c classMongoRepo) GetAllClasses() (*[]*domain.Class, error) {
	log.Println("hit mongo repo get all classes")
	// fetch
	allClasses := []ClassMongo{}
	err := c.session.DB("dnd5e").C("classes").With(c.session.Copy()).Find(bson.M{}).All(&allClasses)
	if err != nil {
		return nil, err
	}
	// map
	domainClasses := []*domain.Class{}
	for _, ac := range allClasses {
		fmt.Printf("%+v", ac)
		domainClasses = append(domainClasses, mapClassToDomain(&ac))
	}
	return &domainClasses, nil
}

func (c classMongoRepo) GetSubClassByName(name string) (*domain.SubClass, error) {
	log.Printf("\nhit get subclass by name: %+v", name)
	scm := SubClassMongo{
		Name: name,
	}
	err := c.session.DB("dnd5e").C("subclasses").With(c.session.Copy()).Find(bson.M{"name": name}).One(&scm)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrSubClassNotFound
		}
		return nil, err
	}
	log.Printf("\nmongo fetched %+v", scm)
	return &domain.SubClass{
		ID:          scm.ID.Hex(),
		Name:        scm.Name,
		Flavor:      scm.Flavor,
		Description: scm.Description[0],
	}, nil
}

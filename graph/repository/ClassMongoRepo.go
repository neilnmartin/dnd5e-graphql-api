package repository

import (
	"errors"
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//SubClassMongo describes the subclass of the Class
type SubClassMongo struct {
	Name string `json:"name" bson:"name"`
}

//ProficiencyChoiceMongo describe the number and choices for Class skill proficiencies
type ProficiencyChoiceMongo struct {
	Choose int          `json:"choose" bson:"choose"`
	Type   string       `json:"type" bson:"type"`
	From   []SkillMongo `json:"from" bson:"from"`
}

//ProficiencyMongo describes a base Class skill
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
	Size               string                   `json:"size" bson:"size"` // Medium etc.
	ProficiencyChoices []ProficiencyChoiceMongo `json:"proficiency_choices" bson:"proficiency_choices"`
	SubClasses         []SubClassMongo          `json:"subclasses" bson:"subclasses"`
}

type classMongoRepo struct {
	session *mgo.Session
}

func mapClassToDomain(cm *ClassMongo) *domain.Class {
	return &domain.Class{
		ID:   cm.ID.Hex(),
		Name: cm.Name,
		// ProficiencyChoices: &cm.ProficiencyChoices,
		// SubClasses: &domain.Class{
		// 	Name: cm.SubClasses.Name,
		// },
	}
}

// GetClassByID receives a domain Class and gets a database Class matching its id
func (r classMongoRepo) GetClassByID(id string) (*domain.Class, error) {
	fmt.Printf("ID %v", id)
	i := bson.IsObjectIdHex(id)
	if !i {
		fmt.Println("not a valid hex")
		return nil, errors.New("Not a valid id")
	}
	rm := ClassMongo{
		ID: bson.ObjectIdHex(id),
	}
	err := r.session.DB("dnd5e").C("Classs").FindId(rm.ID).One(&rm)
	if err != nil {
		return nil, err
	}
	return mapClassToDomain(&rm), nil
}

// GetClassByName receives a domain Class and gets a database Class matching its name
func (r classMongoRepo) GetClassByName(name string) (*domain.Class, error) {
	rm := ClassMongo{
		Name: name,
	}
	err := r.session.DB("dnd5e").C("classes").With(r.session.Copy()).Find(bson.M{"name": name}).One(&rm)
	if err != nil {
		return nil, err
	}
	return mapClassToDomain(&rm), nil
}

func (r classMongoRepo) GetAllClasss() (*[]*domain.Class, error) {
	// fetch
	allClasses := []ClassMongo{}
	err := r.session.DB("dnd5e").C("Classs").With(r.session.Copy()).Find(bson.M{}).All(&allClasses)
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

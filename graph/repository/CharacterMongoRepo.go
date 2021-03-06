package repository

import (
	"log"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type characterMongoRepo struct {
	session *mgo.Session
}

// CharacterMongo is the mongodb type for character documents
type CharacterMongo struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Level    int           `json:"level" bson:"level"`
	Race     string        `json:"race" bson:"race"`
	Class    string        `json:"class" bson:"class"`
	SubClass string        `json:"subClass" bson:"subClass"`
	SubRace  string        `json:"subRace" bson:"subRace"`
	// background string        `json:"background" bson:"background"`
}

// InsertCharacter will create a domain Character in the database
func (cmr characterMongoRepo) InsertCharacter(dc domain.Character, du domain.User) (*domain.Character, error) {
	sc := cmr.session.Copy()
	defer sc.Close()

	c := CharacterMongo{
		ID:       bson.NewObjectId(),
		Race:     dc.Race.Name,
		SubRace:  dc.SubRace.Name,
		Class:    dc.Class.Name,
		SubClass: dc.SubClass.Name,
		Level:    0,
	}

	err := sc.DB("dnd5e").C("characters").Insert(c)
	if err != nil {
		return nil, err
	}
	err = sc.DB("dnd5e").C("characters").FindId(c.ID).One(&c)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrCharacterNotFound
		}
		return nil, err
	}
	log.Printf("\n&u after find: %+v", c)

	// convert to domain user
	cr := domain.Character{
		ID:       c.ID.Hex(),
		Race:     domain.Race{Name: c.Race},
		SubRace:  domain.SubRace{Name: c.SubRace},
		Class:    domain.Class{Name: c.Class},
		SubClass: domain.SubClass{Name: c.SubClass},
	}
	log.Printf("\ninserted, converted to domain user: \n%+v", cr)
	return &cr, nil
}

// GetCharacterByUserID will fetch a character belonging to a specific user
func (cmr characterMongoRepo) GetCharacterByUserID(id string) (*domain.Character, error) {
	sc := cmr.session.Copy()
	defer sc.Close()

	dr := domain.Race{}
	dc := domain.Class{}
	dsc := domain.SubClass{}

	mchar := domain.Character{}

	err := sc.DB("dnd5e").C("characters").Find("{\"userId\": string}").One(&mchar)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrCharacterNotFound
		}
		return nil, err
	}

	// fetch race, class, etc.

	return &domain.Character{
		Race:     dr,
		Class:    dc,
		SubClass: dsc,
	}, nil
}

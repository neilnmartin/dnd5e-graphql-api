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
}

// InsertCharacter will create a domain User in the database
func (m characterMongoRepo) InsertCharacter(dc domain.Character) (*domain.Character, error) {
	sc := m.session.Copy()
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
	log.Printf("\nu.ID: %+v", u.ID)
	err = sc.DB("dnd5e").C("users").FindId(u.ID).One(&u)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	log.Printf("\n&u after find: %+v", u)

	// convert to domain user
	ur := domain.Character{
		ID:    u.ID.Hex(),
		Email: u.Email,
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
	}
	log.Printf("\ninserted, converted to domain user: \n%+v", ur)

	return &ur, nil
}

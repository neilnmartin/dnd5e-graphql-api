package infrastructure

import (
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type raceMongo struct {
	id   bson.ObjectId `json:"id" bson:"_id"`
	name string        `json:"name" bson:"name"`
	age  string        `json:"age" bson:"age"`

	// description []bson.String `json:"desc" bson:"desc"`
	// abilityBonuses AbilityBonus
	// perks String
	// subRaces []SubRace
	// startingProficiencies []bson.String
	// startingProficiencyOptions [String]
	// languages [Language]
	// languageOptions []domain.Language
	// languageDescription String
	// traits [Trait]
	// traitOptions [Trait]
}

type raceMongoRepo struct {
	session *mgo.Session
}

func (r raceMongoRepo) GetRaceByID(ri domain.Race) (*domain.Race, error) {
	fmt.Printf("ID %v", ri.ID)
	fmt.Printf("ID TYPE %T", ri.ID)

	i := bson.IsObjectIdHex(ri.ID)
	if !i {
		fmt.Println("not a valid hex")
	}

	rm := raceMongo{
		id: bson.ObjectIdHex(ri.ID),
	}

	err := r.session.DB("dnd5e").C("races").FindId(rm.id).One(&rm)
	if err != nil {
		return nil, err
	}

	dr := domain.Race{
		ID: rm.id.Hex(),
	}

	return &dr, nil
}

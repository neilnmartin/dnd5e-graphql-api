package repository

import (
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type subRace struct {
	url  string `bson:"url"`
	name string `bson:"name"`
}
type trait struct {
	url  string `bson:"url"`
	name string `bson:"name"`
}

type raceMongo struct {
	id              bson.ObjectId `bson:"_id"`
	name            string        `bson:"name"`
	speed           int           `bson:"speed"`
	age             string        `bson:"age"`
	size            string        `bson:"size"` // Medium etc.
	alignment       string        `bson:"alignment"`
	sizeDescription string        `bson:"size_description"`
	subRaces        []subRace     `bson:"subraces"`
	traits          []trait       `bson:"traits"`

	// description []bson.String `json:"desc" bson:"desc"`
	// abilityBonuses AbilityBonus
	// perks String
	// startingProficiencies []bson.String
	// startingProficiencyOptions [String]
	// languages [Language]
	// languageOptions []domain.Language
	// languageDescription String
	// traitOptions [Trait]
}

type raceMongoRepo struct {
	session *mgo.Session
}

// GetRaceByID receives a domain Race and gets a database race matching its id
func (r raceMongoRepo) GetRaceByID(ri domain.Race) (*domain.Race, error) {
	fmt.Printf("ID %v", ri.ID)
	fmt.Printf("ID TYPE %T", ri.ID)

	i := bson.IsObjectIdHex(ri.ID)
	if !i {
		fmt.Println("not a valid hex")
		return nil, invalidIDError{
			Code:    "Invalid ID",
			Message: "Invalid ID for Race",
		}
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

// GetRaceByName receives a domain Race and gets a database race matching its name
func (r raceMongoRepo) GetRaceByName(name string) (*domain.Race, error) {
	rm := raceMongo{
		name: name,
	}

	err := r.session.DB("dnd5e").C("races").With(r.session.Copy()).Find(bson.M{"name": name}).One(&rm)
	if err != nil {
		return nil, err
	}

	dr := domain.Race{
		ID:              rm.id.Hex(),
		Name:            rm.name,
		Age:             rm.age,
		Speed:           rm.speed,
		Size:            rm.size,
		SizeDescription: rm.sizeDescription,
		Alignment:       rm.alignment,
	}

	return &dr, nil
}

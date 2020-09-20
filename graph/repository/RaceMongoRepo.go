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

//RaceMongo is the mongodb type for race documents
type RaceMongo struct {
	ID              bson.ObjectId `json:"_id" bson:"_id"`
	Name            string        `json:"name" bson:"name"`
	Speed           int           `json:"speed" bson:"speed"`
	Age             string        `json:"age" bson:"age"`
	Size            string        `json:"size" bson:"size"` // Medium etc.
	Alignment       string        `json:"alignment" bson:"alignment"`
	SizeDescription string        `json:"size_description" bson:"size_description"`
	SubRaces        []subRace     `json:"subraces" bson:"subraces"`
	Traits          []trait       `json:"traits" bson:"traits"`

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

	rm := RaceMongo{
		ID: bson.ObjectIdHex(ri.ID),
	}

	err := r.session.DB("dnd5e").C("races").FindId(rm.ID).One(&rm)
	if err != nil {
		return nil, err
	}

	dr := domain.Race{
		ID: rm.ID.Hex(),
	}

	return &dr, nil
}

// GetRaceByName receives a domain Race and gets a database race matching its name
func (r raceMongoRepo) GetRaceByName(name string) (*domain.Race, error) {
	rm := RaceMongo{
		Name: name,
	}

	err := r.session.DB("dnd5e").C("races").With(r.session.Copy()).Find(bson.M{"name": name}).One(&rm)
	if err != nil {
		return nil, err
	}

	dr := domain.Race{
		ID:              rm.ID.Hex(),
		Name:            rm.Name,
		Age:             rm.Age,
		Speed:           rm.Speed,
		Size:            rm.Size,
		SizeDescription: rm.SizeDescription,
		Alignment:       rm.Alignment,
	}

	return &dr, nil
}

func (r *raceMongoRepo) GetAllRaces() (*[]*domain.Race, error) {
	// fetch
	allRaces := []RaceMongo{}
	err := r.session.DB("dnd5e").C("races").With(r.session.Copy()).Find(bson.M{}).All(&allRaces)
	if err != nil {
		return nil, err
	}
	// map
	domainRaces := []*domain.Race{}
	for _, ar := range allRaces {
		fmt.Printf("%+v", ar)
		domainRaces = append(domainRaces, &domain.Race{
			ID:              ar.ID.Hex(),
			Name:            ar.Name,
			Age:             ar.Age,
			Speed:           ar.Speed,
			Size:            ar.Size,
			SizeDescription: ar.SizeDescription,
			Alignment:       ar.Alignment,
		})
	}
	return &domainRaces, nil
}

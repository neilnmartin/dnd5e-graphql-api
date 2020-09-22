package repository

import (
	"errors"
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SubRaceMongo struct {
	url  string `bson:"url"`
	name string `bson:"name"`
}
type TraitMongo struct {
	url  string `bson:"url"`
	name string `bson:"name"`
}

//RaceMongo is the mongodb type for race documents
type RaceMongo struct {
	ID              bson.ObjectId  `json:"_id" bson:"_id"`
	Name            string         `json:"name" bson:"name"`
	Speed           int            `json:"speed" bson:"speed"`
	Age             string         `json:"age" bson:"age"`
	Size            string         `json:"size" bson:"size"` // Medium etc.
	Alignment       string         `json:"alignment" bson:"alignment"`
	SizeDescription string         `json:"size_description" bson:"size_description"`
	SubRaces        []SubRaceMongo `json:"subraces" bson:"subraces"`
	Traits          []TraitMongo   `json:"traits" bson:"traits"`

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

func mapRaceToDomain(rm *RaceMongo) *domain.Race {
	return &domain.Race{
		ID:              rm.ID.Hex(),
		Name:            rm.Name,
		Age:             rm.Age,
		Speed:           rm.Speed,
		Size:            rm.Size,
		SizeDescription: rm.SizeDescription,
		Alignment:       rm.Alignment,
	}
}

// GetRaceByID receives a domain Race and gets a database race matching its id
func (r raceMongoRepo) GetRaceByID(id string) (*domain.Race, error) {
	fmt.Printf("ID %v", id)
	i := bson.IsObjectIdHex(id)
	if !i {
		fmt.Println("not a valid hex")
		return nil, errors.New("Not a valid id")
	}
	rm := RaceMongo{
		ID: bson.ObjectIdHex(id),
	}
	err := r.session.DB("dnd5e").C("races").With(r.session.Copy()).FindId(rm.ID).One(&rm)
	if err != nil {
		return nil, err
	}
	return mapRaceToDomain(&rm), nil
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
	return mapRaceToDomain(&rm), nil
}

func (r raceMongoRepo) GetAllRaces() (*[]*domain.Race, error) {
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
		domainRaces = append(domainRaces, mapRaceToDomain(&ar))
	}
	return &domainRaces, nil
}

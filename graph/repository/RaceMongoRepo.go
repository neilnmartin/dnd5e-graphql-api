package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SubRaceMongo describes the sub-race associated with the Race
type SubRaceMongo struct {
	ID             bson.ObjectId       `json:"_id" bson:"_id"`
	Name           string              `json:"name" bson:"name"`
	AbilityBonuses []AbilityBonusMongo `json:"ability_bonuses" bson:"ability_bonuses"`
}

// AbilityBonusMongo is the abilitybonus associated with the subrace
type AbilityBonusMongo struct {
	Name  string `json:"name" bson:"name"`
	Bonus int    `json:"bonus" bson:"bonus"`
}

//TraitMongo describes the Trait associated with the Race
type TraitMongo struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description []string      `json:"desc" bson:"desc"`
	Races       []RaceMongo   `json:"races" bson:"races"`
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
	dsr := []domain.SubRace{}
	for _, srm := range rm.SubRaces {
		dsr = append(dsr, domain.SubRace{
			Name: srm.Name,
		})
	}
	dt := []domain.Trait{}
	for _, tm := range rm.Traits {
		dt = append(dt, domain.Trait{
			Name: tm.Name,
		})
	}
	return &domain.Race{
		ID:              rm.ID.Hex(),
		Name:            rm.Name,
		Age:             rm.Age,
		Speed:           rm.Speed,
		Size:            rm.Size,
		SizeDescription: rm.SizeDescription,
		Alignment:       rm.Alignment,
		Traits:          dt,
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

func (r raceMongoRepo) GetSubRaceByName(name string) (*domain.SubRace, error) {
	sc := r.session.Copy()
	defer sc.Close()
	srm := SubRaceMongo{
		Name: name,
	}
	err := r.session.DB("dnd5e").C("subraces").With(sc).Find(bson.M{"name": name}).One(&srm)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrSubRaceNotFound
		}
		return nil, err
	}
	return &domain.SubRace{
		Name: srm.Name,
	}, nil
}

func (r raceMongoRepo) GetTraitByName(name string) (*domain.Trait, error) {
	sc := r.session.Copy()
	defer sc.Close()

	tm := TraitMongo{
		Name: name,
	}
	err := r.session.DB("dnd5e").C("traits").With(sc).Find(bson.M{"name": name}).One(&tm)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrTraitNotFound
		}
		return nil, err
	}
	return &domain.Trait{
		ID:          tm.ID.Hex(),
		Name:        tm.Name,
		Description: tm.Description[0],
	}, nil
}

func (r raceMongoRepo) GetTraitsByRaceName(rn string) (*[]*domain.Trait, error) {
	sc := r.session.Copy()
	defer sc.Close()

	tmr := []TraitMongo{}
	err := r.session.DB("dnd5e").C("traits").With(sc).Find(bson.M{"race": rn}).All(&tmr)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrTraitNotFound
		}
		return nil, err
	}
	log.Printf("\nmongo fetched %+v", tmr)
	dt := []*domain.Trait{}
	for _, tm := range tmr {
		dt = append(dt, &domain.Trait{
			ID:          tm.ID.Hex(),
			Name:        tm.Name,
			Description: tm.Description[0],
		}, nil)
	}
	return &dt, nil
}

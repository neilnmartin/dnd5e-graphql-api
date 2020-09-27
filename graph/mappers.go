package graph

import (
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
)

// mapper funcs

// i need to see if there's a more readable way to do this but seems like
// at the moment all i have to work with is for/range

func mapClassFromDomainToAPI(c domain.Class) *model.Class {

	mpc := model.ProficiencyChoices{}
	for _, dpc := range c.ProficiencyChoices {
		mf := []*model.Proficiency{}
		for _, f := range dpc.From {
			mf = append(mf, &model.Proficiency{
				Name: &f.Name,
			})
		}
		mpc = model.ProficiencyChoices{
			Choose: &dpc.Choose,
			From:   mf,
			Type:   &dpc.Type,
		}
	}
	mp := []*model.Proficiency{}
	for _, dp := range c.Proficiencies {
		mp = append(mp, &model.Proficiency{
			Name: &dp.Name,
		})
	}
	return &model.Class{
		ID:                 &c.ID,
		Name:               &c.Name,
		HitDie:             &c.HitDie,
		ProficiencyChoices: &mpc,
		Proficiencies:      mp,
		// map SavingThrows abilities
		// StartingEquipment: c.StartingEquipment,
		// ClassLevels:       c.ClassLevels,
		// SavingThrows       []*Ability         `json:"savingThrows"`
		// StartingEquipment  *StartingEquipment `json:"startingEquipment"`
		// ClassLevels        []*ClassLevel      `json:"classLevels"`
	}
}

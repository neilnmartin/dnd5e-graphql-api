package domain

// Class-related domain types

// Class describes the character class
type Class struct {
	ID                 string
	Name               string
	HitDie             int
	ProficiencyChoices []ProficiencyChoice
	Proficiencies      []Proficiency
	SavingThrows       []Ability
	StartingEquipment  StartingEquipment
	ClassLevels        []ClassLevel
	SubClasses         []SubClass
	Spellcasting       Ability
	url                string
}

//SubClass defines a character subclass
type SubClass struct {
	ID   string
	Name string
}

//ProficiencyChoice describe the number and choices for Class skill proficiencies
type ProficiencyChoice struct {
	Choose int
	Type   string
	From   []Proficiency
}

//Proficiency describes a base Class proficiency
type Proficiency struct {
	Name string
}

//ClassLevel describes the level-specific details for a class
type ClassLevel struct {
}

//StartingEquipment describes the starting equipment of a class
type StartingEquipment struct {
}

// Equipment describes a piece or item of equipment
type Equipment struct {
	ID   string
	Name string
}

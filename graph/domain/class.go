package domain

// Class-related domain types

// Class describes the character class
type Class struct {
	ID                 string
	Name               string
	HitDie             int
	ProficiencyChoices []Skill
	Proficiencies      []Skill
	SavingThrows       []Ability
	StartingEquipment  StartingEquipment
	ClassLevels        []ClassLevel
	SubClasses         []SubClass
	Spellcasting       Ability
	url                string
}

//SubClass defines a character subclass
type SubClass struct {
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

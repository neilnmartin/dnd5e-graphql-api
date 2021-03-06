package domain

// Domain package holds api and persistence-layer agnostic domain type definitions

// User describes the user of the application
type User struct {
	ID       string
	Email    string
	Password string
	Name
}

// Name type describes the name of the user. Follows SCIM attributes.
type Name struct {
	GivenName  string
	FamilyName string
}

// Character is a domain aggregate that describes a user's DnD 5E character
type Character struct {
	ID   string
	Name string // character name
	Race
	SubRace
	Class
	SubClass
	CurrentLevel  int8 // small int max 128
	Proficiencies []Proficiency
	Equipments    []Equipment
}

//Skill describes a character's ability-based skill
type Skill struct {
	Name string
}

//Ability describes the base character abilities and associated scores
type Ability struct {
	Name string
}

package domain

// This is where the actual domain types will be defined (as opposed to the types exposed on the schema)

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

// Character describes a user's DnD 5E character
type Character struct {
	ID   string
	Name string // character name
	Race
	Class
	level      int8 // small int max 128
	Equipments []Equipment
}

// Race describes the DnD 5E character Race
type Race struct {
	ID              string
	Name            string
	Age             int
	Speed           string
	Size            string
	SizeDescription string
	Alignment       string
	// SubRaces        []SubRace
	// Traits          []Traits
}

// Class describes the character class
type Class struct {
	ID   string
	Name string
}

// Equipment describes a piece or item of equipment
type Equipment struct {
	ID   string
	Name string
}

// UserRepo describes the user repository with User related methods
type UserRepo interface {
	CreateUser(u User) (*User, error)
	GetUserByID(u User) (*User, error)
	GetUserByEmail(e string) (*User, error)
}

// RaceRepo describes the race repository with Race related methods
type RaceRepo interface {
	GetRaceByID(r Race) (*Race, error)
	GetRaceByName(r Race) (*Race, error)
}

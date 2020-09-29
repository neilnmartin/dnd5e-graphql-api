package domain

// UserRepo describes the race repository with Race related methods
type UserRepo interface {
	InsertUser(u User) (*User, error)
	GetUserByID(u User) (*User, error)
	GetUserByEmail(e string) (*User, error)
	GenerateNewID() string
}

// RaceRepo describes the race repository with Race related methods
type RaceRepo interface {
	GetRaceByID(id string) (*Race, error)
	GetRaceByName(name string) (*Race, error)
	GetAllRaces() (*[]*Race, error)
	//SubRaces
	GetSubRaceByName(name string) (*SubRace, error)
	//Traits
	GetTraitByName(name string) (*Trait, error)
	GetTraitsByRaceName(rn string) (*[]*Trait, error)
}

// ClassRepo describes the race repository with Race related methods
type ClassRepo interface {
	GetClassByID(id string) (*Class, error)
	GetClassByName(name string) (*Class, error)
	GetAllClasses() (*[]*Class, error)
	//SubClass
	GetSubClassByName(name string) (*SubClass, error)
}

// CharacterRepo describes the character repository with Character related methods
type CharacterRepo interface {
	InsertCharacter(c Character) (*Character, error)
	GetCharacterByUserID(id string) (*Character, error)
}

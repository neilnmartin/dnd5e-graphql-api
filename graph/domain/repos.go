package domain

// UserRepo describes the race repository with Race related methods
type UserRepo interface {
	InsertUser(u User) (*User, error)
	GetUserByID(u User) (*User, error)
	GetUserByEmail(e string) (*User, error)
}

// RaceRepo describes the race repository with Race related methods
type RaceRepo interface {
	GetRaceByID(id string) (*Race, error)
	GetRaceByName(name string) (*Race, error)
	GetAllRaces() (*[]*Race, error)
}

// ClassRepo describes the race repository with Race related methods
type ClassRepo interface {
	GetClassByID(id string) (*Class, error)
	GetClassByName(name string) (*Class, error)
	GetAllClasses() (*[]*Class, error)
}

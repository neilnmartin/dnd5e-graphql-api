package domain

// UserRepo describes the race repository with Race related methods
type UserRepo interface {
	CreateUser(u User) (*User, error)
	GetUserByID(u User) (*User, error)
	GetUserByEmail(e string) (*User, error)
}

// RaceRepo describes the race repository with Race related methods
type RaceRepo interface {
	GetRaceByID(r Race) (*Race, error)
	GetRaceByName(r Race) (*Race, error)
	GetAllRaces() ([]Race, error)
}

// ClassRepo describes the race repository with Race related methods
type ClassRepo interface {
	GetClassByID(r Class) (*Class, error)
	GetClassByName(r Class) (*Class, error)
}

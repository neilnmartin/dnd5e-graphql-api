package domain

// define server types

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

// UserRepo describes the user repository with User related methods
type UserRepo interface {
	CreateUser(u User) User
}

package domain

import "errors"

//Domain-specific errors to avoid unintended null references

var (

	//ErrUserNotFound when repository fails to find a User matching an input
	ErrUserNotFound = errors.New("User not found")

	//ErrCharacterNotFound when repository fails to find a Character matching an input
	ErrCharacterNotFound = errors.New("Character not found")

	//ErrRaceNotFound when repository fails to find a Race matching an input
	ErrRaceNotFound = errors.New("Race not found")

	//ErrClassNotFound when repository fails to find a Class matching an input
	ErrClassNotFound = errors.New("Class not found")

	//ErrSpellNotFound when repository fails to find a Spell matching an input
	ErrSpellNotFound = errors.New("Spell not found")

	// etc.
)
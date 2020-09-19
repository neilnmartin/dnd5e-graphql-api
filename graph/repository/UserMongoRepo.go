package repository

import (
	"fmt"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

type userMongo struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Email      string        `json:"email" bson:"email"`
	GivenName  string        `json:"givenName" bson:"givenName"`
	FamilyName string        `json:"familyName" bson:"familyName"`
	Password   string        `json:"password" bson:"password"`
}

type userMongoRepo struct {
	session *mgo.Session
}

// CreateUser will create a domain User in the database
func (m userMongoRepo) CreateUser(ui domain.User) domain.User {

	u := userMongo{
		ID:         bson.NewObjectId(),
		GivenName:  ui.Name.GivenName,
		FamilyName: ui.Name.FamilyName,
		Password:   ui.Password,
	}

	m.session.DB("rest-game").C("users").Insert(u)
	m.session.DB("rest-game").C("users").FindId(u.ID).One(&u)

	// convert to domain user
	ur := domain.User{
		ID:    u.ID.Hex(),
		Email: u.Email,
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
	}

	return ur
}

// GetUserByID will receive a domain User and return a domain User with a matching id
func (m userMongoRepo) GetUserByID(ui domain.User) domain.User {
	fmt.Printf("ID %v", ui.ID)
	fmt.Printf("ID TYPE %T", ui.ID)

	i := bson.IsObjectIdHex(ui.ID)
	if !i {
		fmt.Println("not a valid hex")
	}

	u := userMongo{
		ID:         bson.ObjectIdHex(ui.ID),
		GivenName:  ui.GivenName,
		FamilyName: ui.FamilyName,
		Password:   ui.Password,
	}

	m.session.DB("rest-game").C("users").FindId(u.ID).One(&u)

	ur := domain.User{
		ID: u.ID.Hex(),
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
		Password: u.Password,
	}

	return ur
}

// GetUserByEmail will accept an email and find a user with a matching email in the database
func (m userMongoRepo) GetUserByEmail(e string) domain.User {
	fmt.Printf("email %v", e)
	fmt.Printf("email TYPE %T", e)

	u := userMongo{
		Email: e,
	}

	m.session.DB("dnd5e").C("users").Find(e).One(&u)

	ur := domain.User{
		ID: u.ID.Hex(),
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
		Password: u.Password,
	}

	return ur
}

// UpdateUser will find and update a database user with the updated values of a domain User
func (m userMongoRepo) UpdateUser(ui domain.User) domain.User {
	i := bson.IsObjectIdHex(ui.ID)
	if !i {
		fmt.Println("not a valid hex")
	}

	u := userMongo{
		ID:         bson.ObjectIdHex(ui.ID),
		GivenName:  ui.Name.GivenName,
		FamilyName: ui.Name.FamilyName,
		Password:   ui.Password,
	}

	m.session.DB("dnd5e").C("users").Update(ui, ui) // TODO

	ur := domain.User{
		ID: u.ID.Hex(),
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
		Password: u.Password,
	}

	return ur
}

// DeleteUser will delete a user and return a boolean indicating success or failure
func (m userMongoRepo) DeleteUser(ui domain.User) bool {
	return false
}

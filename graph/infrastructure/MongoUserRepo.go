package infrastructure

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

// UserMongo struct
type UserMongo struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Email      string        `json:"email" bson:"email"`
	GivenName  string        `json:"givenName" bson:"givenName"`
	FamilyName string        `json:"familyName" bson:"familyName"`
	Password   string        `json:"password" bson:"password"`
}

// MongoUserRepo is the mongodb infrastructure implementation of the UserRepo domain interface
type MongoUserRepo interface {
	CreateUser(domain.User) domain.User
	GetUserByEmail(email string) domain.User
	GetUserByID(domain.User) domain.User
	UpdateUser(domain.User) domain.User
	DeleteUser(domain.User) bool
}

func (m mongoDatasource) CreateUser(ui domain.User) domain.User {

	u := UserMongo{
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

func (m mongoDatasource) GetUserByID(ui domain.User) domain.User {
	fmt.Printf("ID %v", ui.ID)
	fmt.Printf("ID TYPE %T", ui.ID)

	i := bson.IsObjectIdHex(ui.ID)
	if !i {
		fmt.Println("not a valid hex")
	}

	u := UserMongo{
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

func (m mongoDatasource) GetUserByEmail(e string) domain.User {
	fmt.Printf("email %v", e)
	fmt.Printf("email TYPE %T", e)

	u := UserMongo{
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
func (m mongoDatasource) UpdateUser(ui domain.User) domain.User {
	i := bson.IsObjectIdHex(ui.ID)
	if !i {
		fmt.Println("not a valid hex")
	}

	u := UserMongo{
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

func (m mongoDatasource) DeleteUser(ui domain.User) bool {
	return false
}

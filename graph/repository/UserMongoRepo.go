package repository

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

//UserMongo is the mongodb type for a user document
type UserMongo struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Email      string        `json:"email" bson:"email"`
	GivenName  string        `json:"givenName" bson:"givenName"`
	FamilyName string        `json:"familyName" bson:"familyName"`
	Password   string        `json:"password" bson:"password"`
}

type userMongoRepo struct {
	session *mgo.Session
}

// InsertUser will create a domain User in the database
func (m userMongoRepo) InsertUser(ui domain.User) (*domain.User, error) {
	sc := m.session.Copy()
	defer sc.Close()

	u := UserMongo{
		ID:         bson.NewObjectId(),
		Email:      ui.Email,
		GivenName:  ui.Name.GivenName,
		FamilyName: ui.Name.FamilyName,
		Password:   ui.Password,
	}

	err := sc.DB("dnd5e").C("users").Insert(u)
	if err != nil {
		return nil, err
	}
	log.Printf("\nu.ID: %+v", u.ID)
	err = sc.DB("dnd5e").C("users").FindId(u.ID).One(&u)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	log.Printf("\n&u after find: %+v", u)

	// convert to domain user
	ur := domain.User{
		ID:    u.ID.Hex(),
		Email: u.Email,
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
	}
	log.Printf("\ninserted, converted to domain user: \n%+v", ur)

	return &ur, nil
}

// GetUserByID will receive a domain User and return a domain User with a matching id
func (m userMongoRepo) GetUserByID(ui domain.User) (*domain.User, error) {
	sc := m.session.Copy()
	defer sc.Close()

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

	err := sc.DB("dnd5e").C("users").FindId(u.ID).One(&u)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	ur := domain.User{
		ID: u.ID.Hex(),
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
		Password: u.Password,
	}

	return &ur, nil
}

// GetUserByEmail will accept an email and find a user with a matching email in the database
func (m userMongoRepo) GetUserByEmail(e string) (*domain.User, error) {
	sc := m.session.Copy()
	defer sc.Close()
	fmt.Printf("\nemail %v", e)
	fmt.Printf("\nemail TYPE %T", e)

	u := UserMongo{
		Email: e,
	}

	err := sc.DB("dnd5e").C("users").Find(bson.M{"email": e}).One(&u)
	if err != nil {
		if err.Error() == "not found" {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	log.Printf("%+v", u)

	ur := domain.User{
		ID: u.ID.Hex(),
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
		Password: u.Password,
	}

	return &ur, nil
}

// UpdateUser will find and update a database user with the updated values of a domain User
func (m userMongoRepo) UpdateUser(ui domain.User) (*domain.User, error) {
	sc := m.session.Copy()
	defer sc.Close()

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

	sc.DB("dnd5e").C("users").Update(ui, ui) // TODO

	ur := domain.User{
		ID: u.ID.Hex(),
		Name: domain.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
		Password: u.Password,
	}

	return &ur, nil
}

// DeleteUser will delete a user and return a boolean indicating success or failure
func (m userMongoRepo) DeleteUser(ui domain.User) bool {
	return false
}

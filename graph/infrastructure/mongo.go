package infrastructure

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
)

type mongoDatasource struct {
	session *mgo.Session
}

// UserMongo struct
type UserMongo struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Password  string        `json:"password" bson:"password"`
}

// CharacterMongo Struct
type CharacterMongo struct {
	ID       bson.ObjectId `json:"id"`
	Name     string        `json:"name"`
	Age      int           `json:"age"`
	ImageURL string        `json:"imageUrl"`
}

func createMongoDataSource() mongoDatasource {

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	return mongoDatasource{
		session,
	}
}

func (m mongoDatasource) CreateUser(ui domain.User) domain.User {

	u := UserMongo{
		ID:        bson.NewObjectId(),
		FirstName: ui.Name.GivenName,
		LastName:  ui.Name.FamilyName,
		Password:  ui.Password,
	}

	m.session.DB("rest-game").C("users").Insert(u)
	m.session.DB("rest-game").C("users").FindId(u.ID).One(&u)

	ur := domain.User{
		ID:        u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}

	return ur
}

func (m mongoDatasource) CreateCharacter() {

}
func (m mongoDatasource) CreateGame() {

}
func (m mongoDatasource) ReadUser(ui domain.User) domain.User {
	fmt.Printf("ID %v", ui.ID)
	fmt.Printf("ID TYPE %T", ui.ID)

	i := bson.IsObjectIdHex(ui.ID)
	if !i {
		fmt.Println("not a valid hex")
	}

	u := UserMongo{
		ID:        bson.ObjectIdHex(ui.ID),
		FirstName: ui.FirstName,
		LastName:  ui.LastName,
		Password:  ui.Password,
	}

	m.session.DB("rest-game").C("users").FindId(u.ID).One(&u)

	ur := domain.User{
		ID:        u.ID.Hex(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
	}

	return ur
}
func (m mongoDatasource) ReadCharacter() {

}
func (m mongoDatasource) ReadGame() {

}
func (m mongoDatasource) UpdateUser() {

}
func (m mongoDatasource) UpdateCharacter() {

}
func (m mongoDatasource) UpdateGame() {

}
func (m mongoDatasource) DeleteUser() {

}
func (m mongoDatasource) DeleteCharacter() {

}
func (m mongoDatasource) DeleteGame() {

}

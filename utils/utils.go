package utils

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/neilnmartin/dnd5e-graphql-api/config"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
)

// TODO

//PrettyPrint logging utility to print structs as formatted JSON
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// IDToken describes the user data included in the token
type IDToken struct {
	ID   string
	Name model.Name
}

// GenerateJWT takes an IDToken and returns a generated json web token
func GenerateJWT(tokenInput IDToken) (*string, error) {
	secretKey := config.Config.IDTokenSecretKey
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = tokenInput.ID
	claims["name"] = tokenInput.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	ts, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("Error signing jwt: %+v", err)
		return nil, err
	}
	return &ts, nil
}

//ParseJWT parses a json web token string and returns the IDToken type
func ParseJWT(ts string) (*IDToken, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return config.Config.IDTokenSecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &IDToken{
			ID:   claims["id"].(string),
			Name: claims["name"].(model.Name),
		}, nil
	}
	return nil, err
}

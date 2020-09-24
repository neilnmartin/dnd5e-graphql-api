package utils

import "encoding/json"

// TODO

//PrettyPrint logging utility to print structs as formatted JSON
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// IDTokenInput describes the user data included in the token
type IDTokenInput struct {
}

// GenerateJWT takes an IDTokenInput and returns a generated json web token
func GenerateJWT(idt IDTokenInput) {

}

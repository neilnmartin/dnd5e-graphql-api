package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// RunTestAPI sets up and runs tests
func Test(t *testing.T) {
	t.Run("Log In User", func(t *testing.T) {
		var loginResponse struct {
			data struct {
				loginUser struct {
					user struct {
						id   string
						name struct {
							givenName  string
							familyName string
						}
					}
					token string
				}
			}
		}

		reqBody, err := json.Marshal(struct {
			query     string
			variables struct {
				email    string
				password string
			}
		}{
			query: `
				mutation LogInUser($email: String, ) {
					loginUser(input:{
						email:"testuser1@test.com"
						password:"Turk0"
					}){
						user{
							id
							name{
								givenName
								familyName
							}
						}
						token
					}
				}
			`,
			variables: struct {
				email    string
				password string
			}{
				email:    "testuser1@test.com",
				password: "Turk0",
			},
		})
		if err != nil {
			fmt.Println(err)
		}

		res, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			fmt.Println(err)
		}
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		res.Body.Close()
		respBody := json.Unmarshal(resBody, &loginResponse)

		fmt.Printf("%v", respBody)

		require.NotNil(t, respBody)
	})
}

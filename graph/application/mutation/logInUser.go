package mutation

import "fmt"

// LogInResponse consists of the logged in User as well as the token
type LogInResponse struct {
}

// LogInUser logs in a user by comparing password hashes and returning a token
func LogInUser() LogInResponse {
	// TODO
	fmt.Println("reached login service")
	return LogInResponse{}
}

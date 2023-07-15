package model

import (
	"fmt"
	"os"

	"github.com/bjoydeep/simple-microservice-proto/pkg/storage"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//redundant - but may be needed to do memory manipulation later?
/*
var users []User

func GetUsers() []User {
	return users
}

func AddUser(user User) User {

	users = append(users, user)
	return user
}
*/

func SetupModel() {
	println("Setting up GORM")
	err := storage.DB_.AutoMigrate(&User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to execute the query: %v\n", err)
	}
}

func (user User) String() string {
	//return println("User: ",user.ID, " Name: ", user.Name, " Email: ",user.Email)
	return fmt.Sprintf("User: %s Name: %s Email: %s", user.ID, user.Name, user.Email)
}

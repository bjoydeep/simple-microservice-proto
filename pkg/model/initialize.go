package model

import (
	"fmt"
	"os"

	"github.com/bjoydeep/simple-microservice-proto/pkg/storage"
)

type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
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

func UpdateUser(user User) {
	fmt.Println(user)
	if result := storage.DB_.Save(&user); result.Error != nil {
		println("Unable to update the User: ", result.Error)
	}
}

func SetupModel() {
	println("Setting up GORM")
	err := storage.DB_.AutoMigrate(&User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to execute the query: %v\n", err)
	}
}

// works with fmt.Println(user)
func (user User) String() string {
	//return println("User: ",user.ID, " Name: ", user.Name, " Email: ",user.Email)
	return fmt.Sprintf("User: %s Name: %s Email: %s Status: %s", user.ID, user.Name, user.Email, user.Status)
}

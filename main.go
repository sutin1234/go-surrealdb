package main

import (
	"fmt"

	"github.com/surrealdb/surrealdb.go"
	UserModel "github.com/sutin1234/go-surrealdb/User"
)

func main() {
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	}); err != nil {
		panic(err)
	}

	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}

	// Create user
	user := UserModel.User{
		Name:    "Sutin",
		Surname: "Injitt",
	}

	// Insert user
	data, err := db.Create("user", user)
	if err != nil {
		panic(err)
	}

	// Unmarshal data
	createdUser := make([]UserModel.User, 1)
	err = surrealdb.Unmarshal(data, &createdUser)
	if err != nil {
		panic(err)
	}

	// Get user by ID
	data, err = db.Select(createdUser[0].ID)
	if err != nil {
		panic(err)
	}

	// Unmarshal data
	selectedUser := new(UserModel.User)
	if err = surrealdb.Unmarshal(data, &selectedUser); err != nil {
		panic(err)
	}

	fmt.Println("surrealdb Query: ")
	fmt.Printf(
		"Name: %v "+
			"SerName: %v\n",
		selectedUser.Name,
		selectedUser.Surname)

}

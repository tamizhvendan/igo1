package main

import (
	"encoding/json"
	"fmt"

	"github.com/tamizhvendan/igo1/jwt"
)

type User struct {
	Id    jwt.Sub `json:"sub"`
	Name  string
	Admin bool
}

func main() {
	claims := `
	{
		"sub": "1234567890",
		"name": "John D",
		"admin": true
	}
	`
	var user User
	err := json.Unmarshal([]byte(claims), &user)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tamizhvendan/igo1/jwt"
)

type UserJwt struct {
	Sub   string
	Name  string
	Admin bool
}

func (u *UserJwt) Id() (uint, error) {
	v, err := strconv.ParseUint(u.Sub, 10, strconv.IntSize)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

type User struct {
	Id    jwt.Sub `json:"sub"`
	Name  string
	Admin bool
}

func naiveApproach() {
	claims := `
	{
		"sub": "1234567890",
		"name": "John D",
		"admin": true
	}
	`
	var userJwt *UserJwt
	err := json.Unmarshal([]byte(claims), &userJwt)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userJwt.Id())
}

func main() {
	naiveApproach()
}

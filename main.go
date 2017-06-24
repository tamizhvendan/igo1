package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tamizhvendan/igo1/jwt"
)

type UserJwtNaive struct {
	Sub   string
	Name  string
	Admin bool
}

func (u *UserJwtNaive) Id() (uint, error) {
	v, err := strconv.ParseUint(u.Sub, 10, strconv.IntSize)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func naiveApproach(claims string) {

	var userJwt *UserJwtNaive
	err := json.Unmarshal([]byte(claims), &userJwt)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userJwt.Id())
}

type UserJwt struct {
	Id    jwt.Sub `json:"sub"`
	Name  string
	Admin bool
}

func betterApproach(claims string) {

	var userJwt *UserJwt
	err := json.Unmarshal([]byte(claims), &userJwt)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userJwt.Id)
}

func main() {
	claims := `
	{
		"sub": "name/john",
		"name": "John D",
		"admin": true
	}
	`
	//naiveApproach(claims)
	betterApproach(claims)
}

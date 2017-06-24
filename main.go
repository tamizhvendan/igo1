package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type JwtSub uint

func (s *JwtSub) UnmarshalJSON(b []byte) error {
	sub := strings.Replace(string(b), `"`, ``, 2)
	v, err := strconv.ParseUint(sub, 10, strconv.IntSize)
	if err != nil {
		return err
	}
	*s = JwtSub(uint(v))
	return nil
}

type User struct {
	Id    JwtSub `json:"sub"`
	Name  string
	Admin bool
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	claims := `
	{
		"sub": "1234567890",
		"name": "John Doe",
		"admin": true
	}
	`
	var user User
	err := json.Unmarshal([]byte(claims), &user)
	panicOnError(err)

	fmt.Println(user)
}

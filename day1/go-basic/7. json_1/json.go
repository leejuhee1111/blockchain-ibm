package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId string `json:"userId"`,
	UserName string `json:"userName"`,
	UserPassword string `json:"userPassword"`,
	RegDate string `json:"regDate"`,
}

func main() {
	user := User{
		UserId: 		"user01",
		UserName: 		"홍길동",
		UserPassword: 	"1q2w3e4r@",
		RegDate: 		"20190101",
	}

	jsonBytes, err := json.Marshal(user)
	if err = nil {
		panic(err)
	}

	jsonString := string(jsonBytes)

	fmt.Println(jsonString0)
}
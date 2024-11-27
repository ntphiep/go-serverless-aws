package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
)

type Books struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func list() (string, error) {
	books := []Books{
		{Id: 1, Name: "NodeJS", Author: "NodeJS"},
		{Id: 2, Name: "Golang", Author: "Golang"},
		{Id: 3, Name: "Python", Author: "Hiep"},
	}

	res, _ := json.Marshal(&books)
	return string(res), nil
}

// func handleRequest() (string, error) {
// 	return "Hello there my fried friend", nil
// }

func main() {
	lambda.Start(list)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type JsonModel struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	var jsonData JsonModel
	data, err := os.ReadFile("./test.json")
	checkError(err)

	err = json.Unmarshal(data, &jsonData)
	checkError(err)

	jsonData.Email = "johndoe@example.com"
	jsonData.Age++
	fmt.Println(jsonData)
}

func checkError (err error) {
	if err != nil {
		log.Fatal(err)
	}
}
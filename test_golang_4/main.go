package main

import (
	"fmt"
	"net/http"
	"test_golang_4/controller"
)

func main() {
	http.HandleFunc("/items", controller.ItemController)

	http.ListenAndServe(":8067", nil)
	fmt.Println("Server is running in port :8067")
}
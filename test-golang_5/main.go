package main

import (
	"net/http"
	"test_golang_5/app"
	"test_golang_5/controller"

	"github.com/gorilla/mux"
)

func main() {
	app.InitAppAttribute()
	r := mux.NewRouter()

	r.HandleFunc("/product", controller.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/product", controller.CreateProducts).Methods(http.MethodPost)
	r.HandleFunc("/product/{ID}", controller.UpdateProducts).Methods(http.MethodPut)
	r.HandleFunc("/product/{ID}", controller.DeleteProducts).Methods(http.MethodDelete)
	
	http.ListenAndServe(":8073", r)
}
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRounting() {

	router := mux.NewRouter()

	router.HandleFunc("/employee", GetEmployees).Methods("GET")
	router.HandleFunc("/employee/{eid}", GetEmployeeById).Methods("GET")
	router.HandleFunc("/employee", CreateEmployees).Methods("POST")
	router.HandleFunc("/employee/{eid}", UpdateEmployees).Methods("PUT")
	router.HandleFunc("/employee/{eid}", DeleteEmployees).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))

}

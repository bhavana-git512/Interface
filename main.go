package main

import (
	"instance/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ret := controllers.Object()
	if ret != nil {

	}
	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/smile/{name}", controllers.Search).Methods("POST")
	log.Fatal(http.ListenAndServe(":2000", myRoute))

}

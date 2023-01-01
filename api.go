package controllers

import (
	"encoding/json"
	
	"net/http"
	"instance/model"

	"github.com/gorilla/mux"
)

type Product interface {
	Search(res http.ResponseWriter, req *http.Request)
	Signup(res http.ResponseWriter, req *http.Request)
	Order(res http.ResponseWriter, req *http.Request)
}

type Watch struct {
	hp model.Connect // interface objects

	// pointer to struct
}

func Object() *Watch {

	return &Watch{}

}

func Signup(res http.ResponseWriter, req *http.Request) {

}

func Order(res http.ResponseWriter, req *http.Request) {

}

func Search(res http.ResponseWriter, req *http.Request) {

	ds := mux.Vars(req)

	copy := ds["name"]

	if copy != "" {

	}

	nandu := model.Good{"good","nice"}

	//p.ttd.Save(&pradeep)
	

	json.NewEncoder(res).Encode(nandu)

}
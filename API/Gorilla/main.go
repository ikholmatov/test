package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var ff user

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/user/create", Create).Methods("POST")
	router.HandleFunc("/user/get/{id}", Get).Methods("GET")
	router.HandleFunc("/user/update/{id}", Update).Methods("PUT")
	router.HandleFunc("/user/delete/{id}", Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
func Create(w http.ResponseWriter, r *http.Request) {
	var newUser user
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	err = json.Unmarshal(req, &newUser)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	data, err := newUser.Create(newUser)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
func Get(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["id"]
	get, err := ff.Get(Id)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(get)
}
func Update(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["id"]
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	json.Unmarshal(req, &ff)

	update, err := ff.Update(Id, ff.PhoneNumber)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	json.NewEncoder(w).Encode(update)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	Id := mux.Vars(r)["id"]
	delete, err := ff.Delete(Id)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	json.NewEncoder(w).Encode(delete)
}

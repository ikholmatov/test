package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	//ff := user{}
	//get, err := ff.Get(`4170c2ce-1c1d-4e88-8a60-c3845a963c22`)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(get)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/user/create", create).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}
func create(w http.ResponseWriter, r *http.Request) {
	var newUser user
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "we could not dealing")
	}

	err = json.Unmarshal(req, &newUser)
	if err != nil {
		fmt.Println(err)
	}
	data, err := newUser.Create(newUser)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(data)
	fmt.Println(data)
}

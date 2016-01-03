package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// ****************
	// STATIC FILES
	// ****************
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// ****************
	// REST APIs
	// ****************
	restAPIList := func(w http.ResponseWriter, r *http.Request) {

	}
	// REST-API: /list
	router.HandleFunc("/api/list", restAPIList)

	// REST-API: /create
	//           Create takes a JSON create request that was posted.
	//	router.HandlerFunc("/api/create/{
	// TODO: https://github.com/corylanou/tns-restful-json-api/tree/master/v9
	//       Nice logging and clean organization...
	restAPICreate := func(w http.ResponseWriter, r *http.Request) {
		// Read in the posted JSON request.
		// Generate a map from that.
		// Use the information to create a new address.

	}
	router.HandleFunc("/api/create", restAPICreate).Methods("POST")

	fmt.Println("Web server listening...")
	http.ListenAndServe(":8080", router)
}

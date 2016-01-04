package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rcritt/addressdemo/apis"
	"github.com/rcritt/addressdemo/orm"
	"io"
	"io/ioutil"
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
		// Call the List API.
		var addresses []orm.AddressInfo = apis.List()

		// Convert to JSON.
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		// Return JSON.
		if err := json.NewEncoder(w).Encode(addresses); err != nil {
			panic(err)
		}

	}
	// REST-API: /list
	router.HandleFunc("/api/list", restAPIList)

	// REST-API: /create
	//           Create takes a JSON create request that was posted.
	//	router.HandlerFunc("/api/create/{
	// TODO: https://github.com/corylanou/tns-restful-json-api/tree/master/v9
	//       Nice logging and clean organization...
	restAPICreate := func(w http.ResponseWriter, r *http.Request) {
		var addressInfo orm.AddressInfo

		// Read in the posted JSON request.
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		// Generate a map from that.
		if err := json.Unmarshal(body, &addressInfo); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		fmt.Println("Here in create" + string(body))
		// Use the information to create a new address.
		apis.Create(addressInfo)

	}
	router.HandleFunc("/api/create", restAPICreate).Methods("POST")

	fmt.Println("Web server listening...")
	http.ListenAndServe(":8080", router)
}

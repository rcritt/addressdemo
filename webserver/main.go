package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("P=" + http.Dir("static"))
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	fmt.Println("Web server listening...")
	http.ListenAndServe(":8080", router)
}

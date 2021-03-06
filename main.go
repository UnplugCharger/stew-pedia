package main

import (
	"fmt"
	"net/http"

	"github.com/byron/stew-pedia/apis"
	"github.com/gorilla/mux"
)



func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare a static file directory and point it to the assests

	staticFileDirectory := http.Dir("./assets/")

	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/dish", apis.GetStewHandler).Methods("GET")
	r.HandleFunc("/dish", apis.CreateStewHandle).Methods("POST")
	return r
}

func main() {
	// Declare a new route
	r := newRouter()
    
	
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello  Mahinya")
}

package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

var words []EnglishWord

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/words", addWord).Methods("POST")
	router.HandleFunc("/words", getAllWords).Methods("GET")
	router.HandleFunc("/words/{id}", getWord).Methods("GET")
	router.HandleFunc("/words/{id}", deleteWord).Methods("DELETE")
	router.HandleFunc("/words/{id}", patchWord).Methods("PATCH")
	router.HandleFunc("/words/{id}", updateWord).Methods("PUT")

	http.ListenAndServe(":5000", router)
}

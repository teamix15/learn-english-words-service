package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func addWord(w http.ResponseWriter, r *http.Request) {
	var newWord EnglishWord
	json.NewDecoder(r.Body).Decode(&newWord)

	words = append(words, newWord)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

func getAllWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

func getWord(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatalln(err)
	}

	if id >= len(words) {
		w.Write([]byte("Id is more then length of words"))
		return
	}

	word := words[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}

func deleteWord(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatalln(err)
	}

	if id >= len(words) {
		w.Write([]byte("Id is more then length of words"))
		return
	}

	words = append(words[:id], words[id+1:]...)

	w.WriteHeader(200)
}

func patchWord(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatalln(err)
	}

	if id >= len(words) {
		w.Write([]byte("Id is more then length of words"))
		return
	}

	patchWord := &words[id]
	json.NewDecoder(r.Body).Decode(patchWord)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patchWord)
}

func updateWord(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatalln(err)
	}

	if id >= len(words) {
		w.Write([]byte("Id is more then length of words"))
		return
	}

	var updatedWord EnglishWord
	json.NewDecoder(r.Body).Decode(&updatedWord)

	words[id] = updatedWord

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedWord)

}

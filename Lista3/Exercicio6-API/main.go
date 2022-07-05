package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/lista/emcima", ListaEmCima).Methods("GET")
	router.HandleFunc("/lista/aolado", ListaAoLado).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func ListaEmCima(w http.ResponseWriter, r *http.Request) {
	var resultado string

	for i := 1; i <= 20; i++ {
		resultado += fmt.Sprintln(i)
	}

	json.NewEncoder(w).Encode(resultado)
}

func ListaAoLado(w http.ResponseWriter, r *http.Request) {
	var resultado string

	for i := 1; i <= 20; i++ {
		resultado += fmt.Sprint(i)
	}
	json.NewEncoder(w).Encode(resultado)
}

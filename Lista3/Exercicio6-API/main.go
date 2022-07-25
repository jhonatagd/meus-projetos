package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/lista/emcima", ListaEmCima).Methods("GET")
	router.HandleFunc("/lista/aolado", ListaAoLado).Methods("GET")
	router.HandleFunc("/lista/aolado", imprimi).Methods("GET")
	router.HandleFunc("/imparoupar/{id}", ImparOuPar).Methods("GET")
	router.HandleFunc("/imparoupar2", ImparOuPar2).Methods("GET")
	router.HandleFunc("/imparoupar3", ImparOuPar3).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func ListaEmCima(w http.ResponseWriter, r *http.Request) {
	var resultado string

	for i := 1; i <= 20; i++ {
		resultado += fmt.Sprintln(i)


	json.NewEncoder(w).Encode(resultado)
}

func ListaAoLado(w http.ResponseWriter, r *http.Request) {
	var resultado string

	for i := 1; i <= 20; i++ {
		resultado += fmt.Sprint(i)
	}
	json.NewEncoder(w).Encode(resultado)
}

func ImparOuPar(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id := parametros["id"]
	convertInt, _ := strconv.Atoi(id)
	if convertInt%2 == 0 {
		json.NewEncoder(w).Encode("É par")
	} else {
		json.NewEncoder(w).Encode("É Impar")
	}
}

func ImparOuPar2(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("Id")
	convertInt, _ := strconv.Atoi(id)
	if convertInt%2 == 0 {
		json.NewEncoder(w).Encode("É par")
	} else {
		json.NewEncoder(w).Encode("É Impar")
	}
}

func imprimi(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("Id")
	convertInt, _ := strconv.Atoi(id)
	if convertInt%2 == 0 {
		json.NewEncoder(w).Encode("É par")
	} else {
		json.NewEncoder(w).Encode("É Impar")
	}
}

type Requestbody struct {
	Id string
}

func ImparOuPar3(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t Requestbody
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	convertInt, _ := strconv.Atoi(t.Id)
	if convertInt%2 == 0 {
		json.NewEncoder(w).Encode("É par")
	} else {
		json.NewEncoder(w).Encode("É Impar")
	}
}

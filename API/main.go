package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", BuscarUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", BuscarUsuarioPorID).Methods("GET")
	router.HandleFunc("/usuarios", SalvarUsuario).Methods("POST")
	router.HandleFunc("/usuarios/{id}", AtualizarUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", DeletarUsuario).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type Usuario struct {
	ID   int
	Nome string
}

var usuarios []Usuario

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}

func BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id := parametros["id"]
	convertInt, err := strconv.Atoi(id)
	if err != nil {
		json.NewEncoder(w).Encode("400 - Bad Request")
		return
	}

	for i := 0; i < len(usuarios); i++ {
		if usuarios[i].ID == convertInt {
			json.NewEncoder(w).Encode(usuarios[i])
			return
		}

	}
	json.NewEncoder(w).Encode("404 - Not Found - Usuário não encontrado")
}

func SalvarUsuario(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var usuario Usuario
	err := decoder.Decode(&usuario)
	if err != nil {
		json.NewEncoder(w).Encode("400 - Bad Request")
	}

	usuarios = append(usuarios, Usuario{
		ID:   len(usuarios) + 1,
		Nome: usuario.Nome,
	})

	json.NewEncoder(w).Encode("200 - Sucesso - Usuário salvo com sucesso")
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id := parametros["id"]
	convertInt, err := strconv.Atoi(id)
	if err != nil {
		json.NewEncoder(w).Encode("400 - Bad Request")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var usuario Usuario
	err = decoder.Decode(&usuario)
	if err != nil {
		json.NewEncoder(w).Encode("400 - Bad Request")
	}

	for i := 0; i < len(usuarios); i++ {
		if usuarios[i].ID == convertInt {
			usuarios[i].Nome = usuario.Nome
			json.NewEncoder(w).Encode("200 - Usuário alterado com sucesso")
			return
		}

	}
	json.NewEncoder(w).Encode("404 - Not Found - Usuário não encontrado")
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id := parametros["id"]
	convertInt, err := strconv.Atoi(id)
	if err != nil {
		json.NewEncoder(w).Encode("400 - Bad Request")
		return
	}

	for i := 0; i < len(usuarios); i++ {
		if usuarios[i].ID == convertInt {
			//Fazer depois
			json.NewEncoder(w).Encode("200 - Usuário deletado com sucesso")
			return
		}

	}
	json.NewEncoder(w).Encode("404 - Not Found - Usuário não encontrado")
}

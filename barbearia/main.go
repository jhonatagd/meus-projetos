package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//O barbeiro tem uma cadeira de barbeiro em uma sala de corte e uma sala de espera contendo várias cadeiras.
//Quando o barbeiro termina de cortar o cabelo de um cliente
//, ele dispensa o cliente e vai para a sala de espera para ver se há outros esperando.
//Se houver, ele traz um deles de volta para a cadeira e corta o cabelo.
//Se não houver nenhum, ele volta para a cadeira e dorme nela.
//Cada cliente, quando chega, olha para ver o que o barbeiro está fazendo.
//Se o barbeiro estiver dormindo, o cliente o acorda e senta na cadeira da sala de corte.
//Se o barbeiro estiver cortando o cabelo, o cliente fica na sala de espera.
//Se houver uma cadeira livre na sala de espera, o cliente senta-se e espera a sua vez.
//Se não houver cadeira livre, o cliente sai.

var cadeiraBarbeiro bool   //true - ocupado, false - livre
var cadeirasEspera [3]bool //true - ocupada, false - livre

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/verificar-barbeiro", VerificarBarbeiro).Methods("GET")
	router.HandleFunc("/termina-corte", TerminarCorte).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func VerificarBarbeiro(w http.ResponseWriter, r *http.Request) {
	if cadeiraBarbeiro == true {
		for i := 0; i < len(cadeirasEspera); i++ {
			if cadeirasEspera[i] == false {
				cadeirasEspera[i] = true
				json.NewEncoder(w).Encode("O barbeiro está ocupado e tinha lugar na sala de espera")
				cadeiraBarbeiro = true
				return
			}
		}

		json.NewEncoder(w).Encode("O barbeiro está ocupado e não tinha lugar na sala de espera. Vou embora!!!")
		return
	} else {
		json.NewEncoder(w).Encode("Acordei o barbeiro e estou cortando o cabelo")
		cadeiraBarbeiro = true
	}
}

func TerminarCorte(w http.ResponseWriter, r *http.Request) {
	if cadeiraBarbeiro == true {
		cadeiraBarbeiro = false //Liberar cadeira do barbeiro

		for i := 0; i < len(cadeirasEspera); i++ {
			if cadeirasEspera[i] == true {
				cadeirasEspera[i] = false
				cadeiraBarbeiro = true
				ArrumarFila()
				json.NewEncoder(w).Encode("Vou chamar o proxima da fila de espera.")
				return
			}
		}

		if cadeiraBarbeiro == false {
			json.NewEncoder(w).Encode("Não tem ninguem esperando! Vou dormir")
			return
		}
	} else {
		json.NewEncoder(w).Encode("Eu não estou cortando! Estou dormindo!!!")
		cadeiraBarbeiro = true
		return
	}
}

func ArrumarFila() {
	for i := 0; i < len(cadeirasEspera); i++ {
		if i == len(cadeirasEspera)-1 { //ultima cadeira da fila de espera
			cadeirasEspera[i] = false
			return
		}

		cadeirasEspera[i] = cadeirasEspera[i+1]
	}
}

// Fila
// FIFO - First In, First Out
// [Marquito]  [Audisio]  []

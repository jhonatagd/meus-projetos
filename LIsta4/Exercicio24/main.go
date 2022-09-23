package main

import (
	"fmt"
	"math/rand"
)

//Faça um programa que simule um lançamento de dados. Lance o dado 100 vezes e armazene os resultados em um vetor .
//Depois, mostre quantas vezes cada valor foi conseguido. Dica: use um vetor de contadores(1-6) e uma função
//para gerar numeros aleatórios, simulando os lançamentos dos dados.

func main() {
	var armazenaResultados []int
	for i := 0; i <= 15; i++ {
		numerosAleatorios := rand.Intn(6) + 1
		armazenaResultados = append(armazenaResultados, numerosAleatorios)
	}
	fmt.Println(armazenaResultados)
	dado1 := 0
	dado2 := 0
	dado3 := 0
	dado4 := 0
	dado5 := 0
	dado6 := 0

	for i := 0; i < 15; i++ {
		if armazenaResultados[i] == 1 {
			dado1 += 1
		} else if armazenaResultados[i] == 2 {
			dado2 += 1
		} else if armazenaResultados[i] == 3 {
			dado3 += 1
		} else if armazenaResultados[i] == 4 {
			dado4 += 1
		} else if armazenaResultados[i] == 5 {
			dado5 += 1
		} else if armazenaResultados[i] == 6 {
			dado6 += 1
		}
	}

	fmt.Println(dado1, "o numero 1")
	fmt.Println(dado2, "o numero 2")
	fmt.Println(dado3, "o numero 3")
	fmt.Println(dado4, "o numero 4")
	fmt.Println(dado5, "o numero 5")
	fmt.Println(dado6, "o numero 6")
	fmt.Println(armazenaResultados)
}

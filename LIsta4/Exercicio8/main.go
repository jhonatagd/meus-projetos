package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um Programa que peça a idade e a altura de 5 pessoas, armazene cada informação no seu respectivo vetor.
// Imprima a idade e a altura na ordem inversa a ordem lida.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var idade [5]int
	var altura [5]float64

	for i := 0; i < 5; i++ {
		fmt.Println("Idade da", i+1, "pessoa")
		pegaIdade, _ := reader.ReadString('\n')
		limpaIdade := strings.TrimSpace(pegaIdade)
		idadeInt, err := strconv.Atoi(limpaIdade)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		} else if idadeInt < 0 {
			fmt.Println("Numero invalido")
			i--
			continue
		}
		idade[i] = idadeInt

		fmt.Println("Altura da", i+1, "pessoa")
		pegaAltura, _ := reader.ReadString('\n')
		limpaAltura := strings.TrimSpace(pegaAltura)
		alturaFloat, err := strconv.ParseFloat(limpaAltura, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		} else if alturaFloat < 0 {
			fmt.Println("Numero invalido")
			i--
			continue
		}
		altura[i] = alturaFloat

	}

	for j := 0; j < 5; j-- {
		fmt.Println(idade[j])
		fmt.Println(altura[j])
	}

}

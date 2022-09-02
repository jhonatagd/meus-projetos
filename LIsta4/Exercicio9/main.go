//Faça um Programa que leia um vetor A com 10 números inteiros, calcule e mostre a
// soma dos quadrados dos elementos do vetor.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var vetorA [10]int

	for i := 0; i < 10; i++ {
		fmt.Println(i+1, "Numero inteiro")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumeroInt := strings.TrimSpace(pegaNumero)
		numeroInt, err := strconv.Atoi(limpaNumeroInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		vetorA[i] = numeroInt
	}
	somaNumeros := 1
	for j := 0; j < 10; j++ {
		//vetorA[j] *= somaNumeros
		somaNumeros = somaNumeros * vetorA[j]
	}

	fmt.Println("Soma dos numeros", somaNumeros)
}

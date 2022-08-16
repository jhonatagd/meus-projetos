package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um Programa que leia um vetor de 10 números reais e mostre-os na ordem inversa.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var conjuntoVetor [10]int

	for i := 0; i < len(conjuntoVetor); i++ {
		fmt.Println(i+1, "Numero inteiro:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumero := strings.TrimSpace(pegaNumero)
		numero, err := strconv.Atoi(limpaNumero)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		conjuntoVetor[i] = int(numero)
	}
	for j := len(conjuntoVetor) - 1; j >= 0; j-- {
		fmt.Println(conjuntoVetor[j])
	}

}

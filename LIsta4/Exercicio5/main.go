package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um Programa que leia 20 números inteiros e armazene-os num vetor. Armazene os números pares
//no vetor PAR e os números IMPARES no vetor impar. Imprima os três vetores.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var conjuntoVetor [20]int
	var conjuntoVetorPar []int
	var conjuntoVetorImpar []int

	for i := 0; i < 20; i++ {
		fmt.Println(i+1, "Numero inteiro:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumero := strings.TrimSpace(pegaNumero)
		numero, err := strconv.Atoi(limpaNumero)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		if numero%2 == 0 {
			conjuntoVetorPar = append(conjuntoVetorPar, numero)
			conjuntoVetor[i] = numero
		} else {
			conjuntoVetorImpar = append(conjuntoVetorImpar, numero)
			conjuntoVetor[i] = numero
		}
	}
	fmt.Println("Lista com os numeros:", conjuntoVetor)
	fmt.Println("Numeros Pares:", conjuntoVetorPar)
	fmt.Println("Numeros Impares", conjuntoVetorImpar)
}

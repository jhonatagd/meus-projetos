package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um programa que peça 10 números inteiros,
// calcule e mostre a quantidade de números pares e a quantidade de números impares.
func main() {
	numerosPares := 0
	numerosImpares := 0

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println(i+1, "° numero: ")
		primeiro, _ := reader.ReadString('\n')
		limpaPrimeiro := strings.TrimSpace(primeiro)
		primeiroNumero, err := strconv.Atoi(limpaPrimeiro)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i = i - 1
			continue
		}
		if ehPar(primeiroNumero) {
			numerosPares = numerosPares + 1
		} else {
			numerosImpares = numerosImpares + 1
		}
	}
	fmt.Println("Numeros impares", numerosImpares)
	fmt.Println("Numeros Pares", numerosPares)
}
func ehPar(numero int) bool {
	return numero%2 == 0
}

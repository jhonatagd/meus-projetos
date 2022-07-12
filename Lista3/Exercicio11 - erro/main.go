package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Altere o programa anterior para mostrar no final a soma dos números.
func main() {
	//Mesmos problemas do exercicio 10
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Primeiro numero:")
	primeiro, _ := reader.ReadString('\n')
	limpaPrimeiro := strings.TrimSpace(primeiro)
	primeiroNumero, err := strconv.Atoi(limpaPrimeiro)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Segundo numero:")
	segundo, _ := reader.ReadString('\n')
	limpaSegundo := strings.TrimSpace(segundo)
	segundoNumero, err := strconv.Atoi(limpaSegundo)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}

	if primeiroNumero <= segundoNumero {
		for i := primeiroNumero; i <= segundoNumero; i++ {
			fmt.Println(i)
		}
	} else if primeiroNumero >= segundoNumero {
		for b := segundoNumero; b <= primeiroNumero; b++ {
			fmt.Println(b)
		}
	}
	fmt.Println("Soma:", segundoNumero+primeiroNumero)
}

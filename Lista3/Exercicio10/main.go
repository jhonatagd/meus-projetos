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
}

//Faça um programa que receba dois números inteiros e gere os números inteiros que estão no intervalo compreendido por eles.

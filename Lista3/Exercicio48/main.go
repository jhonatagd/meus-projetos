package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um programa que peça um numero inteiro positivo e em seguida mostre
// este numero invertido.

func main() {
	reader := bufio.NewReader(os.Stdin)
	pegaNumeroPositivo := " "

	for {
		fmt.Println("numero inteiro positivo:")
		numeros, _ := reader.ReadString('\n')
		pegaNumeros := strings.TrimSpace(numeros)
		numeroConvertido, err := strconv.Atoi(pegaNumeros)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		if numeroConvertido < 0 {
			fmt.Println("O que você digitou não é um número positivo")
			continue
		}
		pegaNumeroPositivo = pegaNumeros
		break
	}

	for i := len(pegaNumeroPositivo) - 1; i >= 0; i-- {
		fmt.Printf("%c", pegaNumeroPositivo[i])
	}
}

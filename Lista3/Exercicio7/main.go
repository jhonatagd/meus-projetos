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
	maiorValor := 0
	// Faça um programa que leia 5 números e informe o maior número.
	for i := 0; i < 5; i++ {
		fmt.Println("Digitar o ", i+1, "° numero:")
		nUm, _ := reader.ReadString('\n')
		limpaNUm := strings.TrimSpace(nUm)
		numero, err := strconv.Atoi(limpaNUm)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			i = i - 1
		} else {

			if i == 0 {
				maiorValor = numero
			} else {

				if numero > maiorValor {
					maiorValor = numero
				}
			}
		}
	}
	fmt.Println("Maior valor:", maiorValor)
}

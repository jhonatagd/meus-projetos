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

	fmt.Printf("Quantos numeros voce deseja calcular?")
	pegaNumero, _ := reader.ReadString('\n')
	limpaNumero := strings.TrimSpace(pegaNumero)
	numerosLimite, err := strconv.Atoi(limpaNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	menorValor := 0
	maiorValor := 0
	somaValores := 0

	for i := 0; i < numerosLimite; i++ {
		fmt.Println("Digitar o ", i+1, "° numero:")
		nUm, _ := reader.ReadString('\n')
		limpaNUm := strings.TrimSpace(nUm)
		numero, err := strconv.Atoi(limpaNUm)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			i = i - 1
		} else {

			if i == 0 {
				menorValor = numero
				maiorValor = numero
				somaValores = numero
			} else {
				if numero < menorValor {
					menorValor = numero
				}

				if numero > maiorValor {
					maiorValor = numero
				}
				somaValores = somaValores + numero
			}
		}
	}
	fmt.Println("Menor valor:", menorValor)
	fmt.Println("Maior valor:", maiorValor)
	fmt.Println("Soma valores:", somaValores)

} //Faça um programa que, dado um conjunto de N números, determine o menor valor,
// o maior valor e a soma dos valores.

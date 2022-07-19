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
	somaValores := 0
	numerosDeNotas := 0
	for {
		fmt.Println("Numero de notas que deseja calcular:")
		n, _ := reader.ReadString('\n')
		limpaN := strings.TrimSpace(n)
		nNumeros, err := strconv.Atoi(limpaN)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		numerosDeNotas = nNumeros
		break
	}

	for i := 0; i < numerosDeNotas; i++ {
		fmt.Println("Digitar o ", i+1, "° numero:")
		nUm, _ := reader.ReadString('\n')
		limpaNUm := strings.TrimSpace(nUm)
		numero, err := strconv.Atoi(limpaNUm)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			i = i - 1
		} else {

			if i == 0 {

				somaValores = numero
			} else {
				somaValores = somaValores + numero
			}
		}
	}
	media := somaValores / numerosDeNotas
	// Faça um programa que calcule o mostre a média aritmética de N notas.
	fmt.Println("Media aritmética:", media)

}

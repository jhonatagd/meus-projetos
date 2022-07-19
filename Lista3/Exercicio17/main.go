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
	// Errinho de validação na entrada do usuário
	calculaFat := 0
	for {
		fmt.Println("calcular o fatorial do numero:")
		calcula, _ := reader.ReadString('\n')
		limpaCalcula := strings.TrimSpace(calcula)
		calculaFatorial, err := strconv.Atoi(limpaCalcula)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		calculaFat = calculaFatorial
		break
	}

	fatorial := 1
	for i := calculaFat; i >= 1; i-- {
		fatorial = fatorial * i
	}
	fmt.Println(fatorial)

} // Faça um programa que calcule o fatorial de um número inteiro fornecido pelo usuário.
// Ex.: 5!=5.4.3.2.1=120

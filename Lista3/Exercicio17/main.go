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

	fmt.Println("calcular o fatorial do numero:")
	calcula, _ := reader.ReadString('\n')
	limpaCalcula := strings.TrimSpace(calcula)
	calculaFatorial, err := strconv.Atoi(limpaCalcula)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println(calculaFatorial)
} // Faça um programa que calcule o fatorial de um número inteiro fornecido pelo usuário. Ex.: 5!=5.4.3.2.1=120

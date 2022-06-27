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

	for {
		fmt.Println("calcular o fatorial do numero:")
		calcula, _ := reader.ReadString('\n')
		limpaCalcula := strings.TrimSpace(calcula)
		calculaFatorial, err := strconv.Atoi(limpaCalcula)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
		if calculaFatorial < 16 {
			fmt.Println("Numeros menores que 16 invalido.")

			return
		}
		fatorial := 1
		for i := calculaFatorial; i >= 1; i-- {
			fatorial = fatorial * i
		}
		fmt.Println(fatorial)
	}
} // permitindo ao usuário calcular o fatorial várias vezes
//e limitando o fatorial a números inteiros positivos e menores que 16.

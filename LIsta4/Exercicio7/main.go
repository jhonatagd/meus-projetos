package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um Programa que leia um vetor de 5 números inteiros, mostre a soma, a multiplicação e os números.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var mediaAlunos [5]int
	soma := 0
	multiplicacao := 1
	for i := 0; i < 5; i++ {
		fmt.Println(i+1, "Numero inteiro:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumero := strings.TrimSpace(pegaNumero)
		numeroInt, err := strconv.Atoi(limpaNumero)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		mediaAlunos[i] = numeroInt
		soma += mediaAlunos[i]
		multiplicacao *= mediaAlunos[i]

	}

	fmt.Println("Soma dos numeros:", soma)
	fmt.Println("Multiplicação dos numeros:", multiplicacao)
	fmt.Println("Numeros:", mediaAlunos)

}

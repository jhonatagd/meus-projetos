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
	var conjuntoVetor [5]int

	for i := 0; i < 5; i++ {
		fmt.Println(i+1, "Numero inteiro:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumero := strings.TrimSpace(pegaNumero)
		numero, err := strconv.Atoi(limpaNumero)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		conjuntoVetor[i] = int(numero)
	}
	fmt.Println(conjuntoVetor)
}

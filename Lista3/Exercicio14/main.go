package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numerosPares := 0
	numerosImpares := 0

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println("Primeiro numero: ")
		primeiro, _ := reader.ReadString('\n')
		limpaPrimeiro := strings.TrimSpace(primeiro)
		primeiroNumero, err := strconv.Atoi(limpaPrimeiro)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			break
		}
		if ehPar(primeiroNumero) {
			numerosPares = numerosPares + 1
		} else {
			numerosImpares = numerosImpares + 1
		}
	}
}
func ehPar(numero int) bool {
	return numero%2 == 0
}

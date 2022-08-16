//Faça um programa que mostre os n termos da Série a seguir:
//S = 1/1 + 2/3 + 3/5 + 4/7 + 5/9 + ... + n/m.
//Imprima no final a soma da série.
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
	nTermo := 0
	for {
		fmt.Println("N termo:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumero := strings.TrimSpace(pegaNumero)
		numeroInt, err := strconv.Atoi(limpaNumero)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		nTermo = numeroInt
		break
	}
	somaM := 1
	somaTotal := 0.0
	for i := 1; i <= nTermo; i++ {
		somaTotal += float64(i) / float64(somaM)
		somaM = somaM + 2
	}

	fmt.Println("soma da série:", somaTotal)
}

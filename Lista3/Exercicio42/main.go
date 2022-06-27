package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// leia uma quantidade indeterminada de números positivos e conte quantos deles estão nos seguintes intervalos:
// [0-25], [26-50], [51-75] e [76-100]. A entrada de dados deverá terminar quando for lido um número negativo.
func main() {
	reader := bufio.NewReader(os.Stdin)

	entre0E25 := 0
	entre26E50 := 0
	entre51E75 := 0
	entre76E100 := 0

	fmt.Println("Quantos numeros positivos deseja verificar?")
	numeros, _ := reader.ReadString('\n')
	limpaNumeros := strings.TrimSpace(numeros)
	xNumerosPositivos, err := strconv.Atoi(limpaNumeros)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	test := 0
	for i := 0; i < xNumerosPositivos; i++ {

		fmt.Println(test+1, "° Numero")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumerosPositivos := strings.TrimSpace(pegaNumero)
		numeroPositivo, err := strconv.Atoi(limpaNumerosPositivos)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		} else if numeroPositivo < 0 {
			fmt.Printf("numeros de 0 a 25: %v Numeros de 26 a 50: %v Numeros de 51 a 75: %v Numeros de 76 a 100: %v", entre0E25, entre26E50, entre51E75, entre76E100)
		}
		test = test + 1

		if numeroPositivo > 0 && numeroPositivo < 25 {
			entre0E25 = entre0E25 + 1
		}
		if numeroPositivo > 26 && numeroPositivo < 50 {
			entre26E50 = entre26E50 + 1
		}
		if numeroPositivo > 51 && numeroPositivo < 75 {
			entre51E75 = entre51E75 + 1
		}
		if numeroPositivo > 76 && numeroPositivo < 10 {
			entre76E100 = entre76E100 + 1
		}
	}

}

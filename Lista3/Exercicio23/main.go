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
	numerosPrims := 0
	for {
		fmt.Println("Mostrar numeros primos de 1 a :")
		verifica, _ := reader.ReadString('\n')
		limpaVerifica := strings.TrimSpace(verifica)
		numerosPrimos, err := strconv.Atoi(limpaVerifica)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		numerosPrims = numerosPrimos
		break
	}
	numeroDeDivsoes := 0
	for i := 0; i < numerosPrims; i++ {

		if numerosPrims == 2 {
			fmt.Println("é primo")
			numeroDeDivsoes = numeroDeDivsoes + 1
		}
		verificaPar := numerosPrims % 2
		if verificaPar == 0 {
			fmt.Println("Nâo é primo")
		}
		numeroString := strconv.Itoa(numerosPrims)
		soma := 0
		for i := 0; i < len(numeroString); i++ {
			test, _ := strconv.Atoi(string(numeroString[i]))
			soma = soma + test
		}
		if soma%3 == 0 || soma%5 == 0 || soma%7 == 0 || soma%11 == 0 {
			fmt.Println("é primo")
			numeroDeDivsoes = numeroDeDivsoes + 1
		}
	}
}

// mostre todos os primos entre 1 e N sendo N um número inteiro fornecido pelo usuário.
// O programa deverá mostrar também o número de divisões que ele executou para encontrar os números primos

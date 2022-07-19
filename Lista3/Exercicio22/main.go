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

	verificaNumero := 0
	for {
		fmt.Println("Verifica se é numero primo: ")
		verifica, _ := reader.ReadString('\n')
		limpaVerifica := strings.TrimSpace(verifica)
		numero, err := strconv.Atoi(limpaVerifica)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		verificaNumero = numero
		break
	}

	//Altere o programa de cálculo dos números primos, informando,
	//caso o número não seja primo, por quais número ele é divisível.

	if verificaNumero == 2 {
		fmt.Println("é primo")
		//	return
	}

	verificaPar := verificaNumero % 2
	if verificaPar == 0 {
		fmt.Println("Nâo é primo")
		//	return
	}
	//numeroString := strconv.Itoa(numero)
	//for i := 0; i < len(numeroString); i++ {
	//		test, _ := strconv.Atoi(string(numeroString[i]))
	//		soma = soma + test
	//	}

	var divisores []int
	// slice : slice não tem um tamanho fixo, array tem
	for i := 1; i < verificaNumero; i++ {
		testeDv := verificaNumero % i
		if testeDv == 0 {
			divisores = append(divisores, i)
		}
	}
	tamanhoLen := len(divisores)
	if tamanhoLen == 1 {
		fmt.Println("é primo")
		return
	}

	fmt.Println("Divisores de ", verificaNumero, ":", divisores)

}

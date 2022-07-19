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
	nPrimo := 0
	for {
		fmt.Println("N numero primo: ")
		verifica, _ := reader.ReadString('\n')
		limpaVerifica := strings.TrimSpace(verifica)
		numero, err := strconv.Atoi(limpaVerifica)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		nPrimo = numero
		break
	}
	for i := 0; i < nPrimo; i++ {
		if nPrimo == 2 {
			fmt.Println("é primo")
			return
		}
		verificaPar := nPrimo % 2
		if verificaPar == 0 {
			fmt.Println("Nâo é primo")
			return
		}
		numeroString := strconv.Itoa(nPrimo)
		soma := 0
		for i := 0; i < len(numeroString); i++ {
			test, _ := strconv.Atoi(string(numeroString[i]))
			soma = soma + test
		}
		if soma%3 == 0 || soma%5 == 0 || soma%7 == 0 || soma%11 == 0 {
			fmt.Println("é primo")
		}

	}

}

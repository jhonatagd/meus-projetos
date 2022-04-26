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
	fmt.Println("Valor: ")
	pegaValor, _ := reader.ReadString('\n')
	limpaValor := strings.TrimSpace(pegaValor)
	pegaValorInt, err := strconv.Atoi(limpaValor)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	if pegaValorInt == 0 {
		fmt.Println("numero 0 não é positivo nem negativo")
	} else if pegaValorInt > 0 {
		fmt.Println(pegaValorInt, "é um numero positivo")
	} else {
		fmt.Println(pegaValorInt, "é um numero negativo")
	}
}

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

	fmt.Println("Numero base:")
	pegaBase, _ := reader.ReadString('\n')
	limpaBase := strings.TrimSpace(pegaBase)
	base, err := strconv.Atoi(limpaBase)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Numero expoente:")
	pegaExpoente, _ := reader.ReadString('\n')
	limpaExpoente := strings.TrimSpace(pegaExpoente)
	expoente, err := strconv.Atoi(limpaExpoente)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	resultado := 1
	for i := 0; i < expoente; i++ {
		resultado = resultado * base
	}
	fmt.Println(resultado)
}

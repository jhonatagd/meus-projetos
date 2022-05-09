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
	base, _ := reader.ReadString('\n')
	limpaBase := strings.TrimSpace(base)
	pegaBase, err := strconv.Atoi(limpaBase)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Numero expoente:")
	expoente, _ := reader.ReadString('\n')
	limpaExpoente := strings.TrimSpace(expoente)
	pegaExpoente, err := strconv.Atoi(limpaExpoente)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}

	for i := ; i <= pegapegaBaseExpoente; i++ {
		fmt.Println(i)
	}
	for i := pegaBase; i >= pegaExpoente; i++ {
		fmt.Println(i)
	} // terminar -----------------------------------------
}

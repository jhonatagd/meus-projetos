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

	fmt.Println("calcular o fatorial do numero:")
	calcula, _ := reader.ReadString('\n')
	limpaCalcula := strings.TrimSpace(calcula)
	calculaFatorial, err := strconv.Atoi(limpaCalcula)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fatorial := 1
	for i := calculaFatorial; i >= 1; i-- {
		fatorial = fatorial * i
		fmt.Print(calculaFatorial, ".")
		calculaFatorial = calculaFatorial - 1
	}
	fmt.Println(fatorial)

}

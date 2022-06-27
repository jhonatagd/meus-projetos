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

	fmt.Println("N:")
	termo, err := reader.ReadString('\n')
	termo = strings.TrimSpace(termo)
	n, err := strconv.Atoi(termo)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	primeiro := 1
	segundo := 1
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Println(primeiro)
		} else if i == 1 {
			fmt.Println(segundo)
		} else {
			novoNumero := primeiro + segundo
			fmt.Println(novoNumero)
			primeiro = segundo
			segundo = novoNumero
		}
	}

}

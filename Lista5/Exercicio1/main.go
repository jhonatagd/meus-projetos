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
	atualizaN := 0
	for {
		fmt.Println("Digite um N:")
		n, _ := reader.ReadString('\n')
		limpaN := strings.TrimSpace(n)
		pegaN, err := strconv.Atoi(limpaN)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		atualizaN = pegaN
		break
	}
	for i := 1; i <= atualizaN; i++ {
		escreveLinha(i)
	}
}

func escreveLinha(n int) {
	auxilia := ""
	for i := 0; i < n; i++ {
		auxilia += fmt.Sprintf("%d ", n)
	}
	fmt.Println(auxilia)
}

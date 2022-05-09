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

	fmt.Println("Verifica se é numero primo: ")
	verifica, _ := reader.ReadString('\n')
	limpaVerifica := strings.TrimSpace(verifica)
	numero, err := strconv.Atoi(limpaVerifica)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	dividiPor1 := numero / 1
	verificaPrimo := numero / numero
	if dividiPor1 == 0 && verificaPrimo == 0 {
		fmt.Println(numero, "é um numero primo")
	} else {
		fmt.Println(numero, "não é um numero primo")
	}
}

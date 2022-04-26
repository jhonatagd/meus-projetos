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
	fmt.Println("Verifica numero inteiro impar ou par:")
	verifica, _ := reader.ReadString('\n')
	limpaVerifica := strings.TrimSpace(verifica)
	verificaInt, err := strconv.Atoi(limpaVerifica)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	verificaPar := verificaInt % 2

	if verificaPar == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}
}

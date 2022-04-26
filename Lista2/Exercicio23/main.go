package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Descubra se é um número é inteiro ou decimal:")
	verifica, _ := reader.ReadString('\n')
	limpaVerifica := strings.TrimSpace(verifica)
	verificaInt, err := strconv.ParseFloat(limpaVerifica, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	teste := (math.Floor(verificaInt + 0.0/0.5))

	if verificaInt == teste {
		fmt.Println("Inteiro")
	} else {
		fmt.Println("Decimal")
	}
}

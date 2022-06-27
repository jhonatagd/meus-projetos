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

	fmt.Println("Quantidade de itens comprados:")
	i, _ := reader.ReadString('\n')
	limpaI := strings.TrimSpace(i)
	itens, err := strconv.ParseFloat(limpaI, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}

	if itens > 0 && itens <= 50 {
		fmt.Println(itens, "Itens deu um total de:", itens*1.99)
	} else {
		fmt.Println("Máximo de itens 50")
	}
}

//Lojas Quase Dois - Tabela de preços
//1 - R$ 1.99
//2 - R$ 3.98
//50 - R$ 99.50

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
	quantidadeDePaes := 0.0
	for {
		fmt.Println("Quantidade de pães comprados:")
		p, _ := reader.ReadString('\n')
		limpaP := strings.TrimSpace(p)
		quantidadePaes, err := strconv.ParseFloat(limpaP, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		quantidadeDePaes = quantidadePaes
		break
	}

	if quantidadeDePaes > 0 && quantidadeDePaes <= 50 {
		fmt.Println(quantidadeDePaes, "Pães deu um total de:", quantidadeDePaes*0.18)
	} else {
		fmt.Println("Máximo de itens 50")
	}
}

//monta a tabela de preços de pães
//de 1 até 50 pães, a partir do preço do pão informado pelo usuário, conforme o exemplo abaixo:
//Preço do pão: R$ 0.18
//Panificadora Pão de Ontem - Tabela de preços
//1 - R$ 0.18
//2 - R$ 0.36
//...
//50 - R$ 9.00

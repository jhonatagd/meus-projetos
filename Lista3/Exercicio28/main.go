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

	fmt.Println("Quantidade de CDs:")
	quantidadeCD, _ := reader.ReadString('\n')
	limpaQuantidadeCD := strings.TrimSpace(quantidadeCD)
	quantidadeDeCds, err := strconv.Atoi(limpaQuantidadeCD)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fazTotal := 0
	calculaMedia := 0
	for i := 0; i < quantidadeDeCds; i++ {
		fmt.Println("Valor gasto no", i+1, "Cd:")
		Valor, _ := reader.ReadString('\n')
		limpaValor := strings.TrimSpace(Valor)
		valorCd, err := strconv.Atoi(limpaValor)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
		}
		fazTotal = fazTotal + valorCd
		calculaMedia = fazTotal / quantidadeDeCds

	}
	fmt.Println("valor médio gasto em cada um deles:", calculaMedia)
}

//Faça um programa que calcule o valor total investido por um colecionador em sua
//coleção de CDs e o valor médio gasto em cada um deles. O usuário deverá informar
//a quantidade de CDs e o valor para em cada um.

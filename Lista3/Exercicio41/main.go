package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um programa que receba o valor de uma dívida
//e mostre uma tabela com os seguintes dados: valor da dívida, valor dos juros,
// quantidade de parcelas e valor da parcela.
func main() {
	reader := bufio.NewReader(os.Stdin)
	atualizaDivida := 0
	for {
		fmt.Println("Valor da sua divida:")
		pegaDivida, _ := reader.ReadString('\n')
		limpaDivida := strings.TrimSpace(pegaDivida)
		divida, err := strconv.Atoi(limpaDivida)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		atualizaDivida = divida
		break
	}
	juros10 := atualizaDivida / 100 * 10
	juros15 := atualizaDivida / 100 * 15
	juros20 := atualizaDivida / 100 * 20
	juros25 := atualizaDivida / 100 * 25
	fmt.Println("Valor da Dívida    Valor dos Juros    Quantidade de Parcelas     Valor da Parcela")
	fmt.Println(atualizaDivida, 0, "1", atualizaDivida)
	fmt.Println(atualizaDivida, juros10, "3", (atualizaDivida+juros10)/3)
	fmt.Println(atualizaDivida, juros15, "6", (atualizaDivida+juros15)/6)
	fmt.Println(atualizaDivida, juros20, "9", (atualizaDivida+juros20)/9)
	fmt.Println(atualizaDivida, juros25, "12", (atualizaDivida+juros25)/12)
	//Quantidade de Parcelas  % de Juros sobre o valor inicial da dívida
	//1       0
	//3       10
	//6       15
	//9       20
	//12      25
}

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

	fmt.Println("Valor da sua divida:")
	pegaDivida, _ := reader.ReadString('\n')
	limpaDivida := strings.TrimSpace(pegaDivida)
	divida, err := strconv.Atoi(limpaDivida)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	juros10 := divida / 100 * 10
	juros15 := divida / 100 * 15
	juros20 := divida / 100 * 20
	juros25 := divida / 100 * 25
	fmt.Println("Valor da Dívida    Valor dos Juros    Quantidade de Parcelas     Valor da Parcela")
	fmt.Println(divida, 0, "1", divida)
	fmt.Println(divida, juros10, "3", (divida+juros10)/3)
	fmt.Println(divida, juros15, "6", (divida+juros15)/6)
	fmt.Println(divida, juros20, "9", (divida+juros20)/9)
	fmt.Println(divida, juros25, "12", (divida+juros25)/12)
	//Quantidade de Parcelas  % de Juros sobre o valor inicial da dívida
	//1       0
	//3       10
	//6       15
	//9       20
	//12      25
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Desenvolva um gerador de tabuada, capaz de gerar a tabuada de qualquer número inteiro entre 1 a 10.
// O usuário deve informar de qual numero ele deseja ver a tabuada.
// A saída deve ser conforme o exemplo abaixo:
// Tabuada de 5:
// 5 X 1 = 5
// 5 X 2 = 10
// ...
// 5 X 10 = 50
func main() {
	// Erros
	// Não esta validando a entrada do usuário. Está aceitando qualquer valor
	// USAR FOR!!!
	// Mantem a regra criada pq esta certo! Mas faz o mesmo resultado usando for.
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite um numero de 1 a 10:")
	numero, _ := reader.ReadString('\n')
	limpaNumero := strings.TrimSpace(numero)
	pegaNumero, err := strconv.Atoi(limpaNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Tabuada do", pegaNumero, ":")
	fmt.Println(pegaNumero, "* 1 =", pegaNumero*1)
	fmt.Println(pegaNumero, "* 2 =", pegaNumero*2)
	fmt.Println(pegaNumero, "* 3 =", pegaNumero*3)
	fmt.Println(pegaNumero, "* 4 =", pegaNumero*4)
	fmt.Println(pegaNumero, "* 5 =", pegaNumero*5)
	fmt.Println(pegaNumero, "* 6 =", pegaNumero*6)
	fmt.Println(pegaNumero, "* 7 =", pegaNumero*7)
	fmt.Println(pegaNumero, "* 8 =", pegaNumero*8)
	fmt.Println(pegaNumero, "* 9 =", pegaNumero*9)
	fmt.Println(pegaNumero, "* 10 =", pegaNumero*10)
}

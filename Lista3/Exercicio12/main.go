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
	pegaN := 0

	for {
		fmt.Println("Digite um numero de 1 a 10:")
		numero, _ := reader.ReadString('\n')
		limpaNumero := strings.TrimSpace(numero)
		pegaNumero, err := strconv.Atoi(limpaNumero)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
		pegaN = pegaNumero
		break
	}

	fmt.Println("Tabuada do", pegaN, ":")
	fmt.Println(pegaN, "* 1 =", pegaN*1)
	fmt.Println(pegaN, "* 2 =", pegaN*2)
	fmt.Println(pegaN, "* 3 =", pegaN*3)
	fmt.Println(pegaN, "* 4 =", pegaN*4)
	fmt.Println(pegaN, "* 5 =", pegaN*5)
	fmt.Println(pegaN, "* 6 =", pegaN*6)
	fmt.Println(pegaN, "* 7 =", pegaN*7)
	fmt.Println(pegaN, "* 8 =", pegaN*8)
	fmt.Println(pegaN, "* 9 =", pegaN*9)
	fmt.Println(pegaN, "* 10 =", pegaN*10)
}

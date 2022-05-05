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

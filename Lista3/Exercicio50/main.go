package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Sendo H= 1 + 1/2 + 1/3 + 1/4 + ... + 1/N, Faça um programa que calcule o valor de H com N termos.

func main() {
	reader := bufio.NewReader(os.Stdin)
	pegaTermoN := 0
	for {
		fmt.Println("Numero de N termos:")
		pegaTermo, _ := reader.ReadString('\n')
		limpaTermo := strings.TrimSpace(pegaTermo)
		pegaNtermos, err := strconv.Atoi(limpaTermo)
		if err != nil {
			fmt.Println("O que você digitou não é um termo")
			continue
		} else if pegaNtermos <= 0 {
			fmt.Println("Termo invalido, digite novamente")
			continue
		}
		pegaTermoN += pegaNtermos
		break
	}
	valorDeHComNtermos := 0
	for i := 1; i <= pegaTermoN; i++ {
		n := 1
		nDivide := n / i
		valorDeHComNtermos += nDivide
	}
	fmt.Println("Valor de H com", pegaTermoN, "termos:", valorDeHComNtermos)
}

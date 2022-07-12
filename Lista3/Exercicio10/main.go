package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um programa que receba dois números inteiros e gere os números inteiros que estão no intervalo compreendido por eles.
func main() {
	//Vc esta validando o valor digitado, mas esta aceitando o erro. Repetir a solicitação ate que o usuário digite um valor correto
	//Dica, for pra cada input
	reader := bufio.NewReader(os.Stdin)
	primeiroNumero1 := 0
	segundoNumero2 := 0
	for {

		fmt.Println("Primeiro numero:")
		primeiro, _ := reader.ReadString('\n')
		limpaPrimeiro := strings.TrimSpace(primeiro)
		primeiroNumero, err := strconv.Atoi(limpaPrimeiro)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else {
			primeiroNumero1 = primeiroNumero1 + primeiroNumero
			break
		}

	}

	for {

		fmt.Println("Segundo numero:")
		segundo, _ := reader.ReadString('\n')
		limpaSegundo := strings.TrimSpace(segundo)
		segundoNumero, err := strconv.Atoi(limpaSegundo)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else {
			segundoNumero2 = segundoNumero2 + segundoNumero
			break
		}

	}
	//O exercicio pede para printar os inteiros que estão no intervalo
	//Vc esta printando tbm os valores digitados
	//Exemplo: 5 a 10 -> 6, 7, 8 e 9
	//Dica: O segredo esta no for
	if primeiroNumero1 < segundoNumero2 {
		primeiroNumero1 = primeiroNumero1 + 1
		segundoNumero2 = segundoNumero2 - 1
		for i := primeiroNumero1; i <= segundoNumero2; i++ {
			fmt.Println(i)
		}
	} else if primeiroNumero1 > segundoNumero2 {
		primeiroNumero1 = primeiroNumero1 + 1
		segundoNumero2 = segundoNumero2 - 1
		for b := segundoNumero2; b <= primeiroNumero1; b++ {
			fmt.Println(b)
		}
	}
}

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

	fmt.Println("Primeiro numero:")
	primeiro, _ := reader.ReadString('\n')
	limpaPrimeiro := strings.TrimSpace(primeiro)
	primeiroNumero, err := strconv.Atoi(limpaPrimeiro)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Segundo numero:")
	segundo, _ := reader.ReadString('\n')
	limpaSegundo := strings.TrimSpace(segundo)
	segundoNumero, err := strconv.Atoi(limpaSegundo)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}

	//O exercicio pede para printar os inteiros que estão no intervalo
	//Vc esta printando tbm os valores digitados
	//Exemplo: 5 a 10 -> 6, 7, 8 e 9
	//Dica: O segredo esta no for
	if primeiroNumero <= segundoNumero {
		for i := primeiroNumero; i <= segundoNumero; i++ {
			fmt.Println(i)
		}
	} else if primeiroNumero >= segundoNumero {
		for b := segundoNumero; b <= primeiroNumero; b++ {
			fmt.Println(b)
		}
	}
}

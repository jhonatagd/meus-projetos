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
	fmt.Println("terceiro numero:")
	terceiro, _ := reader.ReadString('\n')
	limpaTerceiro := strings.TrimSpace(terceiro)
	terceiroNumero, err := strconv.Atoi(limpaTerceiro)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Quarto numero:")
	quarto, _ := reader.ReadString('\n')
	limpaQuarto := strings.TrimSpace(quarto)
	quartoNumero, err := strconv.Atoi(limpaQuarto)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Quinto numero:")
	quinto, _ := reader.ReadString('\n')
	limpaQuinto := strings.TrimSpace(quinto)
	quintoNumero, err := strconv.Atoi(limpaQuinto)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	// Faça um programa que leia 5 números e informe a soma e a média dos números.
	soma := primeiroNumero + segundoNumero + terceiroNumero + quartoNumero + quintoNumero
	media := ((primeiroNumero + segundoNumero + terceiroNumero + quartoNumero + quintoNumero) / 5)

	fmt.Println("Soma dos números:", soma)
	fmt.Println("Média dos números:", media)
	fmt.Println("o")
}

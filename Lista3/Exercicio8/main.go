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
	primeiro, err := reader.ReadString('\n')
	primeiroNumero := strings.TrimSpace(primeiro)
	primeiroNumero = strings.ToUpper(primeiroNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Segundo numero:")
	segundo, err := reader.ReadString('\n')
	segundoNumero := strings.TrimSpace(segundo)
	segundoNumero = strings.ToUpper(segundoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("terceiro numero:")
	terceiro, err := reader.ReadString('\n')
	terceiroNumero := strings.TrimSpace(terceiro)
	terceiroNumero = strings.ToUpper(terceiroNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Quarto numero:")
	quarto, err := reader.ReadString('\n')
	quartoNumero := strings.TrimSpace(quarto)
	quartoNumero = strings.ToUpper(quartoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Quinto numero:")
	quinto, err := reader.ReadString('\n')
	quintoNumero := strings.TrimSpace(terceiro)
	quintoNumero = strings.ToUpper(quintoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	// Faça um programa que leia 5 números e informe a soma e a média dos números.
	soma := primeiroNumero + segundoNumero + terceiroNumero + quartoNumero + quintoNumero
	media := (primeiroNumero + segundoNumero + terceiroNumero + quartoNumero + quintoNumero) / 5
	
	fmt.Println("Soma dos números:", soma)
	fmt.Println("Média dos números:", media)
}
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
	quintoNumero := strings.TrimSpace(quinto)
	quintoNumero = strings.ToUpper(quintoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Sexto numero:")
	sexto, err := reader.ReadString('\n')
	sextoNumero := strings.TrimSpace(sexto)
	sextoNumero = strings.ToUpper(sextoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("sétimo numero:")
	setimo, err := reader.ReadString('\n')
	setimoNumero := strings.TrimSpace(setimo)
	setimoNumero = strings.ToUpper(setimoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Oitavo numero:")
	oitavo, err := reader.ReadString('\n')
	oitavoNumero := strings.TrimSpace(oitavo)
	oitavoNumero = strings.ToUpper(oitavoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Nono numero:")
	nono, err := reader.ReadString('\n')
	nonoNumero := strings.TrimSpace(nono)
	nonoNumero = strings.ToUpper(nonoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Decimo numero:")
	decimo, err := reader.ReadString('\n')
	decimoNumero := strings.TrimSpace(decimo)
	decimoNumero = strings.ToUpper(decimoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	numerosPares := 0
	numerosImpares := 0
	if primeiroNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if segundoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if terceiroNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if quartoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if quintoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if sextoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if setimoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if oitavoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if nonoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	if decimoNumero == 0 {
		numerosPares = numerosPares + 1
	} else {
		numerosImpares = numerosImpares + 1
	}

	fmt.Println(numerosPares, "Numeros pares", numerosImpares, "numeros impares")
}
package main

// leia três números e mostre o maior deles.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Primeiro numero: ")
	nUm, _ := reader.ReadString('\n')
	limpaNUm := strings.TrimSpace(nUm)
	nUmInt, err := strconv.Atoi(limpaNUm)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Segundo numero: ")
	nDois, _ := reader.ReadString('\n')
	limpaNDois := strings.TrimSpace(nDois)
	nDoisInt, err := strconv.Atoi(limpaNDois)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Terceiro numero: ")
	nTres, _ := reader.ReadString('\n')
	limpaNTres := strings.TrimSpace(nTres)
	nTresInt, err := strconv.Atoi(limpaNTres)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}

	numero1 := nUmInt
	numero2 := nDoisInt
	numero3 := nTresInt

	if numero1 > numero2 && numero1 > numero3 {
		fmt.Println("O numero maior é ", numero1)
	} else if numero2 > numero3 && numero2 > numero1 {
		fmt.Println("Numero maior é o: ", numero2)
	} else if numero3 > numero2 && numero3 > numero1 {
		fmt.Println("Numero maior é o: ", numero3)
	}

}

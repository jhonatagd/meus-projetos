package main

// leia três números e mostre-os em ordem decrescente.

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

	if numero1 >= numero2 && numero1 >= numero3 && numero2 >= numero3 {
		fmt.Println("A ordem decrescente é ", numero1, numero2, numero3)
	} else if numero1 >= numero2 && numero1 >= numero3 && numero3 >= numero2 {
		fmt.Println("A ordem decrescente é ", numero1, numero3, numero2)
	} else if numero2 >= numero1 && numero2 >= numero3 && numero1 >= numero3 {
		fmt.Println("A ordem decrescente é ", numero2, numero1, numero3)
	} else if numero2 >= numero1 && numero2 >= numero3 && numero3 >= numero1 {
		fmt.Println("A ordem decrescente é ", numero2, numero3, numero1)
	} else if numero3 >= numero1 && numero3 >= numero2 && numero1 >= numero2 {
		fmt.Println("A ordem decrescente é ", numero3, numero1, numero2)
	} else if numero3 >= numero1 && numero3 >= numero2 && numero2 >= numero1 {
		fmt.Println("A ordem decrescente é ", numero3, numero2, numero1)
	}

}

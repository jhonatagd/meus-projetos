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
 
		if primeiroNumero <= segundoNumero {
			for i := primeiroNumero; i <= segundoNumero; i++ {
				fmt.Println(primeiroNumero + segundoNumero)
			}
		} else if primeiroNumero >= segundoNumero {
				for b := segundoNumero; b <= primeiroNumero; b++ {
					fmt.Println(segundoNumero + primeiroNumero)
				}
			}
		}
}
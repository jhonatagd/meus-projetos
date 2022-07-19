package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Faça um programa que peça dois números, base e expoente,
// calcule e mostre o primeiro número elevado ao segundo número.
// Não utilize a função de potência da linguagem.
func main() {
	// Errinho: Validar os inputs. Está seguindo o código com qualquer valor
	reader := bufio.NewReader(os.Stdin)
	nB := 0
	nE := 0

	for {
		fmt.Println("Numero base:")
		pegaBase, _ := reader.ReadString('\n')
		limpaBase := strings.TrimSpace(pegaBase)
		base, err := strconv.Atoi(limpaBase)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
		nB = base
		break
	}

	for {
		fmt.Println("Numero expoente:")
		pegaExpoente, _ := reader.ReadString('\n')
		limpaExpoente := strings.TrimSpace(pegaExpoente)
		expoente, err := strconv.Atoi(limpaExpoente)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		nE = expoente
		break
	}

	resultado := 1
	for i := 0; i < nE; i++ {
		resultado = resultado * nB
	}
	fmt.Println(resultado)
}

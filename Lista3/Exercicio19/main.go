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

	fmt.Println("Primeiro numero: ")
	nUm, _ := reader.ReadString('\n')
	limpaNUm := strings.TrimSpace(nUm)
	numeroUm, err := strconv.Atoi(limpaNUm)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}

	fmt.Println("Segundo numero: ")
	nDois, _ := reader.ReadString('\n')
	limpaNDois := strings.TrimSpace(nDois)
	numeroDois, err := strconv.Atoi(limpaNDois)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Terceiro numero: ")
	nTres, _ := reader.ReadString('\n')
	limpaNTres := strings.TrimSpace(nTres)
	numeroTres, err := strconv.Atoi(limpaNTres)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Quarto numero: ")
	nQuatro, _ := reader.ReadString('\n')
	limpaQuatro := strings.TrimSpace(nQuatro)
	numeroQuatro, err := strconv.Atoi(limpaQuatro)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}

	menorValor := numeroUm
	maiorValor := numeroUm

	if numeroUm < 0 || numeroUm > 1000 {
		fmt.Println("Não aceitamos numeros menores que 0 e maiores que 1000.")
	}
	if numeroDois < 0 || numeroDois > 1000 {
		fmt.Println("Não aceitamos numeros menores que 0 e maiores que 1000.")
	}
	if numeroTres < 0 || numeroTres > 1000 {
		fmt.Println("Não aceitamos numeros menores que 0 e maiores que 1000.")
	}
	if numeroQuatro < 0 || numeroQuatro > 1000 {
		fmt.Println("Não aceitamos numeros menores que 0 e maiores que 1000.")
	}

	if numeroUm >= numeroDois && numeroUm >= numeroTres && numeroUm >= numeroQuatro {
		maiorValor = numeroUm
	}
	if numeroDois >= numeroUm && numeroDois >= numeroTres && numeroDois >= numeroQuatro {
		maiorValor = numeroDois
	}
	if numeroTres >= numeroUm && numeroTres >= numeroDois && numeroTres >= numeroQuatro {
		maiorValor = numeroTres
	}
	if numeroQuatro >= numeroUm && numeroQuatro >= numeroDois && numeroQuatro >= numeroTres {
		maiorValor = numeroQuatro
	}

	if numeroUm <= numeroDois && numeroUm <= numeroTres && numeroUm <= numeroQuatro {
		menorValor = numeroUm
	}
	if numeroDois <= numeroUm && numeroDois <= numeroTres && numeroDois <= numeroQuatro {
		menorValor = numeroDois
	}
	if numeroTres <= numeroUm && numeroTres <= numeroDois && numeroTres <= numeroQuatro {
		menorValor = numeroTres
	}
	if numeroQuatro <= numeroUm && numeroQuatro <= numeroDois && numeroQuatro <= numeroTres {
		menorValor = numeroQuatro
	}
	soma := numeroUm + numeroDois + numeroTres + numeroQuatro

	fmt.Println("Maior numero:", maiorValor, "menor numero:", menorValor, "soma:", soma)
} //Faça um programa que, dado um conjunto de N números, determine o menor valor, o maior valor e a soma dos valores.

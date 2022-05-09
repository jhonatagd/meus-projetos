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
	fmt.Println("Quinto numero: ")
	ncinco, _ := reader.ReadString('\n')
	limpaCinco := strings.TrimSpace(ncinco)
	numeroCinco, err := strconv.Atoi(limpaCinco)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	if numeroUm >= numeroDois && numeroUm >= numeroTres && numeroUm >= numeroQuatro && numeroUm >= numeroCinco {
		fmt.Println("O maior numero é:", numeroUm)
	} else if numeroDois >= numeroUm && numeroDois >= numeroTres && numeroDois >= numeroQuatro && numeroDois >= numeroCinco {
		fmt.Println("O maior numero é:", numeroDois)
	} else if numeroTres >= numeroUm && numeroTres >= numeroDois && numeroTres >= numeroQuatro && numeroTres >= numeroCinco {
		fmt.Println("O maior numero é:", numeroTres)
	} else if numeroQuatro >= numeroUm && numeroQuatro >= numeroDois && numeroQuatro >= numeroTres && numeroQuatro >= numeroCinco {
		fmt.Println("O maior numero é:", numeroQuatro)
	} else if numeroCinco >= numeroUm && numeroCinco >= numeroDois && numeroCinco >= numeroTres && numeroCinco >= numeroQuatro {
		fmt.Println("O maior numero é:", numeroCinco)
	}

}

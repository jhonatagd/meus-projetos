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
	totalDeEleitores := 0
	for {
		fmt.Println("Total de eleitores:")
		pegaTotal, _ := reader.ReadString('\n')
		limpaTotal := strings.TrimSpace(pegaTotal)
		totalEleitores, err := strconv.Atoi(limpaTotal)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		totalDeEleitores = totalEleitores
		break
	}

	fmt.Println("Eleitores: 1 2 ou 3")
	eleitor1 := 0
	eleitor2 := 0
	eleitor3 := 0

	for i := 0; i < totalDeEleitores; i++ {
		fmt.Println(i+1, "° eleitor, vote no seu candidato:")
		pegoVoto, _ := reader.ReadString('\n')
		limpaVoto := strings.TrimSpace(pegoVoto)
		veVoto, err := strconv.Atoi(limpaVoto)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
		}
		if veVoto == 1 {
			eleitor1 = eleitor1 + 1
		} else if veVoto == 2 {
			eleitor2 = eleitor2 + 1
		} else if veVoto == 3 {
			eleitor3 = eleitor3 + 1
		} else {
			fmt.Println("Voto não aceito")
		}
	}
	fmt.Println("Total de votos Primeiro eleitor:", eleitor1, ". Segundo eleitor:", eleitor2, ". Terceiro eleitor:", eleitor3)
}

// 3 candidatos
// pedir o numero total de eleitores
// cada eleitor votar e no final mostrar quantos votos cada canditado
//Numa eleição existem três candidatos. Faça um programa que peça o número total de eleitores.
//Peça para cada eleitor votar e ao final mostrar o número de votos de cada candidato.

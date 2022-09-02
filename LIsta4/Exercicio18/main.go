package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//a telefonista digitará um número, entre 1 e 23, correspondente ao número da camisa do jogador.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var armazenaVotosAtletas []int
	//numeroAtleta := 0
	votosComputados := 0
	for {
		fmt.Println("Digite um numero entre 1 e 23:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumeroInt := strings.TrimSpace(pegaNumero)
		pegaNumeroAtleta, err := strconv.Atoi(limpaNumeroInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else if pegaNumeroAtleta == 0 {
			break
		} else if pegaNumeroAtleta > 23 {
			fmt.Println("Numero Invalido digite novamente:")
			continue
		}
		armazenaVotosAtletas = append(armazenaVotosAtletas, pegaNumeroAtleta)
		votosComputados += 1
	}

	for i := 0; i < votosComputados; i++ {
		// if armazenaVotosAtletas[i]
	}

	fmt.Println("Total de votos computados:", votosComputados)
	fmt.Println(armazenaVotosAtletas)
}

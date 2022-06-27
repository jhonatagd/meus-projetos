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

	armazenaMenoresAcidentes := 0
	armazenaMaioresAcidentes := 0
	cidadeMenorAcidente := 0
	cidadeMaiorAcidente := 0
	divisor := 0
	mediaVeiculos := 0
	mediaTotal := 0
	for i := 0; i < 5; i++ {

		fmt.Println("Codigo da cidade:")
		pegaCodigo, _ := reader.ReadString('\n')
		limpaCodigo := strings.TrimSpace(pegaCodigo)
		codigoCidade, err := strconv.Atoi(limpaCodigo)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
		fmt.Println("Número de veículos de passeio (em 1999);")
		numeroVeiculos, _ := reader.ReadString('\n')
		limpaVeiculos := strings.TrimSpace(numeroVeiculos)
		numeroVeiculosDePasseio, err := strconv.Atoi(limpaVeiculos)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
		fmt.Println("Número de acidentes de trânsito com vítimas (em 1999);")
		numeroAcidentes, _ := reader.ReadString('\n')
		limpaAcidentes := strings.TrimSpace(numeroAcidentes)
		numeroDeAcidentes, err := strconv.Atoi(limpaAcidentes)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
		mediaVeiculos = mediaVeiculos + numeroVeiculosDePasseio

		if i == 0 {
			armazenaMenoresAcidentes = numeroDeAcidentes
			armazenaMaioresAcidentes = numeroDeAcidentes
			cidadeMenorAcidente = codigoCidade
			cidadeMaiorAcidente = codigoCidade

		} else {
			if numeroDeAcidentes < armazenaMenoresAcidentes {
				armazenaMenoresAcidentes = numeroDeAcidentes
				cidadeMenorAcidente = codigoCidade
			}
			if numeroDeAcidentes > armazenaMaioresAcidentes {
				armazenaMaioresAcidentes = numeroDeAcidentes
				cidadeMaiorAcidente = codigoCidade
			}
		}

		if numeroVeiculosDePasseio < 2.000 {
			divisor = divisor + 1
			mediaTotal = mediaTotal + numeroDeAcidentes
		}

	}
	verificaMediaDoisMil := 0
	if divisor > 0 {
		verificaMediaDoisMil = mediaTotal / divisor
	}
	fmt.Println()
	verificaMedia := mediaVeiculos / 5
	fmt.Printf("Menor índice de acidentes de transito: %v  e a que cidade pertence: %v", armazenaMenoresAcidentes, cidadeMenorAcidente)
	fmt.Printf("Maior índice de acidentes de transito: %v  e a que cidade pertence: %v", armazenaMaioresAcidentes, cidadeMaiorAcidente)
	fmt.Println("Media de veiculos de passeio:", verificaMedia)
	fmt.Println("Média de acidentes de trânsito nas cidades com menos de 2.000 veículos de passeio:", verificaMediaDoisMil)

	//Foi feita uma estatística em cinco cidades brasileiras para coletar dados sobre acidentes de trânsito.
	//Foram obtidos os seguintes dados:
	//Código da cidade;
	//Número de veículos de passeio (em 1999);
	//Número de acidentes de trânsito com vítimas (em 1999).
	// Deseja-se saber:

	//Qual o maior e menor índice de acidentes de transito e a que cidade pertence;
	//Qual a média de veículos nas cinco cidades juntas;
	//Qual a média de acidentes de trânsito nas cidades com menos de 2.000 veículos de passeio.
}

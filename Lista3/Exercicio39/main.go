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
	atualizaNumeroAluno := 0
	atualizaAlturaAluno := 0
	alunoMaisAlto := 0
	alunoMaisBaixo := 0
	armazenaNumeroAluno := 0
	for i := 0; i < 10; i++ {
		for {
			fmt.Println("Seu número como aluno:")
			pegaNumero, _ := reader.ReadString('\n')
			limpaNumero := strings.TrimSpace(pegaNumero)
			numeroAluno, err := strconv.Atoi(limpaNumero)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
			atualizaNumeroAluno = numeroAluno
			break
		}

		for {
			fmt.Println("Sua altura em centimetros:")
			pegaAltura, _ := reader.ReadString('\n')
			limpaAltura := strings.TrimSpace(pegaAltura)
			alturaAluno, err := strconv.Atoi(limpaAltura)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
			atualizaAlturaAluno = alturaAluno
			break
		}

		if atualizaNumeroAluno > armazenaNumeroAluno && atualizaAlturaAluno > alunoMaisAlto {
			armazenaNumeroAluno = armazenaNumeroAluno + atualizaNumeroAluno
			alunoMaisAlto = alunoMaisAlto + atualizaAlturaAluno
		}
		if atualizaNumeroAluno > armazenaNumeroAluno && atualizaAlturaAluno > alunoMaisBaixo {
			armazenaNumeroAluno = armazenaNumeroAluno + atualizaAlturaAluno
			alunoMaisBaixo = alunoMaisBaixo + atualizaAlturaAluno
		}
		fmt.Println("Codigo do aluno:", armazenaNumeroAluno)
		fmt.Println("Mais alto:", alunoMaisAlto)
		fmt.Println("Codigo do aluno:", armazenaNumeroAluno)
		fmt.Println("Mais baixo:", alunoMaisBaixo)
	}

	//Faça um programa que leia dez conjuntos de dois valores
	// o primeiro representando o número do aluno e o segundo representando a sua altura em centímetros.
	// Encontre o aluno mais alto e o mais baixo. Mostre o número do aluno mais alto e o número do
	// aluno mais baixo, junto com suas alturas.

}

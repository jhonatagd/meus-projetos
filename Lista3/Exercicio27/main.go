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
	quantidadeTurmas := 0

	for {
		fmt.Println("Quantidade de turmas:")
		quantidadeT, _ := reader.ReadString('\n')
		limpaQuantidadeTurmas := strings.TrimSpace(quantidadeT)
		quantidadeDeTurmas, err := strconv.Atoi(limpaQuantidadeTurmas)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		quantidadeTurmas = quantidadeDeTurmas
		break
	}
	somaTurmas := 0
	mediaDeAlunos := 0

	for i := 0; i < quantidadeTurmas; i++ {
		for {
			fmt.Println("Quantidade de alunos na", i+1, "turma:  ")
			alunos, _ := reader.ReadString('\n')
			limpaAlunos := strings.TrimSpace(alunos)
			quantidadeDeAlunosPorTurma, err := strconv.Atoi(limpaAlunos)
			if err != nil {
				fmt.Println("O que você digitou não é um numero")
				continue
			} else if quantidadeDeAlunosPorTurma > 40 {
				fmt.Println("As turmas não podem ter mais de 40 alunos.")
				i = i - 1
			} else {
				somaTurmas = somaTurmas + quantidadeDeAlunosPorTurma
				mediaDeAlunos = somaTurmas / quantidadeTurmas

			}
			break
		}
	}
	fmt.Println("Media de:", mediaDeAlunos)
}

//Faça um programa que calcule o número médio de alunos por turma.
//Para isto, peça a quantidade de turmas e a quantidade de alunos para cada turma.
//As turmas não podem ter mais de 40 alunos.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Foram anotadas as idades e alturas de 30 alunos. Faça um Programa que determine quantos alunos
//com mais de 13 anos possuem altura inferior à média de altura desses alunos.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var idade [5]int
	var altura [5]float64
	alunosAbaixoDaMedia := 0
	verificaMediaAluno := 0.0
	mediaAluno := 0.0
	for i := 0; i < 5; i++ {
		fmt.Println("Idade", i+1, "aluno:")
		pegaIdade, _ := reader.ReadString('\n')
		limpaIdadeInt := strings.TrimSpace(pegaIdade)
		idadeAluno, err := strconv.Atoi(limpaIdadeInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		idade[i] = idadeAluno

		fmt.Println("Altura", i+1, "aluno:")
		pegaAltura, _ := reader.ReadString('\n')
		limpaAlturaInt := strings.TrimSpace(pegaAltura)
		alturaAluno, err := strconv.ParseFloat(limpaAlturaInt, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		altura[i] = alturaAluno
		mediaAluno += alturaAluno

		verificaMediaAluno = mediaAluno / 10

	}
	for i := 0; i < 5; i++ {
		if idade[i] > 13 {
			if verificaMediaAluno < altura[i] {
				alunosAbaixoDaMedia += 1
			}
		}
	}
	fmt.Println("Alunos inferior a media de altura da sala:", alunosAbaixoDaMedia)
}

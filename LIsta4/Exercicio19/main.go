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

	for {
		fmt.Println("aluno:")
		pegaIdade, _ := reader.ReadString('\n')
		limpaIdadeInt := strings.TrimSpace(pegaIdade)
		idadeAluno, err := strconv.Atoi(limpaIdadeInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")

			continue
		}
	}
}

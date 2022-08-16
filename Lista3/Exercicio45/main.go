package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Desenvolver um programa para verificar a nota do aluno em uma prova com 10 questões,
//o programa deve perguntar ao aluno a resposta de cada questão e ao final comparar com o
//gabarito da prova e assim calcular o total de acertos e a nota (atribuir 1 ponto por resposta certa).
//Após cada aluno utilizar o sistema deve ser feita uma pergunta se outro aluno vai utilizar o sistema.
//Após todos os alunos terem respondido informar:
//Maior e Menor Acerto;
//Total de Alunos que utilizaram o sistema;
//A Média das Notas da Turma.

func main() {
	reader := bufio.NewReader(os.Stdin)
	test := 0
	respostaCerta := 0
	respostaErrada := 0
	quantidadeDeAlunos := 0
	tantasRespostas := 0
	maiorAcertos := 0
	menorAcertos := 0
	gabarito := []string{"A", "B", "C", "D", "E", "E", "D", "C", "B", "A"}
	for {
		for i := 0; i < 10; i++ {

			fmt.Printf("Resposta da %v° pergunta:", test+1)
			pegaNota, _ := reader.ReadString('\n')
			nota := strings.TrimSpace(pegaNota)

			if nota == gabarito[i] {
				respostaCerta = respostaCerta + 1
				maiorAcertos = maiorAcertos + 1
			} else {
				respostaErrada = respostaErrada + 1
				menorAcertos = menorAcertos + 1
			}

		}
		verificaTeste := 0
		for {
			fmt.Println("Deseja verificar mais uma vez? 1- Sim / 2- Não")
			pegaVerifica, _ := reader.ReadString('\n')
			limpaVerifica := strings.TrimSpace(pegaVerifica)
			verifica, err := strconv.Atoi(limpaVerifica)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
			verificaTeste = verifica
			break

		}
		if verificaTeste == 1 {
			quantidadeDeAlunos = quantidadeDeAlunos + 1
			tantasRespostas = quantidadeDeAlunos + 10
			continue
		}
		media := tantasRespostas / quantidadeDeAlunos
		fmt.Println("Maior acerto:", maiorAcertos)
		fmt.Println("Menor acerto:", menorAcertos)
		fmt.Println("Total de Alunos que utilizaram o sistema:", quantidadeDeAlunos)
		fmt.Println("A Média das Notas da Turma:", media)
		break
	}
}

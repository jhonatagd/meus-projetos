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
		if maiorAcertos > 

		fmt.Println("Deseja verificar mais uma vez? 1- Sim / 2- Não")
			pegaVerifica, _ := reader.ReadString('\n')
			limpaVerifica := strings.TrimSpace(pegaVerifica)
			verifica, err := strconv.Atoi(limpaVerifica)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
		if verifica == 1 {
			quantidadeDeAlunos = quantidadeDeAlunos + 1
			tantasRespostas = quantidadeDeAlunos + 10
			continue
		}
		media := tantasRespostas / quantidadeDeAlunos
		fmt.Println("Maior e Menor Acerto;")
		fmt.Println("Total de Alunos que utilizaram o sistema:", quantidadeDeAlunos)
		fmt.Println("A Média das Notas da Turma:", media)

	}
}

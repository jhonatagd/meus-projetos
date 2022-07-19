package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	test := 0
	respostaCerta := 0
	respostaErrada := 0
	quantidadeDeAlunos := 0
	tantasRespostas := 0
	gabarito := []string{"A", "B", "C", "D", "E", "E", "D", "C", "B", "A"}
	for {
		for i := 0; i < 10; i++ {

			fmt.Printf("Resposta da %v° pergunta:", test+1)
			pegaNota, _ := reader.ReadString('\n')
			nota := strings.TrimSpace(pegaNota)

			if nota == gabarito[i] {
				respostaCerta = respostaCerta + 1
			} else {
				respostaErrada = respostaErrada + 1
			}

		}

		fmt.Printf("Deseja verificar as respostas de novo? (S para sim e N para não)")
		pegaVerifica, _ := reader.ReadString('\n')
		verifica := strings.TrimSpace(pegaVerifica)

		if verifica == "S" {
			quantidadeDeAlunos = quantidadeDeAlunos + 1
			tantasRespostas := quantidadeDeAlunos + 10
			continue
		}
		media := respostaCerta / re

	}
}

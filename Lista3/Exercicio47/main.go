package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Em uma competição de ginástica, cada atleta recebe votos de sete jurados.
//A melhor e a pior nota são eliminadas. A sua nota fica sendo a média dos votos restantes.
//Você deve fazer um programa que receba o nome do ginasta e as notas dos sete jurados
//alcançadas pelo atleta em sua apresentação e depois informe a sua média, conforme a
//descrição acima informada *****(retirar o melhor e o pior salto e depois calcular a média
//com as notas restantes).**** As notas não são informados ordenadas. Um exemplo de saída
//do programa deve ser conforme o exemplo abaixo:
func main() {
	reader := bufio.NewReader(os.Stdin)
	nomeDoAtleta := "0"
	for {
		fmt.Println("Nome do atleta:")
		pegNome, _ := reader.ReadString('\n')
		nomeAluno, _ := strings.TrimSpace(pegNome)
		if err != nil {
			fmt.Println("O que você digitou não é um nome")
			continue
		}
		nomeDoAtleta = numeroAluno
	}

	for i := 0; i < 7; i++ {
		pegaNota := 0
		contaNota := 0
		maiorNota := 0
		menorNota := 0
		for {
			fmt.Println(contaNota+1, "° nota:")
			pegaNota, _ := reader.ReadString('\n')
			limpaNota := strings.TrimSpace(pegaNota)
			numeroAluno, err := strconv.Atoi(limpaNota)
			if err != nil {
				fmt.Println("O que você digitou não é uma nota")
				continue
			}
			pegaNota = numeroAluno
		}

	}
}

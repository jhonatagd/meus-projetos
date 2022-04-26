//Faça um Programa para leitura de três notas parciais de um aluno. O programa deve calcular a média alcançada por aluno e presentar:
//A mensagem "Aprovado", se a média for maior ou igual a 7, com a respectiva média alcançada;
//A mensagem "Reprovado", se a média for menor do que 7, com a respectiva média alcançada;
//A mensagem "Aprovado com Distinção", se a média for igual a 10.

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
	fmt.Println("Valor da primeira nota:")
	primeira, _ := reader.ReadString('\n')
	limpaPrimeira := strings.TrimSpace(primeira)
	primeiraInt, err := strconv.Atoi(limpaPrimeira)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	fmt.Println("Valor da segunda nota:")
	segunda, _ := reader.ReadString('\n')
	limpaSegunda := strings.TrimSpace(segunda)
	segundaInt, err := strconv.Atoi(limpaSegunda)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	fmt.Println("Valor da terceira nota:")
	terceira, _ := reader.ReadString('\n')
	limpaTerceira := strings.TrimSpace(terceira)
	terceiraInt, err := strconv.Atoi(limpaTerceira)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}

	media := (primeiraInt + segundaInt + terceiraInt) / 3

	if media == 10 {
		fmt.Println("Aprovado com Distinção, media de:", media)
	} else if media >= 7 {
		fmt.Println("Aprovado, media de:", media)
	} else if media < 7 {
		fmt.Println("Reprovado, media de:", media)
	}
}

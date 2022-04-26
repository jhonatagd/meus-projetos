package main

// leitura de duas notas parciais de um aluno
// deve calcular a média alcançada por aluno e apresentar:
// "Aprovado", se a média alcançada for maior ou igual a sete
// "Reprovado", se a média for menor do que sete
//  "Aprovado com Distinção", se a média for igual a dez

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Primeira nota: ")
	notaUm, _ := reader.ReadString('\n')
	limpaNotaUm := strings.TrimSpace(notaUm)
	notaUmInt, err := strconv.Atoi(limpaNotaUm)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	fmt.Println("Segunda nota: ")
	notaDois, _ := reader.ReadString('\n')
	limpaNotaDois := strings.TrimSpace(notaDois)
	notaDoisInt, err := strconv.Atoi(limpaNotaDois)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	somaNota := notaUmInt + notaDoisInt
	calculaMedia := somaNota / 2

	if calculaMedia == 10 {
		fmt.Println(calculaMedia, "Aprovado com distinção.")
	} else if calculaMedia < 7 {
		fmt.Println(calculaMedia, "Reprovado")
	} else if calculaMedia >= 7 {
		fmt.Println(calculaMedia, "Aprovado")
	}
}

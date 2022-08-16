package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um Programa que peça as quatro notas de 10 alunos, calcule e armazene
//num vetor a média de cada aluno, imprima o número de alunos com média maior ou igual a 7.0.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var mediaAlunos [2]float64
	//var acimaDaMedia []float64
	contaACimaDaMedia := 0.0

	for i := 0; i < 2; i++ {
		media := 0.0
		for j := 0; j < 4; j++ {
			fmt.Println(j+1, "Nota:")
			pegaNota, _ := reader.ReadString('\n')
			limpaNota := strings.TrimSpace(pegaNota)
			nota, err := strconv.ParseFloat(limpaNota, 64)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				j--
				continue
			} else if nota > 10 || nota < 0 {
				fmt.Println("O que você digitou não é um número")
				j--
				continue
			}
			media += nota / 4
			mediaAlunos[i] = float64(media)
		}

		if media >= 7 {
			contaACimaDaMedia += 1
		}

	}

	fmt.Println("Media dos alunos", mediaAlunos)
	fmt.Println("Alunos com a media superior a 7", contaACimaDaMedia)
}

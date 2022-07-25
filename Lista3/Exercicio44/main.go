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
	votosJose := 0
	votosJoao := 0
	votosMarcos := 0
	votosJhona := 0
	votosNulo := 0
	votosBranco := 0
	atualizaVoto := 0
	for {

		for {
			fmt.Println("1 , 2, 3, 4  - Votos para os respectivos candidatos ")
			fmt.Println("1- Jose/ 2- João/ 3- Marcos/ 4- Jhona/ 5- Voto nulo/ 6- Voto em branco")
			pegaVoto, _ := reader.ReadString('\n')
			limpaVoto := strings.TrimSpace(pegaVoto)
			voto, err := strconv.Atoi(limpaVoto)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
			atualizaVoto = voto
			break
		}

		if atualizaVoto == 1 {
			votosJose = votosJose + 1
		}
		if atualizaVoto == 2 {
			votosJoao = votosJoao + 1
		}
		if atualizaVoto == 3 {
			votosMarcos = votosMarcos + 1
		}
		if atualizaVoto == 4 {
			votosJhona = votosJhona + 1
		}
		if atualizaVoto == 5 {
			votosNulo = votosNulo + 1
		}
		if atualizaVoto == 6 {
			votosBranco = votosBranco + 1
		}

		if atualizaVoto == 0 {
			fmt.Printf("O total de votos para cada candidato: Jose %v, João %v, Marcos %v, Jhona %v", votosJose, votosJoao, votosMarcos, votosJhona)
			fmt.Println("O total de votos nulos:", votosNulo)
			fmt.Println("O total de votos em branco:", votosBranco)
			fmt.Println("A percentagem de votos nulos sobre o total de votos:")
			fmt.Println("A percentagem de votos em branco sobre o total de votos:")
		}
	}

}

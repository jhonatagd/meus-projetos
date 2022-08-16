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
		pegaNome, err := reader.ReadString('\n')
		pegaNomeString := strings.TrimSpace(pegaNome)
		pegaNomeString = strings.ToUpper(pegaNomeString)
		if err != nil {
			fmt.Println("O que você digitou não é uma letra")
		}
		nomeDoAtleta = pegaNomeString
		break
	}
	var listaFloat [7]float64
	for i := 0; i < 7; i++ {
		for {
			fmt.Println("Nota do", i+1, "jurado:")
			pegaNota, _ := reader.ReadString('\n')
			limpaNota := strings.TrimSpace(pegaNota)
			pegaNotaJurado, err := strconv.ParseFloat(limpaNota, 64)
			if err != nil {
				fmt.Println("O que você digitou não é uma nota")
				i = i - 1
				continue
			}
			listaFloat[i] = pegaNotaJurado
			break
		}
	}
	fmt.Println("Atleta:", nomeDoAtleta)

	for i := 0; i < 7; i++ {
		fmt.Println(i+1, "° Nota:", listaFloat[i])
	}

	melhorSalto := 0.0
	piorSalto := 0.0
	for i := 0; i < 7; i++ {
		if i == 0 {
			melhorSalto = listaFloat[i]
			piorSalto = listaFloat[i]
		} else {
			if listaFloat[i] > melhorSalto {
				melhorSalto = listaFloat[i]
			}
			if listaFloat[i] < piorSalto {
				piorSalto = listaFloat[i]
			}
		}
	}

	fmt.Println("Pior salto:", piorSalto)
	fmt.Println("Melhor Salto:", melhorSalto)

	totalSaltos := 0.0
	for i := 0; i < 7; i++ {
		totalSaltos = totalSaltos + listaFloat[i]
	}

	media := (totalSaltos - (melhorSalto + piorSalto)) / 5

	fmt.Println("Média dos demais saltos:", media)

}

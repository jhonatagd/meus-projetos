package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Em uma competição de salto em distância cada atleta tem direito a cinco saltos.
//No final da série de saltos de cada atleta, o melhor e o pior resultados são eliminados.
//O seu resultado fica sendo a média dos três valores restantes. Você deve fazer um
//programa que receba o nome e as cinco distâncias alcançadas pelo atleta em seus saltos
//e depois informe a média dos saltos conforme a descrição acima informada
//(retirar o melhor e o pior salto e depois calcular a média).
//Faça uso de uma lista para armazenar os saltos. Os saltos são informados
//na ordem da execução, portanto não são ordenados. O programa deve ser encerrado
//quando não for informado o nome do atleta. A saída do programa deve ser conforme
//o exemplo abaixo:
func main() {
	reader := bufio.NewReader(os.Stdin)
	nomeAtleta := "0"
	for {
		fmt.Println("Nome do atleta:")
		pegaLetra, err := reader.ReadString('\n')
		pegaLetraString := strings.TrimSpace(pegaLetra)
		pegaLetraString = strings.ToUpper(pegaLetraString)
		if err != nil {
			fmt.Println("O que você digitou não é uma letra")
			continue
		}
		nomeAtleta = pegaLetraString
		break
	}

	var listaFloat [5]float64

	for i := 0; i < 5; i++ { // {}, {}, {}, {}{}, {}

		for {
			fmt.Println("Distancia do", i+1, "salto:")
			pegaSalto, _ := reader.ReadString('\n')
			limpaSalto := strings.TrimSpace(pegaSalto)
			distanciaSalto, err := strconv.ParseFloat(limpaSalto, 64)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				i = i - 1
				continue
			}
			listaFloat[i] = distanciaSalto
			break
		}

	}
	fmt.Println("Atleta:", nomeAtleta)

	for i := 0; i < 5; i++ {
		fmt.Println(i+1, "° salto", listaFloat[i])
	}

	melhorSalto := 0.0
	piorSalto := 0.0
	for i := 0; i < 5; i++ {
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
	for i := 0; i < 5; i++ {
		totalSaltos = totalSaltos + listaFloat[i]

	}
	media := (totalSaltos - (melhorSalto + piorSalto)) / 3

	fmt.Println("Média dos demais saltos:", media)

}

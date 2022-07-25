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
		pegaNome, _ := reader.ReadString('\n')
		limpaNomes := strings.TrimSpace(pegaNome)
		PegaNomeAtleta, err := strconv.Atoi(limpaNomes)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		nomeAtleta = PegaNomeAtleta
	}

	mediaDeSalto := 0
	maiorSalto := 0
	menorSalto := 0

	for i := 0; i < 5; i++ {
		contaSalto := 0
		distanciaSaltos := 0
		for {
			fmt.Printf("Distancia do v% salto:", contaSalto+1)
			pegaSalto, _ := reader.ReadString('\n')
			limpaSalto := strings.TrimSpace(pegaSalto)
			distanciaSalto, err := strconv.Atoi(limpaSalto)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
			distanciaSaltos = distanciaSalto
		}
		if distanciaSaltos > maiorSalto {
			maiorSalto = distanciaSaltos
		}

	}
}

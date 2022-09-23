package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um programa que carregue uma lista com os modelos de cinco carros (exemplo de modelos: FUSCA, GOL, VECTRA etc).
//Carregue uma outra lista com o consumo desses carros
// isto é, quantos quilômetros cada um desses carros faz com um litro de combustível.
//Calcule e mostre:
//O modelo do carro mais econômico;

func main() {
	reader := bufio.NewReader(os.Stdin)
	var modelosCarros [4]string
	var consumoDeCadaCarro [4]int
	for i := 0; i < 5; i++ {

		fmt.Println(i+1, "° modelo de carro:")
		modeloCarro, err := reader.ReadString('\n')
		pegaModeloCarro := strings.TrimSpace(modeloCarro)
		pegaModeloCarro = strings.ToUpper(pegaModeloCarro)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		modelosCarros[i] = pegaModeloCarro
	}

	for j := 0; j < 5; j++ {

		fmt.Println("consumo desse carro por km:", modelosCarros[j])
		consumoCarro, _ := reader.ReadString('\n')
		pegaConsumoCarro := strings.TrimSpace(consumoCarro)
		consumoDoCarro, err := strconv.Atoi(pegaConsumoCarro)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			j--
			continue
		}
		consumoDeCadaCarro[j] = consumoDoCarro
	}

	carroMaisEconomico := 0
	for a := 0; a < 5; a++ {
		if carroMaisEconomico < consumoDeCadaCarro[a] {
			carroMaisEconomico = consumoDeCadaCarro[a]
		}
	}

	//	for b := 0; b < 5; b++ {
	//		if consumoDeCadaCarro[i] == carroMaisEconomico {
	//			carroMaisEconomico = consumoDeCadaCarro[a]
	//		}
	//	}

}

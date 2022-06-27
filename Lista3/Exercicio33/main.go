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

	menorTemperatura := 0
	maiorTemperatura := 0
	somaTemperaturas := 0
	quantasVezesPercorreu := 0
	//desenvolver um programa que leia as um conjunto indeterminado de temperaturas
	//informe ao final a menor e a maior temperaturas informadas, bem como a média das temperaturas.
	for {
		mediaTemperatura := somaTemperaturas / quantasVezesPercorreu

		fmt.Println("Informe a temperatura (X - para encerrar):")
		pegaTemperatura, _ := reader.ReadString('\n')
		limpaTemperatura := strings.TrimSpace(pegaTemperatura)
		if strings.ToLower(limpaTemperatura) == "x" {
			fmt.Println("Maior temperatura:", maiorTemperatura)
			fmt.Println("Menor temperatura:", menorTemperatura)
			fmt.Println("Media das temperatura:", mediaTemperatura)
			break
		}
		temperatura, err := strconv.Atoi(limpaTemperatura)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
		}
		quantasVezesPercorreu = quantasVezesPercorreu + 1
		somaTemperaturas = somaTemperaturas + temperatura

		if menorTemperatura < temperatura {
			menorTemperatura = menorTemperatura + temperatura
		} else if maiorTemperatura > temperatura {
			maiorTemperatura = maiorTemperatura + temperatura
		}

	}

}

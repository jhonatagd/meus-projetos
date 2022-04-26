package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Responda com Sim ou Não!")
	fmt.Println("Telefonou para a vítima?")
	primeiraPergunta, err := reader.ReadString('\n')
	pegaPrimeiraPergunta := strings.TrimSpace(primeiraPergunta)
	pegaPrimeiraPergunta = strings.ToUpper(pegaPrimeiraPergunta)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}

	fmt.Println("Esteve no local do crime?")
	segundaPergunta, err := reader.ReadString('\n')
	pegaSegundaPergunta := strings.TrimSpace(segundaPergunta)
	pegaSegundaPergunta = strings.ToUpper(pegaSegundaPergunta)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}

	fmt.Println("Mora perto da vítima?")
	terceiraPergunta, err := reader.ReadString('\n')
	pegaTerceiraPergunta := strings.TrimSpace(terceiraPergunta)
	pegaTerceiraPergunta = strings.ToUpper(pegaTerceiraPergunta)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}
	fmt.Println("Devia para a vítima?")
	quartaPergunta, err := reader.ReadString('\n')
	pegaQuartaPergunta := strings.TrimSpace(quartaPergunta)
	pegaQuartaPergunta = strings.ToUpper(pegaQuartaPergunta)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}
	fmt.Println("Já trabalhou com a vítima?")
	quintaPergunta, err := reader.ReadString('\n')
	pegaQuintaPergunta := strings.TrimSpace(quintaPergunta)
	pegaQuintaPergunta = strings.ToUpper(pegaQuintaPergunta)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}
	totalSim := 0
	totalNao := 0
	if pegaPrimeiraPergunta == "SIM" {
		totalSim = totalSim + 1
	} else if pegaPrimeiraPergunta != "SIM" {
		totalNao = totalNao + 1
	} else {
		fmt.Println("Responde direito Sim ou Não")
	}

	if pegaSegundaPergunta == "SIM" {
		totalSim = totalSim + 1
	} else if pegaSegundaPergunta != "SIM" {
		totalNao = totalNao + 1
	} else {
		fmt.Println("Responde direito Sim ou Não")
	}

	if pegaTerceiraPergunta == "SIM" {
		totalSim = totalSim + 1
	} else if pegaTerceiraPergunta != "SIM" {
		totalNao = totalNao + 1
	} else {
		fmt.Println("Responde direito Sim ou Não")
	}

	if pegaQuartaPergunta == "SIM" {
		totalSim = totalSim + 1
	} else if pegaQuartaPergunta != "SIM" {
		totalNao = totalNao + 1
	} else {
		fmt.Println("Responde direito Sim ou Não")
	}

	if pegaQuintaPergunta == "SIM" {
		totalSim = totalSim + 1
	} else if pegaQuintaPergunta != "SIM" {
		totalNao = totalNao + 1
	} else {
		fmt.Println("Responde direito Sim ou Não")
	}

	if totalSim == 2 {
		fmt.Println("Suspeita")
	} else if totalSim == 3 || totalSim == 4 {
		fmt.Println("Cúmplice")
	} else if totalSim == 5 {
		fmt.Println("Assasino")
	} else {
		fmt.Println("Inocente")
	}
	// Se a pessoa responder positivamente a 2 questões ela deve ser classificada como "Suspeita"
	// entre 3 e 4 como "Cúmplice" e 5 como "Assassino". Caso contrário, ele será classificado como "Inocente"
}

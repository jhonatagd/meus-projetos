package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Utilizando listas faça um programa que faça 5 perguntas para uma pessoa sobre um crime. As perguntas são:
//"Telefonou para a vítima?"
//"Esteve no local do crime?"
//"Mora perto da vítima?"
//"Devia para a vítima?"
//"Já trabalhou com a vítima?" O programa deve no final emitir uma classificação sobre a participação da pessoa no crime.
//Se a pessoa responder positivamente a 2 questões ela deve ser classificada como "Suspeita", entre 3 e 4 como
//"Cúmplice" e 5 como "Assassino". Caso contrário, ele será classificado como "Inocente".
func main() {
	reader := bufio.NewReader(os.Stdin)
	resposta1 := " "
	resposta2 := " "
	resposta3 := " "
	resposta4 := " "
	resposta5 := " "
	for {
		fmt.Println("Telefonou para a vítima?")
		pegaPergunta1, err := reader.ReadString('\n')
		pegaPergunta1String := strings.TrimSpace(pegaPergunta1)
		pegaPergunta1String = strings.ToUpper(pegaPergunta1String)
		if err != nil {
			fmt.Println("O que você digitou não é uma resposta")
			continue
		}
		resposta1 = pegaPergunta1String
		break
	}

	for {
		fmt.Println("Esteve no local do crime?")
		pegaPergunta2, err := reader.ReadString('\n')
		pegaPergunta2String := strings.TrimSpace(pegaPergunta2)
		pegaPergunta2String = strings.ToUpper(pegaPergunta2String)
		if err != nil {
			fmt.Println("O que você digitou não é uma resposta")
			continue
		}
		resposta2 = pegaPergunta2String
		break
	}

	for {
		fmt.Println("Mora perto da vítima?")
		pegaPergunta3, err := reader.ReadString('\n')
		pegaPergunta3String := strings.TrimSpace(pegaPergunta3)
		pegaPergunta3String = strings.ToUpper(pegaPergunta3String)
		if err != nil {
			fmt.Println("O que você digitou não é uma resposta")
			continue
		}
		resposta3 = pegaPergunta3String
		break
	}
	for {
		fmt.Println("Devia para a vítima?")
		pegaPergunta4, err := reader.ReadString('\n')
		pegaPergunta4String := strings.TrimSpace(pegaPergunta4)
		pegaPergunta4String = strings.ToUpper(pegaPergunta4String)
		if err != nil {
			fmt.Println("O que você digitou não é uma resposta")
			continue
		}
		resposta4 = pegaPergunta4String
		break
	}
	for {
		fmt.Println("Já trabalhou com a vítima?")
		pegaPergunta5, err := reader.ReadString('\n')
		pegaPergunta5String := strings.TrimSpace(pegaPergunta5)
		pegaPergunta5String = strings.ToUpper(pegaPergunta5String)
		if err != nil {
			fmt.Println("O que você digitou não é uma resposta")
			continue
		}
		pegaPergunta5 = pegaPergunta5String
		break
	}
	totalRespostasSim := 0
	if resposta1 == "SIM" {
		totalRespostasSim = totalRespostasSim + 1
	} else if resposta2 == "SIM" {
		totalRespostasSim = totalRespostasSim + 1
	} else if resposta3 == "SIM" {
		totalRespostasSim = totalRespostasSim + 1
	} else if resposta4 == "SIM" {
		totalRespostasSim = totalRespostasSim + 1
	} else if resposta5 == "SIM" {
		totalRespostasSim = totalRespostasSim + 1
	}

	if totalRespostasSim < 2 {
		fmt.Println("Suspeita")
	}
	if totalRespostasSim > 2 && totalRespostasSim < 4 {
		fmt.Println("Cúmplice")
	}
	if totalRespostasSim == 5 {
		fmt.Println("Assassino")
	}
	fmt.Println(totalRespostasSim)
	//Se a pessoa responder positivamente a 2 questões ela deve ser classificada como "Suspeita", entre 3 e 4 como
	//"Cúmplice" e 5 como "Assassino". Caso contrário, ele será classificado como "Inocente".
}

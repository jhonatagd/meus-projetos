package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Errinhos: Programa esta terminando, mesmo que o valor digitado seja inválido
	//Sugestão: Um for para cada valor pedido, terminando o for somente quando o valor digitado for valido
	reader := bufio.NewReader(os.Stdin)
	nome := "0"
	idade := 0
	salario := 0
	sexo := "0"

	for {
		fmt.Println("Nome:")
		um, err := reader.ReadString('\n')
		pegaNome := strings.TrimSpace(um)
		pegaNome = strings.ToUpper(pegaNome)
		_, err = strconv.Atoi(pegaNome)
		if err == nil {
			fmt.Println("Nome incorreto, digite novamente:")
			continue
		} else {

			nomeLen := len(nome)

			if nomeLen > 3 {
				nome = pegaNome
				break
			} else {
				fmt.Println("Nome: maior que 3 caracteres")
				continue
			}

		}

	}

	fmt.Println("Idade:")
	dois, _ := reader.ReadString('\n')
	dois = strings.TrimSpace(dois)
	Idade, err := strconv.Atoi(dois)
	if err != nil {
		fmt.Println("O que você digitou não é uma idade")
	}
	fmt.Println("Salario:")
	tres, _ := reader.ReadString('\n')
	tres = strings.TrimSpace(tres)
	Salario, err := strconv.Atoi(tres)
	if err != nil {
		fmt.Println("O que você digitou não é uma idade")
	}
	fmt.Println("Sexo: 'f' ou 'm'")
	quatro, err := reader.ReadString('\n')
	sexo := strings.TrimSpace(quatro)
	sexo = strings.ToUpper(sexo)
	if err != nil {
		fmt.Println("O que você digitou não é um nome")
	}
	fmt.Println("Estado Civil:'s', 'c', 'v', 'd'")
	cinco, err := reader.ReadString('\n')
	estadoCivil := strings.TrimSpace(cinco)
	estadoCivil = strings.ToUpper(estadoCivil)
	if err != nil {
		fmt.Println("O que você digitou não é um nome")
	}

	if Idade >= 0 && Idade <= 150 {
		fmt.Println("Idade ok")
	} else {
		fmt.Println("Idade: entre 0 e 150")
	}
	if Salario > 0 {
		fmt.Println("Salario ok")
	} else {
		fmt.Println("Salário: maior que zero")
	}
	if sexo == "F" || sexo == "M" {
		fmt.Println("Sexo ok")
	} else {
		fmt.Println("Sexo: 'f' ou 'm'")
	}
	if (estadoCivil == "S") || (estadoCivil == "C") || (estadoCivil == "V") || (estadoCivil == "D") {
		fmt.Println("Estado Civil ok")
	} else {
		fmt.Println("Estado Civil: 's', 'c', 'v' ou 'd'")
	}

	//Faça um programa que leia e valide as seguintes informações:
	//Nome: maior que 3 caracteres;
	//Idade: entre 0 e 150;
	//Salário: maior que zero;
	//Sexo: 'f' ou 'm';
	//Estado Civil: 's', 'c', 'v', 'd';
}

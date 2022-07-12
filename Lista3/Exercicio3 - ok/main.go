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

	for {
		fmt.Println("Nome:")
		armazenaNome, err := reader.ReadString('\n')
		pegaNome := strings.TrimSpace(armazenaNome)
		pegaNome = strings.ToUpper(pegaNome)
		_, err = strconv.Atoi(pegaNome)
		if err == nil {
			fmt.Println("Nome incorreto, digite novamente:")
			continue
		} else {

			nomeLen := len(pegaNome)

			if nomeLen > 3 {
				//nome = pegaNome
				break
			} else {
				fmt.Println("Nome: maior que 3 caracteres")
				continue
			}
		}
	}

	for {
		fmt.Println("Idade:")
		ArmazenaIdade, _ := reader.ReadString('\n')
		ArmazenaIdade = strings.TrimSpace(ArmazenaIdade)
		pegaIdade, err := strconv.Atoi(ArmazenaIdade)
		if err != nil {
			fmt.Println("idade incorreta, digite novamente:")
			continue
		} else {

			if pegaIdade > 0 || pegaIdade < 150 {
				// idade = idade + pegaIdade
				break
			} else {
				fmt.Println("Idade: entre 0 e 150")
				continue
			}
		}
	}

	for {
		fmt.Println("Salario:")
		armazenaSalario, _ := reader.ReadString('\n')
		armazenaSalario = strings.TrimSpace(armazenaSalario)
		pegaSalario, err := strconv.Atoi(armazenaSalario)
		if err != nil {
			fmt.Println("Salario incorreto, digite novamente:")
			continue
		} else {

			if pegaSalario > 0 {
				//	salario = salario + pegaSalario
				break
			} else {
				fmt.Println("Salário: maior que zero")
				continue
			}
		}
	}

	for {
		fmt.Println("Sexo: 'f' ou 'm'")
		armazenaSexo, err := reader.ReadString('\n')
		pegaSexo := strings.TrimSpace(armazenaSexo)
		pegaSexo = strings.ToUpper(pegaSexo)
		if err != nil {
			fmt.Println("incorreto, Digite o sexo novamente")
			continue
		} else {

			if pegaSexo == "F" || pegaSexo == "M" {
				//	sexo = sexo + pegaSexo
				break
			} else {
				fmt.Println("Sexo: 'f' ou 'm':")
				continue
			}
		}
	}

	for {
		fmt.Println("Estado Civil:'s', 'c', 'v', 'd'")
		armazenaEstadoCivil, err := reader.ReadString('\n')
		estadoCivil := strings.TrimSpace(armazenaEstadoCivil)
		estadoCivil = strings.ToUpper(estadoCivil)
		if err != nil {
			fmt.Println("incorreto, Digite novamente")
			continue
		} else {

			if (estadoCivil == "S") || (estadoCivil == "C") || (estadoCivil == "V") || (estadoCivil == "D") {
				break
			} else {
				fmt.Println("Estado Civil: 's', 'c', 'v' ou 'd'")
				continue
			}
		}
	}
	fmt.Println("Nome: ok")
	fmt.Println("Idade: ok")
	fmt.Println("Salario: ok")
	fmt.Println("Sexo: ok")
	fmt.Println("Estado civil: ok")

	//Faça um programa que leia e valide as seguintes informações:
	//Nome: maior que 3 caracteres;
	//Idade: entre 0 e 150;
	//Salário: maior que zero;
	//Sexo: 'f' ou 'm';
	//Estado Civil: 's', 'c', 'v', 'd';
}

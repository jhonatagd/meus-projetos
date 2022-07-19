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
	// faça a tabuada de um número qualquer inteiro que será digitado pelo usuário
	//Montar a tabuada de: 5
	//Começar por: 4
	//Terminar em: 7
	//Vou montar a tabuada de 5 começando em 4 e terminando em 7:
	// 5 X 4 = 20
	tabuadado1 := 0
	comecaPor := 0
	terminaEm := 0
	for {
		fmt.Println("Tabuada do:")
		pegaTabuada, _ := reader.ReadString('\n')
		limpaTabuada := strings.TrimSpace(pegaTabuada)
		tabuada, err := strconv.Atoi(limpaTabuada)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		tabuadado1 = tabuada
		break
	}

	for {
		fmt.Println("Começa pelo:")
		pegaComeco, _ := reader.ReadString('\n')
		limpaComeco := strings.TrimSpace(pegaComeco)
		comeco, err := strconv.Atoi(limpaComeco)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		comecaPor = comeco
		break
	}

	for {
		fmt.Println("Termina em:")
		pegaFim, _ := reader.ReadString('\n')
		limpaFim := strings.TrimSpace(pegaFim)
		termina, err := strconv.Atoi(limpaFim)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else if termina > comecaPor {
			fmt.Println("Final maior que o começo.")
			break
		}
		terminaEm = termina
		break
	}

	for i := comecaPor; i <= terminaEm; i++ {

		fmt.Println(tabuadado1, " X ", i, "=", tabuadado1*i)

	}

}

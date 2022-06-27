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
	fmt.Println("Tabuada do:")
	pegaTabuada, _ := reader.ReadString('\n')
	limpaTabuada := strings.TrimSpace(pegaTabuada)
	tabuada, err := strconv.Atoi(limpaTabuada)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Começa pelo:")
	pegaComeco, _ := reader.ReadString('\n')
	limpaComeco := strings.TrimSpace(pegaComeco)
	comeco, err := strconv.Atoi(limpaComeco)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	fmt.Println("Termina em:")
	pegaFim, _ := reader.ReadString('\n')
	limpaFim := strings.TrimSpace(pegaFim)
	termina, err := strconv.Atoi(limpaFim)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	} else if termina > comeco {
		fmt.Println("Final maior que o começo.")
	}
	for i := comeco; i <= termina; i++ {

		fmt.Println(tabuada, " X ", i, "=", tabuada*i)

	}

}

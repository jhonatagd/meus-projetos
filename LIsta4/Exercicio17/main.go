package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fazer um programa que receba o nome e as cinco distâncias alcançadas pelo atleta em seus
//saltos e depois informe o nome, os saltos e a média dos saltos.
func main() {
	reader := bufio.NewReader(os.Stdin)
	nomeAtleta := " "
	for {
		fmt.Println("Nome do atleta:")
		pegaCaractere, err := reader.ReadString('\n')
		pegaCaractereStrings := strings.TrimSpace(pegaCaractere)
		pegaCaractereStrings = strings.ToUpper(pegaCaractereStrings)
		if err != nil {
			fmt.Println("O que você digitou não é um nome")
			continue
		}
		nomeAtleta = pegaCaractereStrings
		break
	}
	var listaSaltos [5]float64
	contaSalto := 0
	for i := 0; i < 5; i++ {

		fmt.Println(contaSalto+1, "Salto:")
		pegaSalto, _ := reader.ReadString('\n')
		limpaSaltoInt := strings.TrimSpace(pegaSalto)
		saltoAtleta, err := strconv.ParseFloat(limpaSaltoInt, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		listaSaltos[i] = saltoAtleta
	}
	somaSalto := 0.0
	for i := 0; i < 5; i++ {
		somaSalto += listaSaltos[i]
	}

	fmt.Println("Atleta:", nomeAtleta)
	fmt.Println("Primeiro salto:", listaSaltos[0])
	fmt.Println("Segundo salto:", listaSaltos[1])
	fmt.Println("Terceiro salto:", listaSaltos[2])
	fmt.Println("Quarto salto:", listaSaltos[3])
	fmt.Println("Quinto salto:", listaSaltos[4])
	fmt.Println("Resultado final:")
	fmt.Println("Atleta:", nomeAtleta)
	fmt.Println("Média dos saltos", somaSalto/5)

}

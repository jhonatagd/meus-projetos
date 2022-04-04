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
	fmt.Print("Digite a nota do primeiro bimestre: ")
	NotaBiUm, _ := reader.ReadString('\n')       // Capturando a imput e jogando em uma variável após apertar "enter"
	NotaBiUmLimpa := strings.TrimSpace(NotaBiUm) // Limpando a variável
	NotaBiUmint, err := strconv.ParseFloat(NotaBiUmLimpa, 64)

	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	fmt.Print("Digite a nota do segundo bimestre: ")
	NotaBiDois, _ := reader.ReadString('\n')
	NotaBiDoisLimpa, err := strconv.ParseFloat(strings.TrimSpace(NotaBiDois), 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	fmt.Print("Digite a nota do terceiro bimestre: ")
	NotaBitres, _ := reader.ReadString('\n')
	NotaBitresLimpa, err := strconv.ParseFloat(strings.TrimSpace(NotaBitres), 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	fmt.Print("Digite a nota do quarto bimestre: ")
	NotaBiquatro, _ := reader.ReadString('\n')
	NotaBiquatroLimpa, err := strconv.ParseFloat(strings.TrimSpace(NotaBiquatro), 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma nota")
	}
	SomaNotasBimestre := strconv.FormatFloat((NotaBiUmint+NotaBiDoisLimpa+NotaBitresLimpa+NotaBiquatroLimpa)/4, 'f', -1, 64)
	fmt.Println(" Media das notas = " + SomaNotasBimestre)
}

// diminuir o codigo "reader"
// extrair calculo da linha 39 e colocar em uma variavel

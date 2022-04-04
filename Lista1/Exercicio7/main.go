package main

// Faça um Programa que calcule a área de um quadrado, em seguida mostre o dobro desta área para o usuário.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" area de um quadrado: ")
	areaDoQuadrado, _ := reader.ReadString('\n')
	LimpaAreaDoQuadrado := strings.TrimSpace(areaDoQuadrado)
	areaDoQuadradoFloat, err := strconv.ParseFloat(LimpaAreaDoQuadrado, 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma area ")
	}
	calcularAreaQuadrado := areaDoQuadradoFloat * areaDoQuadradoFloat
	dobroDaArea := calcularAreaQuadrado * 2

	fmt.Println("O dobro da sua área é: ")
	fmt.Println(dobroDaArea)
}

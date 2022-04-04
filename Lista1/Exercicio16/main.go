package main

// pedir o tamanho em metros quadrados da área a ser pintada.
// cobertura da tinta é de 1 litro para cada 3 metros quadrados ( / 3 ? )
// tinta é vendida em latas de 18 litros, que custam R$ 80,00.
// Informe ao usuário a quantidades de latas de tinta a serem compradas e o preço total.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tamanho em metros quadrados da área a ser pintada: ")
	metros, _ := reader.ReadString('\n')
	limpaMetros := strings.TrimSpace(metros)
	metroFloat, err := strconv.ParseFloat(limpaMetros, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um metro quadrado ")
	}
	preco := 80.0
	litrosNaLata := 18.0
	// 1 l para 3 m }
	litros := metroFloat / 3.0
	latas := litros / litrosNaLata
	total := latas * preco

	fmt.Println("vai precisar de", latas, "latas")
	fmt.Println("vai sair um total de", total, "Reais")

}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	atualizaSalario := 0.0
	for {
		fmt.Println("Digite o seu salario:")
		pegaSalario, _ := reader.ReadString('\n')
		limpasalario := strings.TrimSpace(pegaSalario)
		salarioInicial, err := strconv.ParseFloat(limpasalario, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		atualizaSalario = salarioInicial
		break
	}
	salarioFinal := 0.0
	aumentaPorcentagem := 0.015
	for i := 1995; i < 2022; i++ {
		aumentaPorcentagem = aumentaPorcentagem * 2
		salarioFinal = atualizaSalario * (aumentaPorcentagem + 1)
		fmt.Printf("Ano: %v Porcentagem: %v Salario: %.2f \n ", i, aumentaPorcentagem, salarioFinal) // só funciona com o printF
		fmt.Println("ano:", i)
	}
	arredondaSalarioFinal := math.Floor(salarioFinal*100) / 100
	fmt.Printf("Salario final de: %.2f ", arredondaSalarioFinal)
}

package main

import (
	"fmt"
	"math"
)

func main() {

	//Esse funcionário foi contratado em 1995, com salário inicial de R$ 1.000,00;
	//Em 1996 recebeu aumento de 1,5% sobre seu salário inicial;
	//A partir de 1997 (inclusive), os aumentos salariais sempre correspondem ao dobro do
	// percentual do ano anterior
	// Faça um programa que determine o salário atual desse funcionário.
	salarioInicial := 1000.00
	salarioFinal := 0.0
	aumentaPorcentagem := 0.015
	for i := 1995; i < 2022; i++ {
		aumentaPorcentagem = aumentaPorcentagem * 2
		salarioFinal = salarioInicial * (aumentaPorcentagem + 1)
		fmt.Printf("Ano: %v Porcentagem: %v Salario: %.2f \n ", i, aumentaPorcentagem, salarioFinal)
	}
	arredondaSalarioFinal := math.Floor(salarioFinal*100) / 100
	fmt.Printf("Salario final de: %.2f ", arredondaSalarioFinal)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Altere o programa anterior permitindo ao usuário informar as populações e as taxas de crescimento iniciais.
//Valide a entrada e permita repetir a operação.
func main() {
	//Erros/Dicas:
	//Ter a atenção para calcular os anos de crescimento com base na cidade com maior taxa de crescimento
	// Se a taxa da cidade A for maior, comparar sempre com a cidade A
	// Se a taxa da cidade B for maior, comparar sempre com a cidade B
	// Ideal é comparar por taxa!
	// Ficar atento quando as taxas forem iguais.
	// No calculo de anos necessários, falta fazer o for
	reader := bufio.NewReader(os.Stdin)
	habitantesA1 := 0
	habitantesB2 := 0
	taxA1 := 0
	taxB2 := 0

	for {
		fmt.Println("Habitantes A:")
		habitantesA, _ := reader.ReadString('\n')
		limpaHabitantesA := strings.TrimSpace(habitantesA)
		habA, err := strconv.Atoi(limpaHabitantesA)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		} else {
			habitantesA1 = habA
			break
		}
	}
	for {
		fmt.Println("Habitantes B:")
		habitantesB, _ := reader.ReadString('\n')
		limpaHabitantesB := strings.TrimSpace(habitantesB)
		habB, err := strconv.Atoi(limpaHabitantesB)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		} else {
			habitantesB2 = habB
			break
		}
	}

	for {
		fmt.Println("Taxa A:")
		taxA, _ := reader.ReadString('\n')
		limpaTaxA := strings.TrimSpace(taxA)
		taxaA, err := strconv.Atoi(limpaTaxA)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		} else {
			taxA1 = taxaA
			break
		}

	}

	for {

		fmt.Println("Taxa B:")
		taxB, _ := reader.ReadString('\n')
		limpaTaxB := strings.TrimSpace(taxB)
		taxaB, err := strconv.Atoi(limpaTaxB)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		} else {
			taxB2 = taxaB
			break
		}
	}

	anos := 0
	if habitantesA1 > habitantesB2 {
		taxA1 = ((habitantesA1 * taxA1) / 100)
		taxB2 = ((habitantesB2 * taxB2) / 100)
		anos = anos + 1.0
		fmt.Println("Serão necessarios", anos, "anos para que a população do pais A ultrapasse ou iguale a população do Pais B")
	}
}

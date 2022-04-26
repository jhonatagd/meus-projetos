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
	fmt.Println("Primeiro numero:")
	um, _ := reader.ReadString('\n')
	limpaUm := strings.TrimSpace(um)
	numeroum, err := strconv.ParseFloat(limpaUm, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Segundo numero:")
	dois, _ := reader.ReadString('\n')
	limpaDois := strings.TrimSpace(dois)
	numeroDois, err := strconv.ParseFloat(limpaDois, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Digite a operação que deseja realizar: (+, -, /, *):")
	pegaOperacao, err := reader.ReadString('\n')
	limpaOperacao := strings.TrimSpace(pegaOperacao)
	limpaOperacao = strings.ToUpper(limpaOperacao)
	if err != nil {
		fmt.Println("O que você digitou não é uma operação")
	}
	mais := "+"
	menos := "-"
	divisao := "/"
	multiplicar := "*"

	if limpaOperacao == mais {
		negativoOuPositivo := numeroum + numeroDois
		verificaInteiroOuDecimal := (math.Floor(negativoOuPositivo + 0.0/0.5))
		somaFloat := numeroum + numeroDois
		converte := int(somaFloat)

		if converte%2 == 0 {
			fmt.Println(negativoOuPositivo, "Numero par")
		} else {
			fmt.Println(negativoOuPositivo, "Numero Impar")
		}
		if negativoOuPositivo < 0 {
			fmt.Println(negativoOuPositivo, "É negativo")
		} else {
			fmt.Println(negativoOuPositivo, "É Positivo")
		}
		if negativoOuPositivo == verificaInteiroOuDecimal {
			fmt.Println(negativoOuPositivo, "Inteiro")
		} else {
			fmt.Println(negativoOuPositivo, "Decimal")
		}
	}

	if limpaOperacao == menos {
		negativoOuPositivo := numeroum - numeroDois
		verificaInteiroOuDecimal := (math.Floor(negativoOuPositivo + 0.0/0.5))
		somaFloat := numeroum - numeroDois
		converte := int(somaFloat)

		if converte%2 == 0 {
			fmt.Println(negativoOuPositivo, "Numero par")
		} else {
			fmt.Println(negativoOuPositivo, "Numero Impar")
		}
		if negativoOuPositivo < 0 {
			fmt.Println(negativoOuPositivo, "É negativo")
		} else {
			fmt.Println(negativoOuPositivo, "É Positivo")
		}
		if negativoOuPositivo == verificaInteiroOuDecimal {
			fmt.Println(negativoOuPositivo, "Inteiro")
		} else {
			fmt.Println(negativoOuPositivo, "Decimal")
		}
	}
	if limpaOperacao == divisao {
		negativoOuPositivo := numeroum / numeroDois
		verificaInteiroOuDecimal := (math.Floor(negativoOuPositivo + 0.0/0.5))
		somaFloat := numeroum / numeroDois
		converte := int(somaFloat)

		if converte%2 == 0 {
			fmt.Println(negativoOuPositivo, "Numero par")
		} else {
			fmt.Println(negativoOuPositivo, "Numero Impar")
		}
		if negativoOuPositivo < 0 {
			fmt.Println(negativoOuPositivo, "É negativo")
		} else {
			fmt.Println(negativoOuPositivo, "É Positivo")
		}
		if negativoOuPositivo == verificaInteiroOuDecimal {
			fmt.Println(negativoOuPositivo, "Inteiro")
		} else {
			fmt.Println(negativoOuPositivo, "Decimal")
		}
	}
	if limpaOperacao == multiplicar {
		negativoOuPositivo := numeroum * numeroDois
		verificaInteiroOuDecimal := (math.Floor(negativoOuPositivo + 0.0/0.5))
		somaFloat := numeroum * numeroDois
		converte := int(somaFloat)

		if converte%2 == 0 {
			fmt.Println(negativoOuPositivo, "Numero par")
		} else {
			fmt.Println(negativoOuPositivo, "Numero Impar")
		}
		if negativoOuPositivo < 0 {
			fmt.Println(negativoOuPositivo, "É negativo")
		} else {
			fmt.Println(negativoOuPositivo, "É Positivo")
		}
		if negativoOuPositivo == verificaInteiroOuDecimal {
			fmt.Println(negativoOuPositivo, "Inteiro")
		} else {
			fmt.Println(negativoOuPositivo, "Decimal")
		}
	}
}

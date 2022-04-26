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
	fmt.Println("Litros colocado:")
	comb, _ := reader.ReadString('\n')
	limpaComb := strings.TrimSpace(comb)
	combustível, err := strconv.ParseFloat(limpaComb, 64)
	if err != nil {
		fmt.Println("O que voce digitou não é um numero")
	}
	fmt.Println("Tipo de combustível: (A-álcool - G-gasolina)")
	gaso, err := reader.ReadString('\n')
	tipoDaGasosa := strings.TrimSpace(gaso)
	tipoDaGasosa = strings.ToUpper(tipoDaGasosa)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}

	if combustível <= 20 && tipoDaGasosa == "A" {
		litros := combustível * 1.90
		descontaDesconto := ((litros * 3) / 100)
		teste := litros - descontaDesconto
		fmt.Println("Total de:", teste)
	} else if combustível > 20 && tipoDaGasosa == "A" {
		litros := combustível * 1.90
		descontaDesconto := ((litros * 5) / 100)
		teste := litros - descontaDesconto
		fmt.Println("Total de:", teste)
	} else if combustível <= 20 && tipoDaGasosa == "G" {
		litros := combustível * 2.50
		descontaDesconto := ((litros * 4) / 100)
		teste := litros - descontaDesconto
		fmt.Println("Total de:", teste)
	} else if combustível > 20 && tipoDaGasosa == "G" {
		litros := combustível * 2.50
		descontaDesconto := ((litros * 6) / 100)
		teste := litros - descontaDesconto
		fmt.Println("Total de:", teste)
	}

}

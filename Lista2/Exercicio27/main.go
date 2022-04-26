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
	fmt.Println("Quantidade em Kg de morangos:")
	morang, _ := reader.ReadString('\n')
	limpaMo := strings.TrimSpace(morang)
	kgMorango, err := strconv.ParseFloat(limpaMo, 64)
	if err != nil {
		fmt.Println("O que voce digitou não é um kg")
	}
	fmt.Println("Quantidade em Kg de maças:")
	maca, _ := reader.ReadString('\n')
	limpaMaca := strings.TrimSpace(maca)
	kgMaca, err := strconv.ParseFloat(limpaMaca, 64)
	if err != nil {
		fmt.Println("O que voce digitou não é um kg")
	}
	//                     Até 5 Kg           Acima de 5 Kg
	//Morango         R$ 2,50 por Kg          R$ 2,20 por Kg
	//Maçã            R$ 1,80 por Kg          R$ 1,50 por Kg
	testeMaca := kgMaca
	testeMorango := kgMorango
	if kgMaca > 5 {
		kgMaca = kgMaca * 1.50
	} else {
		kgMaca = kgMaca * 1.80
	}
	if kgMorango > 5 {
		kgMorango = kgMorango * 2.20
	} else {
		kgMorango = kgMorango * 2.50
	}

	totalKgFrutas := kgMaca + kgMorango
	teste := kgMaca + kgMorango
	if (totalKgFrutas > 8) || (teste > 25) {
		descobreDesconto := (5 * teste) / 100
		descontaDesconto := teste - descobreDesconto
		fmt.Println("Quantidade de:", testeMorango, "kg de morangos e", testeMaca, "kg de maça, total de:", descontaDesconto)
	} else if totalKgFrutas < 8 {
		fmt.Println("Quantidade de:", testeMorango, "kg de morangos e", testeMaca, "kg de maça, total de:", teste)
	}

}

//Se o cliente comprar mais de 8 Kg em frutas ou o valor total da compra ultrapassar R$ 25,00
// receberá ainda um desconto de 10% sobre este total.
//ler a quantidade (em Kg) de morangos e a quantidade (em Kg) de maças adquiridas e escreva o valor a ser pago pelo cliente.

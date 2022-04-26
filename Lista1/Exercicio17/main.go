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
	var metroint float64
	var err error

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Tamanho em metros quadrados da área a ser pintada: ")
		metros, _ := reader.ReadString('\n')
		limpaMetros := strings.TrimSpace(metros)
		metroint, err = strconv.ParseFloat(limpaMetros, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um metro quadrado ")
		} else {
			break
		}
	}

	preco := 80.0
	litrosNaLata := 18.0
	litros := metroint / 6.0
	latas := (math.Ceil(litros / litrosNaLata))
	total := latas * preco

	preco2 := 25.0
	litrosNoGalao2 := 3.6
	litros2 := metroint / 6.0
	galoes := (math.Ceil(litros2 / litrosNoGalao2))
	total2 := galoes * preco2

	fmt.Println("vai precisar de", latas, "lata de 18 litros.")
	fmt.Println("vai sair um total de", total, "Reais.")
	fmt.Println(" ")

	fmt.Println("vai precisar de", galoes, "galões de 3,6 litros.")
	fmt.Println("vai sair um total de", total2, "reais.")

}

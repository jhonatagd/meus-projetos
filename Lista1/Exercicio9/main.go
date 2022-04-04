package main

// peça a temperatura em graus Fahrenheit, transforme e mostre a temperatura em graus Celsius. C = 5 * ((F-32) / 9).

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("temperatura em graus Fahrenheit: ")
	pegaTemperatura, _ := reader.ReadString('\n')
	limpaTemperatura := strings.TrimSpace(pegaTemperatura)
	TemperaturaFloat, err := strconv.ParseFloat(limpaTemperatura, 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma temperatura em graus Fahrenheit ")
	}
	var Formula = 5 * ((TemperaturaFloat - 32) / 9)
	fmt.Println("em graus Celsius: ")
	fmt.Println(Formula)

}

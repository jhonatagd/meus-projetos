package main

// peça a temperatura em graus Celsius, transforme e mostre em graus Fahrenheit.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("temperatura em graus Celsius: ")
	pegaTemperatura, _ := reader.ReadString('\n')
	limpaTemperatura := strings.TrimSpace(pegaTemperatura)
	TemperaturaFloat, err := strconv.ParseFloat(limpaTemperatura, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um graus Celsius ")
	}
	var descobreGrau = (TemperaturaFloat * 1.8) + 32
	fmt.Println("em graus Fahrenheit: ")
	fmt.Println(descobreGrau)

}

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
	fmt.Println("Valor da primeira nota: ")
	primeiraNota, _ := reader.ReadString('\n')
	limpaPrimeiraNota := strings.TrimSpace(primeiraNota)
	primeiraNotaFloat, err := strconv.ParseFloat(limpaPrimeiraNota, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	fmt.Println("Valor da segunda nota: ")
	segundaNota, _ := reader.ReadString('\n')
	LimpasegundaNota := strings.TrimSpace(segundaNota)
	segundaNotaFloat, err := strconv.ParseFloat(LimpasegundaNota, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	media := ((primeiraNotaFloat + segundaNotaFloat) / 2.0)

	if media >= 9.0 && media <= 10.0 {
		fmt.Println("Media de:", media)
		fmt.Println("Conceito correspondente: A")
		fmt.Print("APROVADO")
	} else if media >= 7.5 && media < 9.0 {
		fmt.Println("Media de:", media)
		fmt.Println("Conceito correspondente: B")
		fmt.Print("APROVADO")
	} else if media >= 6.0 && media < 7.5 {
		fmt.Println("Media de:", media)
		fmt.Println("Conceito correspondente: C")
		fmt.Print("APROVADO")
	} else if media >= 4.0 && media < 6.0 {
		fmt.Println("Media de:", media)
		fmt.Println("Conceito correspondente: D")
		fmt.Print("REPROVADO")
	} else if media >= 0.0 && media < 4.0 {
		fmt.Println("Media de:", media)
		fmt.Println("Conceito correspondente: E")
		fmt.Print("REPROVADO")
	}

}

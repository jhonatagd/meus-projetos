package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um Programa que leia 4 notas, mostre as notas e a média na tela.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var armazenaNotas [4]float64
	somaMedia := 0.0
	for i := 0; i < len(armazenaNotas); i++ {
		fmt.Println(i+1, "Nota:")
		pegaNota, _ := reader.ReadString('\n')
		limpaNota := strings.TrimSpace(pegaNota)
		nota, err := strconv.ParseFloat(limpaNota, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		} else if nota < 0 {
			fmt.Println("Valor invalido, digite novamente")
			i--
			continue
		}
		armazenaNotas[i] = nota
		somaMedia += nota
	}
	fmt.Println("Media das notas:", somaMedia/4)
	fmt.Println("Notas:", armazenaNotas)

}

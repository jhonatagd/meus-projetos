package main

//Para homens: (72.7*h) - 58 // Para mulheres: (62.1*h) - 44.7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" sua altura: ")
	altura, _ := reader.ReadString('\n')
	LimpaAltura := strings.TrimSpace(altura)
	alturaFloat, err := strconv.ParseFloat(LimpaAltura, 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma altura ")
	}

	medePesoHomem := (72.7 * alturaFloat) - 58
	medepesoMulher := (62.1 * alturaFloat) - 44.7

	fmt.Println("Seu peso ideal se voce for homem é: ")
	fmt.Println(medePesoHomem)
	fmt.Println("Seu peso ideal se voce for mulher é: ")
	fmt.Println(medepesoMulher)
}

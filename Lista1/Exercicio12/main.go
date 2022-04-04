package main

// dados de entrada a altura de uma pessoa, construa um algoritmo que calcule seu peso ideal ==> (72.7*altura) - 58

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
	medePeso := (72.7 * alturaFloat) - 58

	fmt.Println("Seu peso ideal é: ")
	fmt.Println(medePeso)
}

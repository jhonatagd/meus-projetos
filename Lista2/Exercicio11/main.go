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
	fmt.Println("Salário atual: ")
	salario, _ := reader.ReadString('\n')
	limpaSalario := strings.TrimSpace(salario)
	salarioInt, err := strconv.Atoi(limpaSalario)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}

	s20 := "20%"
	s15 := "15%"
	s10 := "10%"
	s5 := "5%"
	aumento20 := (salarioInt * 20) / 100
	aumento15 := (salarioInt * 15) / 100
	aumento10 := (salarioInt * 10) / 100
	aumento5 := (salarioInt * 5) / 100

	novoSalario20 := salarioInt + aumento20
	novoSalario15 := salarioInt + aumento15
	novoSalario10 := salarioInt + aumento10
	novoSalario5 := salarioInt + aumento5

	if salarioInt <= 280 {
		fmt.Println("O salário antes do reajuste:", salarioInt)
		fmt.Println("O percentual de aumento aplicado:", s20)
		fmt.Println("O valor do aumento:", aumento20)
		fmt.Println("O novo salário, após o aumento:", novoSalario20)
	} else if salarioInt > 280 && salarioInt <= 700 {
		fmt.Println("O salário antes do reajuste:", salarioInt)
		fmt.Println("O percentual de aumento aplicado:", s15)
		fmt.Println("O valor do aumento: ", aumento15)
		fmt.Println("O novo salário, após o aumento:", novoSalario15)
	} else if salarioInt > 700 && salarioInt <= 1500 {
		fmt.Println("O salário antes do reajuste:", salarioInt)
		fmt.Println("O percentual de aumento aplicado:", s10)
		fmt.Println("O valor do aumento:", aumento10)
		fmt.Println("O novo salário, após o aumento:", novoSalario10)
	} else if salarioInt > 1500 {
		fmt.Println("o salário antes do reajuste:", salarioInt)
		fmt.Println("o percentual de aumento aplicado:", s5)
		fmt.Println("o valor do aumento:", aumento5)
		fmt.Println("o novo salário, após o aumento:", novoSalario5)
	}

}

package main

// deverá pedir ao usuário o valor da sua hora e a quantidade de horas trabalhadas no mês.
//Salário Bruto: (5 * 220)        : R$ 1100,00
//       (-) IR (5%)                     : R$   55,00
//       (-) INSS ( 10%)                 : R$  110,00
//       FGTS (11%)                      : R$  121,00
//      Total de descontos              : R$  165,00
//      Salário Liquido                 : R$  935,
//Salário Bruto até 900 (inclusive) - isento
//Salário Bruto até 1500 (inclusive) - desconto de 5%
//Salário Bruto até 2500 (inclusive) - desconto de 10%
//salario Bruto acima de 2500 - desconto de 20%

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Valor da sua hora trabalhada: ")
	valorHora, _ := reader.ReadString('\n')
	limpaHora := strings.TrimSpace(valorHora)
	valorHoraInt, err := strconv.Atoi(limpaHora)
	if err != nil {
		fmt.Println("O que você digitou não é uma hora")
	}
	fmt.Println("Quantidade de horas trabalhadas no mês: ")
	horasTrabalhadas, _ := reader.ReadString('\n')
	limpaHoraTrabalhada := strings.TrimSpace(horasTrabalhadas)
	horaTrabalhadaInt, err := strconv.Atoi(limpaHoraTrabalhada)
	if err != nil {
		fmt.Println("O que você digitou não é uma hora")
	}
	salarioBruto := valorHoraInt * horaTrabalhadaInt
	iNSS := (10 * salarioBruto) / 100 // 10%
	fGTS := (11 * salarioBruto) / 100 // 11%
	if salarioBruto <= 900 {
		fmt.Println("salario bruto de:", salarioBruto)
		fmt.Println("Imposto de renda: 0")
		fmt.Println("INSS: 0")
		fmt.Println("FGTS: 0")
		fmt.Println("Total de descontos: 0")
		fmt.Println("salario liquido:", salarioBruto)
	} else if salarioBruto > 900 && salarioBruto <= 1500 {
		fmt.Println("salario bruto de:", salarioBruto)
		fmt.Println("Imposto de renda:", ((5 * salarioBruto) / 100))
		fmt.Println("INSS:", (10*salarioBruto)/100)
		fmt.Println("FGTS:", fGTS)
		fmt.Println("Total de descontos", ((5*salarioBruto)/100)+iNSS+fGTS)
		fmt.Println("salario liquido:", ((5 * salarioBruto) / 100), iNSS+fGTS, -salarioBruto)
	} else if salarioBruto > 1500 && salarioBruto <= 2500 {
		fmt.Println("salario bruto de:", salarioBruto)
		fmt.Println("Imposto de renda:", ((10 * salarioBruto) / 100))
		fmt.Println("INSS:", (10*salarioBruto)/100)
		fmt.Println("FGTS:", fGTS)
		fmt.Println("Total de descontos", ((10 * salarioBruto) / 100), iNSS+fGTS)
		fmt.Println("salario liquido:", ((10 * salarioBruto) / 100), iNSS+fGTS, -salarioBruto)
	} else if salarioBruto > 2500 {
		fmt.Println("salario bruto de:", salarioBruto)
		fmt.Println("Imposto de renda:", ((20 * salarioBruto) / 100))
		fmt.Println("INSS:", (10*salarioBruto)/100)
		fmt.Println("FGTS:", fGTS)
		fmt.Println("Total de descontos", ((20 * salarioBruto) / 100), iNSS+fGTS)
		fmt.Println("salario liquido:", ((20 * salarioBruto) / 100), iNSS+fGTS, -salarioBruto)
	}
}

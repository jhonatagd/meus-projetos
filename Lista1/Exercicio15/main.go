package main

// ganha por hora e o número de horas trabalhadas no mês.
//  total do seu salário no referido mês = gph * nh
// escontados 11% para o Imposto de Renda, 8% para o INSS e 5% para o sindicato
// a  salário bruto.
// b quanto pagou ao INSS.
// c quanto pagou ao sindicato.
// d o salário líquido.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Quanto ganha por hora: ")
	ganhaPorHora, _ := reader.ReadString('\n')
	LimpaHora := strings.TrimSpace(ganhaPorHora)
	GanhaPorHoraFloat, err := strconv.ParseFloat(LimpaHora, 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma hora")
	}

	fmt.Println("Horas trabalhadas no mes: ")
	horasNoMes, _ := reader.ReadString('\n')
	limpaHorasNoMes := strings.TrimSpace(horasNoMes)
	horasNoMesFloat, err := strconv.ParseFloat(limpaHorasNoMes, 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma hora")
	}
	ir := 11.0 / 100.0
	inss := 8.0 / 100.0
	sindicat0 := 5.0 / 100.0
	descobreQuantoGanhouNoMes := GanhaPorHoraFloat * horasNoMesFloat // descontar 11%
	// %/100
	salariodescontoImpostoRendaAtualizado := ir * descobreQuantoGanhouNoMes
	//  8% para o INSS
	salarioAtualizadoComDescontoInss := inss * descobreQuantoGanhouNoMes
	// 5% para o sindicato
	salarioAtualizadoComDescontoSindicato := sindicat0 * descobreQuantoGanhouNoMes
	ValorDescontado := salariodescontoImpostoRendaAtualizado + salarioAtualizadoComDescontoInss + salarioAtualizadoComDescontoSindicato
	salarioLiquido := descobreQuantoGanhouNoMes - ValorDescontado

	fmt.Println("voce ganhou por mes: ")
	fmt.Println(descobreQuantoGanhouNoMes)
	fmt.Println("Desconto do Imposto de renda: ")
	fmt.Println(salariodescontoImpostoRendaAtualizado)
	fmt.Println("Desconto do Inss: ")
	fmt.Println(salarioAtualizadoComDescontoInss)
	fmt.Println("Desconto do Sindicato: ")
	fmt.Println(salarioAtualizadoComDescontoSindicato)
	fmt.Println("total de descontos:")
	fmt.Println(ValorDescontado)
	fmt.Println(" salario liquido:")
	fmt.Println(salarioLiquido)

}

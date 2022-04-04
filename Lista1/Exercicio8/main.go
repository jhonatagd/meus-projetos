package main

//quanto você ganha por hora e o número de horas trabalhadas no mês. Calcule e mostre o total do seu salário no referido mês.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite quantas horas voce trabalhou no mes: ")
	horasPorDia, _ := reader.ReadString('\n')
	horasPorDiaLimpa := strings.TrimSpace(horasPorDia)
	horasPorDiaFloat, err := strconv.ParseFloat(horasPorDiaLimpa, 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma hora")
	}

	fmt.Print("Digite quanto voce ganha por hora: ")
	ganhaPorHora, _ := reader.ReadString('\n')
	ganhaPorHoraLimpa, err := strconv.ParseFloat(strings.TrimSpace(ganhaPorHora), 64)
	if err != nil {
		fmt.Println("O que você digitou não é uma hora")
	}
	DescobreHorasPorGanho := horasPorDiaFloat * ganhaPorHoraLimpa
	fmt.Println("ganho por horas trabalhadas: ")
	fmt.Println(DescobreHorasPorGanho)
}

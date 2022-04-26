package main

//leia um número e exiba o dia correspondente da semana.
//(1-Domingo, 2- Segunda, etc.), se digitar outro valor deve aparecer valor inválido.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Numero correspondente da semana: ")
	semana, _ := reader.ReadString('\n')
	limpaSemana := strings.TrimSpace(semana)
	numeroSemanaInt, err := strconv.Atoi(limpaSemana)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	diaInvalido := "Valor invalido"

	if numeroSemanaInt == 1 {
		fmt.Println("domingo")
	} else if numeroSemanaInt == 2 {
		fmt.Println("segunda")
	} else if numeroSemanaInt == 3 {
		fmt.Println("terça")
	} else if numeroSemanaInt == 4 {
		fmt.Println("quarta")
	} else if numeroSemanaInt == 5 {
		fmt.Println("quinta")
	} else if numeroSemanaInt == 6 {
		fmt.Println("sexta")
	} else if numeroSemanaInt == 7 {
		fmt.Println("sabado")
	} else {
		fmt.Println(diaInvalido)
	}
}

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

	fmt.Println("Numero de pessoas:")
	pegaIdade, _ := reader.ReadString('\n')
	limpaIdade := strings.TrimSpace(pegaIdade)
	nPessoas, err := strconv.Atoi(limpaIdade)

	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	valorNota := 0
	for i := 0; i < nPessoas; i++ {
		fmt.Println("Digitar o ", i+1, "° numero:")
		nUm, _ := reader.ReadString('\n')
		limpaNUm := strings.TrimSpace(nUm)
		nota, err := strconv.Atoi(limpaNUm)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			i = i - 1
		}
		if nota > 0 {

			valorNota = valorNota + nota
		} else {
			fmt.Println("Nota invalida")
			i = i - 1
		}
	}

	calculaMedia := valorNota / nPessoas
	if calculaMedia >= 0 && calculaMedia <= 25 {
		fmt.Println("Turma jovem")
	} else if calculaMedia >= 26 && calculaMedia <= 60 {
		fmt.Println("Turma adulta")
	} else {
		fmt.Println("Turma idosa")
	}
	fmt.Println("Soma:", valorNota)

} //Faça um programa que peça para n pessoas a sua idade, ao final o programa devera verificar
//se a média de idade da turma varia entre 0 e 25,26 e 60 e maior que 60; e então, dizer se
// a turma é jovem, adulta ou idosa, conforme a média calculada.

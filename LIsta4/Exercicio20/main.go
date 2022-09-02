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
	var listaSalario []int
	var listaAbono []int
	quantidadeDeSalario := 0
	for {
		fmt.Println("Digite seu salario de Dezembro:")
		pegaSalario, _ := reader.ReadString('\n')
		limpaASalarioInt := strings.TrimSpace(pegaSalario)
		salario, err := strconv.Atoi(limpaASalarioInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else if salario == 0 {
			break
		}
		listaSalario = append(listaSalario, salario)
		quantidadeDeSalario += 1
	}
	valorMinimoAbono := 0
	for i := 0; i < quantidadeDeSalario; i++ {
		if listaSalario[i] < 200 {
			listaAbono = append(listaAbono, 100)
			valorMinimoAbono += 1
		} else {
			descobrePorcentagem := ((20 / 100) * listaSalario[i])
			listaAbono = append(listaAbono, listaSalario[i]+descobrePorcentagem)
		}
	}

	//gastosabono
	totalAbono := 0
	for i := 0; i < quantidadeDeSalario; i++ {
		totalAbono = totalAbono + listaAbono[i]
	}

	fmt.Println("Salarios:", listaSalario)
	fmt.Print("Abono:", listaAbono)
	fmt.Println("Foram processados", quantidadeDeSalario, "colaboradores")
	fmt.Println("Total gasto com abonos:", totalAbono)
	fmt.Println("Total gasto com abonos:", totalAbono)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// regulamento de pesca do estado de São Paulo (50 quilos)
// deve pagar uma multa de R$ 4,00 por quilo excedente.
// leia a variável peso (peso de peixes) e calcule o excesso.
// Gravar na variável excesso a quantidade de quilos além do limite
// na variável multa o valor da multa que João deverá pagar.
// Imprima os dados do programa com as mensagens adequadas.

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("kilo do peixe: ")
	kgDoPeixe, _ := reader.ReadString('\n')
	limpaKgDoPeixe := strings.TrimSpace(kgDoPeixe)
	kgDoPeixeFloat, err := strconv.ParseFloat(limpaKgDoPeixe, 64) // variavel onde esta o kg do peixe

	if err != nil {
		fmt.Println("Digite o kilo novamente")
	}
	multa := 4.0
	kgPx := 50.0
	KilosExcedente := 0.0

	if kgDoPeixeFloat >= kgPx {
		KilosExcedente = kgDoPeixeFloat - kgPx
	}

	fmt.Println("Sua multa é de: ")
	fmt.Println(KilosExcedente * multa)

}

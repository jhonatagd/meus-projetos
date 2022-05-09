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

	fmt.Println("Habitantes A:")
	habitantesA, _ := reader.ReadString('\n')
	limpaHabitantesA := strings.TrimSpace(habitantesA)
	habA, err := strconv.ParseFloat(limpaHabitantesA, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Habitantes B:")
	habitantesB, _ := reader.ReadString('\n')
	limpaHabitantesB := strings.TrimSpace(habitantesB)
	habB, err := strconv.ParseFloat(limpaHabitantesB, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Taxa A:")
	taxA, _ := reader.ReadString('\n')
	limpaTaxA := strings.TrimSpace(taxA)
	taxaA, err := strconv.ParseFloat(limpaTaxA, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	fmt.Println("Taxa B:")
	taxB, _ := reader.ReadString('\n')
	limpaTaxB := strings.TrimSpace(taxB)
	taxaB, err := strconv.ParseFloat(limpaTaxB, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	anos := 0
	if habA > habB {
		taxaA = ((habA * taxaA) / 100)
		taxaB = ((habB * taxaB) / 100)
		anos = anos + 1.0
		fmt.Println("Serão necessarios", anos, "anos para que a população do pais A ultrapasse ou iguale a população do Pais B")
	}
}

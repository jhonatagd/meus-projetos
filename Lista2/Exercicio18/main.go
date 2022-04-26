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
	fmt.Println("Dia: ")
	dia, _ := reader.ReadString('\n')
	limpaDia := strings.TrimSpace(dia)
	diaInt, err := strconv.Atoi(limpaDia)
	if err != nil {
		fmt.Println("O que você digitou não é um dia")
	}
	fmt.Println("mês: ")
	mes, _ := reader.ReadString('\n')
	limpaMes := strings.TrimSpace(mes)
	mesInt, err := strconv.Atoi(limpaMes)
	if err != nil {
		fmt.Println("O que você digitou não é um dia")
	}
	fmt.Println("Ano: ")
	ano, _ := reader.ReadString('\n')
	limpaAno := strings.TrimSpace(ano)
	anoInt, err := strconv.Atoi(limpaAno)
	if err != nil {
		fmt.Println("O que você digitou não é um dia")
	}
	restoDivisão4 := anoInt % 4
	restoDivisão100 := anoInt % 100
	restoDivisão400 := anoInt % 400

	if diaInt <= 31 && mesInt == 1 {
		fmt.Println("Data valida")
	} else if diaInt <= 28 && mesInt == 2 {
		fmt.Println("Data valida")
	} else if diaInt <= 31 && mesInt == 3 {
		fmt.Println("Data valida")
	} else if diaInt <= 31 && mesInt == 5 {
		fmt.Println("Data valida")
	} else if diaInt <= 31 && mesInt == 7 {
		fmt.Println("Data valida")
	} else if diaInt <= 31 && mesInt == 8 {
		fmt.Println("Data valida")
	} else if diaInt <= 31 && mesInt == 10 {
		fmt.Println("Data valida")
	} else if diaInt <= 31 && mesInt == 12 {
		fmt.Println("Data valida")
	} else if diaInt <= 30 && mesInt == 4 {
		fmt.Println("Data valida")
	} else if diaInt <= 30 && mesInt == 6 {
		fmt.Println("Data valida")
	} else if diaInt <= 30 && mesInt == 9 {
		fmt.Println("Data valida")
	} else if diaInt <= 30 && mesInt == 11 {
		fmt.Println("Data valida")
	} else if diaInt <= 29 && (((restoDivisão4 == 0 && restoDivisão100 != 0) || (restoDivisão100 == 0 && restoDivisão400 == 0)) && (mesInt == 2)) {
		fmt.Println("Data valida")
	} else {
		fmt.Println("Data invalida")
	}
}

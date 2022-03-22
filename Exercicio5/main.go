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
	fmt.Println("Converta metros para centímetros: ")
	pegaMetro, _ := reader.ReadString('\n')
	limpaMetro := strings.TrimSpace(pegaMetro)
	metroFloat, err := strconv.ParseFloat(limpaMetro, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um metro ")
	}
	var metroParaCentimetros = metroFloat * 100
	fmt.Println("em centímetros: ")
	fmt.Println(metroParaCentimetros)

}

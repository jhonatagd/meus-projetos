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
	fmt.Println("Primeiro numero: ")
	pegaPrimeiroNumero, _ := reader.ReadString('\n')
	limpaPrimeiroNumero := strings.TrimSpace(pegaPrimeiroNumero)
	limpaPrimeiroNumeroInt, err := strconv.Atoi(limpaPrimeiroNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um numero ")
	}
	fmt.Println("Segundo numero: ")
	pegaSegundoNumero, _ := reader.ReadString('\n')
	limpaSegundoNumero := strings.TrimSpace(pegaSegundoNumero)
	limpaSegundoNumeroInt, err := strconv.Atoi(limpaSegundoNumero)
	if err != nil {
		fmt.Println("O que você digitou não é um numero ")
	}
	numerosIguais := limpaPrimeiroNumeroInt == limpaSegundoNumeroInt
	if numerosIguais {
		fmt.Println("Os numeros são iguais.")
	} else if limpaPrimeiroNumeroInt > limpaSegundoNumeroInt {
		fmt.Println(limpaPrimeiroNumeroInt, "maior que", limpaSegundoNumeroInt)
	} else {
		fmt.Println(limpaSegundoNumeroInt, "maior que", limpaPrimeiroNumeroInt)
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite a letra:")
	pegaLetra, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}
	limpaLetra := strings.TrimSpace(pegaLetra)

	vogalA := "A"
	vogalE := "E"
	vogalI := "I"
	vogalO := "O"
	vogalU := "U"

	if limpaLetra == vogalA {
		fmt.Println("é vogal.")
	} else if limpaLetra == vogalE {
		fmt.Println("é vogal.")
	} else if limpaLetra == vogalI {
		fmt.Println("é vogal.")
	} else if limpaLetra == vogalO {
		fmt.Println("é vogal.")
	} else if limpaLetra == vogalU {
		fmt.Println("é vogal.")
	} else {
		fmt.Println("é consoante.")
	}
}

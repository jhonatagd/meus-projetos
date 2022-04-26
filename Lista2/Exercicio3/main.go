package main

// Conforme a letra escrever: F - Feminino, M - Masculino, Sexo Inválido.
// verifique se uma letra digitada é "F" ou "M".

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite a letra:")
	pegaSexo, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}
	limpaSexo := strings.TrimSpace(pegaSexo)
	mDeMacho := "M"
	fDeFelinas := "F"

	if limpaSexo == mDeMacho {
		fmt.Println("Masculino")
	} else if limpaSexo == fDeFelinas {
		fmt.Println("Feminino")
	} else {
		fmt.Println("Sexo Inválido")
	}
}

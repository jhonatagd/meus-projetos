package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Faça um Programa que leia um vetor de 10 caracteres, e diga quantas consoantes foram lidas. Imprima as consoantes.
func main() {
	reader := bufio.NewReader(os.Stdin)
	var conjuntoCaracteres []string
	consoantes := 0
	for i := 0; i < 10; i++ {
		fmt.Println("Digite um caracter:")
		pegaCaractere, err := reader.ReadString('\n')
		pegaCaractereStrings := strings.TrimSpace(pegaCaractere)
		pegaCaractereStrings = strings.ToUpper(pegaCaractereStrings)
		if err != nil {
			fmt.Println("O que você digitou não é um caracter")
			i--
			continue
		}
		pegaCaracterLen := len(pegaCaractereStrings)
		if pegaCaracterLen > 1 {
			fmt.Println("O que você digitou não é um caracter")
			i--
			continue
		}

		if pegaCaractereStrings == "A" || pegaCaractereStrings == "E" || pegaCaractereStrings == "I" || pegaCaractereStrings == "O" || pegaCaractereStrings == "U" {
			continue
		} else {

			consoantes += +1
			conjuntoCaracteres = append(conjuntoCaracteres, pegaCaractereStrings)
		}

	}

	fmt.Println("quantas consoantes foram lidas:", consoantes)
	fmt.Println("Consoantes digitadas:", conjuntoCaracteres)
}

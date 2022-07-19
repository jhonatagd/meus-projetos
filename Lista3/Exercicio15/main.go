package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//A série de Fibonacci é formada pela seqüência 1,1,2,3,5,8,13,21,34,55,...
//Faça um programa capaz de gerar a série até o n−ésimo termo.
func main() {
	reader := bufio.NewReader(os.Stdin)
	//Errinho: Validação da entrada de dados. O código segue com o erro do usuário.
	nesimoNumero := 0
	for {
		fmt.Println("N:")
		termo, _ := reader.ReadString('\n')
		termo = strings.TrimSpace(termo)
		nesimo, err := strconv.Atoi(termo)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		nesimoNumero = nesimo
		break
	}
	primeiro := 1
	segundo := 1
	for i := 0; i < nesimoNumero; i++ {
		if i == 0 {
			fmt.Println(primeiro)
		} else if i == 1 {
			fmt.Println(segundo)
		} else {
			novoNumero := primeiro + segundo
			fmt.Println(novoNumero)
			primeiro = segundo
			segundo = novoNumero
		}
	}

}

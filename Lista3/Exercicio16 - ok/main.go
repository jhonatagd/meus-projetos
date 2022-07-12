package main

import (
	"fmt"
)

//A série de Fibonacci é formada pela seqüência 0,1,1,2,3,5,8,13,21,34,55,...
//Faça um programa que gere a série até que o valor seja maior que 500.
func main() {

	primeiro := 0
	segundo := 1
	i := 0
	for {
		if i == 0 {
			fmt.Println(primeiro)
		} else if i == 1 {
			fmt.Println(segundo)
		} else {
			novoNumero := primeiro + segundo
			fmt.Println(novoNumero)
			primeiro = segundo
			segundo = novoNumero

			if novoNumero > 500 {
				break
			}
		}

		i++
	}

}

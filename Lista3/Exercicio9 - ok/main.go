package main

import "fmt"

//Faça um programa que imprima na tela apenas os números ímpares entre 1 e 50.
func main() {
	for i := 1; i <= 50; i++ {
		verifica := i % 2
		if verifica != 0 {
			fmt.Println(i)
		}
	}
}

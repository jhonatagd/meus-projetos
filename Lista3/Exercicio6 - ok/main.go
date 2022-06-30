package main

import "fmt"

//Faça um programa que imprima na tela os números de 1 a 20, um abaixo do outro.
//Depois modifique o programa para que ele mostre os números um ao lado do outro.
func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 20; i++ {
		fmt.Print(i)
	}
}

package main

import (
	"fmt"
)

func main() {

	primeiro := 0
	segundo := 1
	for i := 0; i < 16; i++ {

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

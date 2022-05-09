package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		verifica := i % 2
		if verifica != 0 {
			fmt.Println(i)
		}

	}
}

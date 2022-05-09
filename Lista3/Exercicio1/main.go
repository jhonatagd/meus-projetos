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

	for {
		fmt.Println("Informe uma nota de 0 a 10: ")
		nota, _ := reader.ReadString('\n')
		nota = strings.TrimSpace(nota)
		notaInt, err := strconv.Atoi(nota)
		if err != nil {
			fmt.Println("O que você digitou não é um valor válido")
		}

		if notaInt < 0 || notaInt > 10 {
			fmt.Println("O que você digitou não é um valor válido")
		} else {
			fmt.Println("O que você digitou é válido")

			for i := notaInt; i <= notaInt+10; i++ {
				fmt.Println(i)
			}

			return
		}
	}
}

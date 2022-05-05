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

	fmt.Println("Que termo deseja encontrar:")
	termo, err := reader.ReadString('\n')
	pegaTermo := strings.TrimSpace(termo)
	pegaTermo = strings.ToUpper(pegaTermo)
	if err != nil {
		fmt.Println("O que você digitou não é um número")
	}
	ultimo := 1
	penultimo := 1
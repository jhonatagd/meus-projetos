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
	fmt.Println(" raio de um círculo: ")
	raio, _ := reader.ReadString('\n')
	limpaRaio := strings.TrimSpace(raio)
	raioFloat, err := strconv.ParseFloat(limpaRaio, 64)
	Pi := 3.14
	area := Pi * (raioFloat * raioFloat)
	if err != nil {
		fmt.Println("O que você digitou não é um raio ")
	}
	fmt.Println("Sua área: ")
	fmt.Println(area)
}

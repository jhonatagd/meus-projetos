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
	fmt.Println("Verifica se o ano é bissexto: ")
	pegaAno, _ := reader.ReadString('\n')
	limpaAno := strings.TrimSpace(pegaAno)
	anoInt, err := strconv.Atoi(limpaAno)
	if err != nil {
		fmt.Println("O que você digitou não é um ano")
	}
	restoDivisão4 := (anoInt % 4)
	restoDivisão100 := (anoInt % 100)
	restoDivisão400 := (anoInt % 400)

	if (restoDivisão4 == 0 && restoDivisão100 != 0) || (restoDivisão100 == 0 && restoDivisão400 == 0) {
		fmt.Println("É um ano bissexto")
	} else {
		fmt.Println("Não é bissexto")
	}
}

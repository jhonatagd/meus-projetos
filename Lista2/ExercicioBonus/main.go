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
	fmt.Println("Dinheiro no cartão:")
	colocaSaldo, _ := reader.ReadString('\n')
	limpaCartao := strings.TrimSpace(colocaSaldo)
	cartao, err := strconv.ParseFloat(limpaCartao, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	fmt.Println("Valor do brinquedo:")
	valor, _ := reader.ReadString('\n')
	limpavalor := strings.TrimSpace(valor)
	precoBrinquedo, err := strconv.ParseFloat(limpavalor, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}

	if cartao > precoBrinquedo {
		fmt.Println("Pagamento concluido, saldo de:", (cartao - precoBrinquedo))
	}
}

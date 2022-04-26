package main

// preço de três produtos
// informe qual produto você deve comprar, sabendo que a decisão é sempre pelo mais barato.
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Valor do Primeiro produto: ")
	produtoUm, _ := reader.ReadString('\n')
	limpaProdutoUm := strings.TrimSpace(produtoUm)
	produtoUmFloat, err := strconv.ParseFloat(limpaProdutoUm, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	fmt.Println("Valor do segundo produto: ")
	produtoDois, _ := reader.ReadString('\n')
	limpaProdutoDois := strings.TrimSpace(produtoDois)
	produtoDoisFloat, err := strconv.ParseFloat(limpaProdutoDois, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	fmt.Println("Valor do terceiro produto: ")
	produtoTres, _ := reader.ReadString('\n')
	limpaProdutoTres := strings.TrimSpace(produtoTres)
	produtoTresFloat, err := strconv.ParseFloat(limpaProdutoTres, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}

	numero1 := produtoUmFloat
	numero2 := produtoDoisFloat
	numero3 := produtoTresFloat

	if numero1 < numero2 && numero1 < numero3 {
		fmt.Println("O produto mais barato é ", numero1)
	} else if numero2 < numero3 && numero2 < numero1 {
		fmt.Println("O produto mais barato é ", numero2)
	} else if numero3 < numero2 && numero3 < numero1 {
		fmt.Println("O produto mais barato é ", numero3)
	}
}

// perguntar preço de produto, e dizer qual produto é mais barato

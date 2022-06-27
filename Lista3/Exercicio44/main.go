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
	somaCompra := 0.0
	valorPorItemCachorroQuente := 0.0
	valorPorItemBauruSimples := 0.0
	valorPorItemBauruComOvo := 0.0
	valorPorItemHambúrguer := 0.0
	valorPorItemCheeseburguer := 0.0
	valorPorItemRefrigerante := 0.0

	for {

		fmt.Println("Codigo do produto:  - Digite 0 para finalizar pedido")
		pegaCodigo, _ := reader.ReadString('\n')
		limpaCodigo := strings.TrimSpace(pegaCodigo)
		codigo, err := strconv.ParseFloat(limpaCodigo, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
	}
}

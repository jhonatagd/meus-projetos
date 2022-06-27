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
		fmt.Println("Quantidade de itens comprado:")
		pegaItens, _ := reader.ReadString('\n')
		limpaItem := strings.TrimSpace(pegaItens)
		quantidadeDeItens, err := strconv.ParseFloat(limpaItem, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
		}
		cachorroQuente := 100.00
		bauruSimples := 101.00
		bauruComOvo := 102.00
		hambúrguer := 103.00
		cheeseburguer := 104.00
		refrigerante := 105.00

		if codigo == cachorroQuente {
			somaCompra = somaCompra + (quantidadeDeItens * 1.20)
			valorPorItemCachorroQuente = valorPorItemCachorroQuente + quantidadeDeItens
		}
		if codigo == bauruSimples {
			somaCompra = somaCompra + (quantidadeDeItens * 1.30)
			valorPorItemBauruSimples = valorPorItemBauruSimples + quantidadeDeItens
		}
		if codigo == bauruComOvo {
			somaCompra = somaCompra + (quantidadeDeItens * 1.50)
			valorPorItemBauruComOvo = valorPorItemBauruComOvo + quantidadeDeItens
		}
		if codigo == hambúrguer {
			somaCompra = somaCompra + (quantidadeDeItens * 1.20)
			valorPorItemHambúrguer = valorPorItemHambúrguer + quantidadeDeItens
		}
		if codigo == cheeseburguer {
			somaCompra = somaCompra + (quantidadeDeItens * 1.30)
			valorPorItemCheeseburguer = valorPorItemCheeseburguer + quantidadeDeItens
		}
		if codigo == refrigerante {
			somaCompra = somaCompra + (quantidadeDeItens * 1.00)
			valorPorItemRefrigerante = valorPorItemRefrigerante + quantidadeDeItens
		}

		if codigo == 0 {
			fmt.Println("Valor a ser pago pelo item:")
			fmt.Println("Cachorro Quente:", valorPorItemCachorroQuente)
			fmt.Println("Bauru Simples:", valorPorItemBauruSimples)
			fmt.Println("Bauru com ovo:", valorPorItemBauruComOvo)
			fmt.Println("Hambúrguer:", valorPorItemHambúrguer)
			fmt.Println("Cheeseburguer:", valorPorItemCheeseburguer)
			fmt.Println("Refrigerante:", valorPorItemRefrigerante)
			fmt.Println("Total do seu pedido:", somaCompra)

			break

		}

	}
}

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
	atualizaCodigo := 0.0
	atualizaQuantidadeDeItens := 0.0
	for {
		for {
			fmt.Println("Codigo do produto:  - Digite 0 para finalizar pedido")
			pegaCodigo, _ := reader.ReadString('\n')
			limpaCodigo := strings.TrimSpace(pegaCodigo)
			codigo, err := strconv.ParseFloat(limpaCodigo, 64)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
			atualizaCodigo = codigo
			break
		}

		for {
			fmt.Println("Quantidade de itens comprado:")
			pegaItens, _ := reader.ReadString('\n')
			limpaItem := strings.TrimSpace(pegaItens)
			quantidadeDeItens, err := strconv.ParseFloat(limpaItem, 64)
			if err != nil {
				fmt.Println("O que você digitou não é um número")
				continue
			}
			atualizaQuantidadeDeItens = quantidadeDeItens
			break
		}

		cachorroQuente := 100.00
		bauruSimples := 101.00
		bauruComOvo := 102.00
		hambúrguer := 103.00
		cheeseburguer := 104.00
		refrigerante := 105.00

		if atualizaCodigo == cachorroQuente {
			somaCompra = somaCompra + (atualizaQuantidadeDeItens * 1.20)
			valorPorItemCachorroQuente = valorPorItemCachorroQuente + atualizaQuantidadeDeItens
		}
		if atualizaCodigo == bauruSimples {
			somaCompra = somaCompra + (atualizaQuantidadeDeItens * 1.30)
			valorPorItemBauruSimples = valorPorItemBauruSimples + atualizaQuantidadeDeItens
		}
		if atualizaCodigo == bauruComOvo {
			somaCompra = somaCompra + (atualizaQuantidadeDeItens * 1.50)
			valorPorItemBauruComOvo = valorPorItemBauruComOvo + atualizaQuantidadeDeItens
		}
		if atualizaCodigo == hambúrguer {
			somaCompra = somaCompra + (atualizaQuantidadeDeItens * 1.20)
			valorPorItemHambúrguer = valorPorItemHambúrguer + atualizaQuantidadeDeItens
		}
		if atualizaCodigo == cheeseburguer {
			somaCompra = somaCompra + (atualizaQuantidadeDeItens * 1.30)
			valorPorItemCheeseburguer = valorPorItemCheeseburguer + atualizaQuantidadeDeItens
		}
		if atualizaCodigo == refrigerante {
			somaCompra = somaCompra + (atualizaQuantidadeDeItens * 1.00)
			valorPorItemRefrigerante = valorPorItemRefrigerante + atualizaQuantidadeDeItens
		}

		if atualizaCodigo == 0 {
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

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
	totalCompras := 0.0
	for {
		fmt.Println("Valor do item:")
		pegaI, _ := reader.ReadString('\n')
		limpaI := strings.TrimSpace(pegaI)
		valorItem, err := strconv.ParseFloat(limpaI, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		} else if valorItem > 0 {
			totalCompras = valorItem
		} else if valorItem == 0 {

			// pede troco
			pegaPagamento := 0.0
			for {
				fmt.Println("Dinheiro recebido:")
				dinheiro, _ := reader.ReadString('\n')
				limpaDinheiro := strings.TrimSpace(dinheiro)
				pagamento, err := strconv.ParseFloat(limpaDinheiro, 64)
				if err != nil {
					fmt.Println("O que você digitou não é um numero")
					continue
				}
				pegaPagamento = pagamento
				break
			}
			troco := pegaPagamento - totalCompras
			fmt.Println("Total de:", totalCompras)
			fmt.Println("Dinheiro recebido:", pegaPagamento)
			fmt.Println("troco de:", troco)
		} else {
			fmt.Println("Deseja encerrar? (S / N) ")
			pegaString, _ := reader.ReadString('\n')
			if strings.Contains(strings.ToUpper(pegaString), "S") {
				break
			}
		}
		// troco
	}
}

//Lojas Tabajara
//Produto 1: R$ 2.20
//Produto 2: R$ 5.80
//Produto 3: R$ 0
//Total: R$ 9.00
//Dinheiro: R$ 20.00
//Troco: R$ 11.00

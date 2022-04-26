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
	fmt.Println("Quantidade de carne comprada:")
	quant, _ := reader.ReadString('\n')
	limpaQuant := strings.TrimSpace(quant)
	quantidadeComprada, err := strconv.ParseFloat(limpaQuant, 64)
	if err != nil {
		fmt.Println("O que voce digitou não é um numero")
	}
	fmt.Println("Tipo de carne comprada:(File Duplo),(Alcatra),(Picanha)")
	tp, err := reader.ReadString('\n')
	tipoCarne := strings.TrimSpace(tp)
	tipoCarne = strings.ToUpper(tipoCarne)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}
	if tipoCarne == "FILE DUPLO" && quantidadeComprada < 5 {
		valorPagar := (quantidadeComprada * 4.90)
		desconto := (5 * valorPagar) / 100
		fazDesconto := valorPagar - desconto
		fmt.Println("Tipo da carne:", tipoCarne, "Quantidade de carne:", quantidadeComprada, "Preço total:", valorPagar, "Tipo de pagamento, se for cartão:", fazDesconto, "Se for em dinheiro:", valorPagar, "Valor do desconto:", desconto, "Total de:", valorPagar)
	} else if tipoCarne == "FILE DUPLO" && quantidadeComprada > 5 {
		valorPagar := (quantidadeComprada * 5.80)
		desconto := (5 * valorPagar) / 100
		fazDesconto := valorPagar - desconto
		fmt.Println("Tipo da carne:", tipoCarne, "Quantidade de carne:", quantidadeComprada, "Preço total:", valorPagar, "Tipo de pagamento, se for cartão:", fazDesconto, "Se for em dinheiro:", valorPagar, "Valor do desconto:", desconto, "Total de:", valorPagar)
	} else if tipoCarne == "ALCATRA" && quantidadeComprada < 5 {
		valorPagar := (quantidadeComprada * 5.90)
		desconto := (5 * valorPagar) / 100
		fazDesconto := valorPagar - desconto
		fmt.Println("Tipo da carne:", tipoCarne, "Quantidade de carne:", quantidadeComprada, "Preço total:", valorPagar, "Tipo de pagamento, se for cartão:", fazDesconto, "Se for em dinheiro:", valorPagar, "Valor do desconto:", desconto, "Total de:", valorPagar)
	} else if tipoCarne == "ALCATRA" && quantidadeComprada > 5 {
		valorPagar := (quantidadeComprada * 6.80)
		desconto := (5 * valorPagar) / 100
		fazDesconto := valorPagar - desconto
		fmt.Println("Tipo da carne:", tipoCarne, "Quantidade de carne:", quantidadeComprada, "Preço total:", valorPagar, "Tipo de pagamento, se for cartão:", fazDesconto, "Se for em dinheiro:", valorPagar, "Valor do desconto:", desconto, "Total de:", valorPagar)
	} else if tipoCarne == "PICANHA" && quantidadeComprada < 5 {
		valorPagar := (quantidadeComprada * 6.90)
		desconto := (5 * valorPagar) / 100
		fazDesconto := valorPagar - desconto
		fmt.Println("Tipo da carne:", tipoCarne, "Quantidade de carne:", quantidadeComprada, "Preço total:", valorPagar, "Tipo de pagamento, se for cartão:", fazDesconto, "Se for em dinheiro:", valorPagar, "Valor do desconto:", desconto, "Total de:", valorPagar)
	} else if tipoCarne == "PICANHA" && quantidadeComprada < 5 {
		valorPagar := (quantidadeComprada * 7.80)
		desconto := (5 * valorPagar) / 100
		fazDesconto := valorPagar - desconto
		fmt.Println("Tipo da carne:", tipoCarne, "Quantidade de carne:", quantidadeComprada, "Preço total:", valorPagar, "Tipo de pagamento, se for cartão:", fazDesconto, "Se for em dinheiro:", valorPagar, "Valor do desconto:", desconto, "Total de:", valorPagar)
	}
}

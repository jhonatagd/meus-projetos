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
	guardaCorte := 0
	valorTotal := 0
	for {
		fmt.Println("Escolha seu corte: 1- Blindado / 2- Raspar tudo / 3- Degradê / 4- Moicano.")
		pegaCorte, _ := reader.ReadString('\n')
		limpaCorte := strings.TrimSpace(pegaCorte)
		armazenaCorte, err := strconv.Atoi(limpaCorte)
		if err != nil {
			fmt.Println("O que você digitou não é um corte")
			continue
		}
		if armazenaCorte > 4 {
			fmt.Println("Corte invalido")
			continue
		}
		guardaCorte = armazenaCorte
		break
	}

	if guardaCorte == 1 {
		valorTotal = valorTotal + 25
		// blindado
	}
	if guardaCorte == 2 {
		valorTotal = valorTotal + 15
		//raspar tudo
	}
	if guardaCorte == 3 {
		valorTotal = valorTotal + 28
		// degrade
	}
	if guardaCorte == 4 {
		valorTotal = valorTotal + 23
		// moicano
	}
	armazenaExtra := 0
	for {
		fmt.Println("Serviços extras: 1- Sobrancelha / 2- Barba / 3- Sem extra.")
		pegaExtra, _ := reader.ReadString('\n')
		limpaExtra := strings.TrimSpace(pegaExtra)
		armazenaServicoExtra, err := strconv.Atoi(limpaExtra)
		if err != nil {
			fmt.Println("O que você digitou não é um Extra")
			continue
		}
		if armazenaServicoExtra > 3 {
			fmt.Println("Extra invalido")
			continue
		}
		armazenaExtra = armazenaServicoExtra
		break
	}

	if armazenaExtra == 1 {
		valorTotal = valorTotal + 5
		//sobrancelha
	}
	if armazenaExtra == 2 {
		valorTotal = valorTotal + 15
		//barba
	}
	if armazenaExtra == 3 {
		fmt.Println("Sem extras")
	}
	armazenaProduto := 0
	for {
		fmt.Println("Produtos de beleza Masculina: 1- Pomada modeladora / 2- Perfume / 3- Gel / 4- Minoxidil")
		pegaProdutos, _ := reader.ReadString('\n')
		limpaProdutos := strings.TrimSpace(pegaProdutos)
		armazenaProdutosExtra, err := strconv.Atoi(limpaProdutos)
		if err != nil {
			fmt.Println("O que você digitou não é um Produto")
			continue
		}
		if armazenaProdutosExtra > 4 {
			fmt.Println("Extra invalido")
			continue
		}
		armazenaProduto = armazenaProdutosExtra
		break
	}
	if armazenaProduto == 1 {
		valorTotal = valorTotal + 18
		//pomada
	}
	if armazenaProduto == 2 {
		valorTotal = valorTotal + 25
		//perfume
	}
	if armazenaProduto == 3 {
		valorTotal = valorTotal + 12
		//gel
	}
	if armazenaProduto == 4 {
		valorTotal = valorTotal + 50
		//minoxidil
	}

	fmt.Println("Valor total do seu dia de beleza:", valorTotal)

}

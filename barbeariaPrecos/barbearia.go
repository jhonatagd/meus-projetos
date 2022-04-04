package main

import (
	"fmt"
)

type Servico struct {
	nome  string
	valor float64
}

func main() {
	navalhado := Servico{nome: "navalhado", valor: 23.00}
	social := Servico{nome: "social", valor: 20.00}
	barba := Servico{nome: "barba", valor: 15.00}
	sobrancelha := Servico{nome: "sobrancelha", valor: 10.00}
	camisaBarbearia := Servico{nome: "camisa", valor: 40.00}

	var carrinho []Servico
	carrinho = append(carrinho, navalhado)
	carrinho = append(carrinho, social)
	carrinho = append(carrinho, barba)
	carrinho = append(carrinho, sobrancelha)
	carrinho = append(carrinho, camisaBarbearia)

	fmt.Printf("Carrinho:\n")
	var total float64
	for _, servico := range carrinho {
		total += servico.valor
		fmt.Printf("Nome: %v valor: %.2f - total parcial: %.2f\n", servico.nome, servico.valor, total)
	}
	fmt.Printf("Total: %.2f\n\n", total)

	if total >= 100 {
		totalComDesconto := total * 0.90
		fmt.Printf("Total com desconto: %.2f", totalComDesconto)
	} else {
		fmt.Printf("Total sem desconto: %.2f", total)
	}

}

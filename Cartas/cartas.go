package main

import (
	"fmt"

	"example.com/greetings/Cartas/funcoes"
	"example.com/greetings/Cartas/structs"
)

func main() {
	var baralho1 = structs.NovoBaralho("vermelho")
	//var baralho2 = structs.NovoBaralho("azul")

	fmt.Println(fmt.Sprintf("%v", baralho1))

	baralho1 = funcoes.EmbaralharByJhona(baralho1)
	fmt.Println(fmt.Sprintf("Embaralhado - %v", baralho1))

	baralho1 = funcoes.OrdenarValorACadaNaipe(baralho1)
	fmt.Println(fmt.Sprintf("Ordenar Valor A Cada Naipe - %v", baralho1))

	baralho1 = funcoes.OrdernarBaralhoNaipeEValor(baralho1)
	fmt.Println(fmt.Sprintf("Ordernar Baralho Naipe E Valor - %v", baralho1))

}

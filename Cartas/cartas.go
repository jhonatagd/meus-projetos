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

	baralho1 = funcoes.Embaralhar(baralho1)
	fmt.Println(fmt.Sprintf("Embaralhado - %v", baralho1))

	baralho1 = funcoes.OrdenarDaMenorCartaParaAMaior(baralho1)
	fmt.Println(fmt.Sprintf("ordenarDaMenorCartaParaAMaior - %v", baralho1))
}

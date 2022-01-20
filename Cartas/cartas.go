package main

import "fmt"

type carta struct {
	Nome  string
	Valor int
}

func main() {
	var cartas = []carta{
		{
			Nome:  "Espadas",
			Valor: 3,
		},
		{
			Nome:  "Paus",
			Valor: 5,
		},
		{
			Nome:  "Copas",
			Valor: 6,
		},
		{
			Nome:  "Ouros",
			Valor: 8,
		},
	}

	fmt.Println(fmt.Sprintf("%v", cartas))

	cartas = ordenarDaMenorCartaParaAMaior(cartas)
	fmt.Println(fmt.Sprintf("ordenarDaMenorCartaParaAMaior - %v", cartas))

}

func ordenarDaMenorCartaParaAMaior(cartas []carta) []carta {
	for i := 0; i < len(cartas); i++ {
		var cartaComMenorValor carta
		var indiceJ int = i

		cartaComMenorValor = cartas[i]

		for j := i; j < len(cartas); j++ {
			if cartaComMenorValor.Nome == cartas[j].Nome {
				continue
			} else {
				if cartaComMenorValor.Valor > cartas[j].Valor {
					cartaComMenorValor = cartas[j]
					indiceJ = j
				}
			}
		}

		if cartaComMenorValor.Nome != cartas[i].Nome {
			cartas[indiceJ] = cartas[i]
			cartas[i] = cartaComMenorValor
		}
	}
	return cartas
}

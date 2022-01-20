package funcoes

import (
	"math/rand"

	"example.com/greetings/Cartas/structs"
)

func Embaralhar(cartas []structs.Carta) []structs.Carta {
	for i := 0; i < 100; i++ {
		var posicao1 = rand.Intn(len(cartas))
		var posicao2 = rand.Intn(len(cartas))

		var variavelAuxiliar structs.Carta
		variavelAuxiliar = cartas[posicao1]

		cartas[posicao1] = cartas[posicao2]
		cartas[posicao2] = variavelAuxiliar
	}

	return cartas
}

func EmbaralharByJhona(cartas []structs.Carta) []structs.Carta {
	//Fazer um método de embaralhar utilizando outro algoritmo. Dica: Utilizar o rand.Intn ou alguma logica que vc pensar
	return cartas
}

func OrdernarBaralhoPorValorENaipe(cartas []structs.Carta) []structs.Carta {
	//Order de Naipes - Paus, Copas, Espadas, Ouro. Resultado => A Paus, A Copas, A Espadas, A Ouro, 2 Paus, 2 Copas, ....
	return cartas
}

func OrdernarBaralhoNaipeEValor(cartas []structs.Carta) []structs.Carta {
	//Order de Naipes - Paus, Copas, Espadas, Ouro. Resultado => A Paus, 2 Paus, 3 Paus, ..., A Copas, 2 Copas, 3 Copas ....
	return cartas
}

func OrdenarDaMenorCartaParaAMaior(cartas []structs.Carta) []structs.Carta {
	for i := 0; i < len(cartas); i++ {
		var cartaComMenorValor structs.Carta
		var indiceJ int = i

		cartaComMenorValor = cartas[i]

		for j := i; j < len(cartas); j++ {
			if cartaComMenorValor.Valor == cartas[j].Valor {
				continue
			} else {
				if cartaComMenorValor.Valor > cartas[j].Valor {
					cartaComMenorValor = cartas[j]
					indiceJ = j
				}
			}
		}

		if cartaComMenorValor.Valor != cartas[i].Valor {
			cartas[indiceJ] = cartas[i]
			cartas[i] = cartaComMenorValor
		}
	}

	return cartas
}

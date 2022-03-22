package funcoes

import (
	"math/rand"

	"example.com/greetings/Cartas/structs"
)

func Embaralhar(cartas []structs.Carta) []structs.Carta {
	for i := 0; i < 100; i++ {
		posicao1 := rand.Intn(len(cartas))
		posicao2 := rand.Intn(len(cartas))

		variavelAuxiliar := cartas[posicao1]

		cartas[posicao1] = cartas[posicao2]
		cartas[posicao2] = variavelAuxiliar
	}

	return cartas
}

func EmbaralharByJhona(cartas []structs.Carta) []structs.Carta {
	for i := 0; i < 100; i++ {
		var monteA []structs.Carta
		var monteB []structs.Carta
		var monteC []structs.Carta
		var indiceCartas int = (len(cartas)) / 3
		monteA = cartas[0:indiceCartas]
		monteB = cartas[indiceCartas : indiceCartas+indiceCartas]
		monteC = cartas[indiceCartas+indiceCartas:]
		cartas = append(monteB, monteC...)
		cartas = append(cartas, monteA...)
	}
	return cartas
}

func OrdenarValorACadaNaipe(baralho []structs.Carta) []structs.Carta {
	var valorA []structs.Carta
	var valor2 []structs.Carta
	var valor3 []structs.Carta
	var valor4 []structs.Carta
	var valor5 []structs.Carta
	var valor6 []structs.Carta
	var valor7 []structs.Carta
	var valor8 []structs.Carta
	var valor9 []structs.Carta
	var valor10 []structs.Carta
	var valorJ []structs.Carta
	var valorQ []structs.Carta
	var valorK []structs.Carta

	for _, valorSelecionado := range baralho {
		if valorSelecionado.Valor == "A" {
			valorA = append(valorA, valorSelecionado)
		}
		if valorSelecionado.Valor == "2" {
			valor2 = append(valor2, valorSelecionado)
		}
		if valorSelecionado.Valor == "3" {
			valor3 = append(valor3, valorSelecionado)
		}
		if valorSelecionado.Valor == "4" {
			valor4 = append(valor4, valorSelecionado)
		}
		if valorSelecionado.Valor == "5" {
			valor5 = append(valor5, valorSelecionado)
		}
		if valorSelecionado.Valor == "6" {
			valor6 = append(valor6, valorSelecionado)
		}
		if valorSelecionado.Valor == "7" {
			valor7 = append(valor7, valorSelecionado)
		}
		if valorSelecionado.Valor == "8" {
			valor8 = append(valor8, valorSelecionado)
		}
		if valorSelecionado.Valor == "9" {
			valor9 = append(valor9, valorSelecionado)
		}
		if valorSelecionado.Valor == "10" {
			valor10 = append(valor10, valorSelecionado)
		}
		if valorSelecionado.Valor == "J" {
			valorJ = append(valorJ, valorSelecionado)
		}
		if valorSelecionado.Valor == "K" {
			valorK = append(valorK, valorSelecionado)
		}
		if valorSelecionado.Valor == "Q" {
			valorQ = append(valorQ, valorSelecionado)
		}
	}

	valorA = OrdernarBaralhoPorNaipe(valorA)
	valor2 = OrdernarBaralhoPorNaipe(valor2)
	valor3 = OrdernarBaralhoPorNaipe(valor3)
	valor4 = OrdernarBaralhoPorNaipe(valor4)
	valor5 = OrdernarBaralhoPorNaipe(valor5)
	valor6 = OrdernarBaralhoPorNaipe(valor6)
	valor7 = OrdernarBaralhoPorNaipe(valor7)
	valor8 = OrdernarBaralhoPorNaipe(valor8)
	valor9 = OrdernarBaralhoPorNaipe(valor9)
	valor10 = OrdernarBaralhoPorNaipe(valor10)
	valorK = OrdernarBaralhoPorNaipe(valorK)
	valorJ = OrdernarBaralhoPorNaipe(valorJ)
	valorQ = OrdernarBaralhoPorNaipe(valorQ)

	var juntaBaralho []structs.Carta
	juntaBaralho = append(juntaBaralho, valorA...)
	juntaBaralho = append(juntaBaralho, valor2...)
	juntaBaralho = append(juntaBaralho, valor3...)
	juntaBaralho = append(juntaBaralho, valor4...)
	juntaBaralho = append(juntaBaralho, valor5...)
	juntaBaralho = append(juntaBaralho, valor6...)
	juntaBaralho = append(juntaBaralho, valor7...)
	juntaBaralho = append(juntaBaralho, valor8...)
	juntaBaralho = append(juntaBaralho, valor9...)
	juntaBaralho = append(juntaBaralho, valor10...)
	juntaBaralho = append(juntaBaralho, valorJ...)
	juntaBaralho = append(juntaBaralho, valorQ...)
	juntaBaralho = append(juntaBaralho, valorK...)

	return juntaBaralho
}

func OrdernarBaralhoPorNaipe(cartas []structs.Carta) []structs.Carta {
	var cartasDePaus []structs.Carta
	var cartasDeCopas []structs.Carta
	var cartasDeEspadas []structs.Carta
	var cartasDeOuros []structs.Carta
	for _, cartaSelecionada := range cartas {
		if cartaSelecionada.Naipe == "Paus" {
			cartasDePaus = append(cartasDePaus, cartaSelecionada)

		}
		if cartaSelecionada.Naipe == "Copas" {
			cartasDeCopas = append(cartasDeCopas, cartaSelecionada)

		}

		if cartaSelecionada.Naipe == "Espadas" {
			cartasDeEspadas = append(cartasDeEspadas, cartaSelecionada)
		}

		if cartaSelecionada.Naipe == "Ouros" {
			cartasDeOuros = append(cartasDeOuros, cartaSelecionada)
		}

	}
	var juntaBaralho []structs.Carta
	juntaBaralho = append(juntaBaralho, cartasDePaus...)
	juntaBaralho = append(juntaBaralho, cartasDeCopas...)
	juntaBaralho = append(juntaBaralho, cartasDeEspadas...)
	juntaBaralho = append(juntaBaralho, cartasDeOuros...)
	return juntaBaralho
}

func OrdernarBaralhoNaipeEValor(baralhoInicial []structs.Carta) []structs.Carta {
	var baralhoDePaus []structs.Carta
	var baralhoDeCopas []structs.Carta
	var baralhoDeEspadas []structs.Carta
	var baralhoDeOuros []structs.Carta
	for _, carta := range baralhoInicial {
		if carta.Naipe == "Paus" {
			baralhoDePaus = append(baralhoDePaus, carta)
		} else if carta.Naipe == "Copas" {
			baralhoDeCopas = append(baralhoDeCopas, carta)
		} else if carta.Naipe == "Espadas" {
			baralhoDeEspadas = append(baralhoDeEspadas, carta)
		} else {
			baralhoDeOuros = append(baralhoDeOuros, carta)
		}
	}

	var baralhoDePausOrdenado = ordenarBaralhoPorValor(baralhoDePaus)
	var baralhoDeCopasOrdenado = ordenarBaralhoPorValor(baralhoDeCopas)
	var baralhoDeEspadasOrdenado = ordenarBaralhoPorValor(baralhoDeEspadas)
	var baralhoDeOurosOrdenado = ordenarBaralhoPorValor(baralhoDeOuros)

	var juntaUltimosNaipes []structs.Carta
	juntaUltimosNaipes = append(juntaUltimosNaipes, baralhoDePausOrdenado...)
	juntaUltimosNaipes = append(juntaUltimosNaipes, baralhoDeCopasOrdenado...)
	juntaUltimosNaipes = append(juntaUltimosNaipes, baralhoDeEspadasOrdenado...)
	juntaUltimosNaipes = append(juntaUltimosNaipes, baralhoDeOurosOrdenado...)
	return juntaUltimosNaipes

}

func ordenarBaralhoPorValor(baralhoDeEntrada []structs.Carta) []structs.Carta {
	var baralhoDeA []structs.Carta
	var baralhoDe2 []structs.Carta
	var baralhoDe3 []structs.Carta
	var baralhoDe4 []structs.Carta
	var baralhoDe5 []structs.Carta
	var baralhoDe6 []structs.Carta
	var baralhoDe7 []structs.Carta
	var baralhoDe8 []structs.Carta
	var baralhoDe9 []structs.Carta
	var baralhoDe10 []structs.Carta
	var baralhoDeJ []structs.Carta
	var baralhoDeQ []structs.Carta
	var baralhoDeK []structs.Carta

	for _, carta := range baralhoDeEntrada {
		if carta.Valor == "A" {
			baralhoDeA = append(baralhoDeA, carta)
		}
		if carta.Valor == "2" {
			baralhoDe2 = append(baralhoDe2, carta)
		}
		if carta.Valor == "3" {
			baralhoDe3 = append(baralhoDe3, carta)
		}
		if carta.Valor == "4" {
			baralhoDe4 = append(baralhoDe4, carta)
		}
		if carta.Valor == "5" {
			baralhoDe5 = append(baralhoDe5, carta)
		}
		if carta.Valor == "6" {
			baralhoDe6 = append(baralhoDe6, carta)
		}
		if carta.Valor == "7" {
			baralhoDe7 = append(baralhoDe7, carta)
		}
		if carta.Valor == "8" {
			baralhoDe8 = append(baralhoDe8, carta)
		}
		if carta.Valor == "9" {
			baralhoDe9 = append(baralhoDe9, carta)
		}
		if carta.Valor == "10" {
			baralhoDe10 = append(baralhoDe10, carta)
		}
		if carta.Valor == "J" {
			baralhoDeJ = append(baralhoDeJ, carta)
		}
		if carta.Valor == "Q" {
			baralhoDeQ = append(baralhoDeQ, carta)
		}
		if carta.Valor == "K" {
			baralhoDeK = append(baralhoDeK, carta)
		}
	}

	var resultadoOrdenado []structs.Carta
	resultadoOrdenado = append(resultadoOrdenado, baralhoDeA...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe2...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe3...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe4...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe5...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe6...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe7...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe8...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe9...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDe10...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDeJ...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDeQ...)
	resultadoOrdenado = append(resultadoOrdenado, baralhoDeK...)

	return resultadoOrdenado
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

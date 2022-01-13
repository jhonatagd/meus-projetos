package main

import "fmt"

type Carro struct {
	Nome  string
	Valor float64
}

func main() {
	var carros = []Carro{
		Carro{
			Nome:  "Lamborghini",
			Valor: 500000,
		},
		Carro{
			Nome:  "Jipe",
			Valor: 46000,
		},
		Carro{
			Nome:  "Brasilia",
			Valor: 16000,
		},
		Carro{
			Nome:  "Smart",
			Valor: 46000,
		},
		Carro{
			Nome:  "Fusca",
			Valor: 17000,
		},
	}

	fmt.Println(fmt.Sprintf("%v", carros))

	carros = OrdernarPorValorDoMaisCaroParaOMaisBarato(carros)
	fmt.Println(fmt.Sprintf("%v", carros))
}

func OrdernarPorValorDoMaisCaroParaOMaisBarato(carros []Carro) []Carro {
	//len -> quantidade de itens em uma lista. Devolve um inteiro. Nesse exemplo, quanto é o len de carros 5
	for i := 0; i < len(carros); i++ {
		//Variavel auxiliar para salvar qual o é o carro mais caro
		var carroMaisCaro Carro
		//Variavel auxiliar para salvar a posição do carro mais caro
		var indiceJ int = i
		//Atribuir o carro da posição i como o carro mais caro
		carroMaisCaro = carros[i]

		for j := i; j < len(carros); j++ {
			if carroMaisCaro.Nome == carros[j].Nome {
				continue
			} else {
				if carroMaisCaro.Valor < carros[j].Valor {
					carroMaisCaro = carros[j]
					indiceJ = j
				}
			}
		}

		if carroMaisCaro.Nome != carros[i].Nome {
			carros[indiceJ] = carros[i]
			carros[i] = carroMaisCaro
		}
	}

	return carros
}

func OrdernarPorValorDoMaisBaratoParaOMaisCaro(carros []Carro) []Carro {
	//Coloque aqui seu codigo
	return carros
}

func OrdernarPorOrdemAlfabetica(carros []Carro) []Carro {
	//Coloque aqui seu codigo
	return carros
}

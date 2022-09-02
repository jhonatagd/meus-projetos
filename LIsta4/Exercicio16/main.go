//Utilize uma lista para resolver o problema a seguir. Uma empresa paga seus vendedores com base em comissões.
// O vendedor recebe $200 por semana mais 9 por cento de suas vendas brutas daquela semana. Por exemplo,
// um vendedor que teve vendas brutas de $3000 em uma semana recebe $200 mais 9 por cento de $3000, ou seja,
//um total de $470. Escreva um programa (usando um array de contadores) que determine quantos vendedores
// receberam salários nos seguintes intervalos de valores:
//$200 - $299
//$300 - $399
//$400 - $499
//$500 - $599
//$600 - $699
//$700 - $799
//$800 - $899
//$900 - $999
//$1000 em diante
//Desafio: Crie ma fórmula para chegar na posição da lista a partir do salário, sem fazer vários ifs aninhados.
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
	//pegaSemanasTrabalhadas := 0
	//pegaVendasBrutas := 0
	a := 0
	b := 0
	c := 0
	d := 0
	e := 0
	f := 0
	g := 0
	h := 0
	acimaK := 0

	salariosComputados := 0
	var listaTotalSemanaEValor []int
	for {
		fmt.Println("Semanas trabalhadas:")
		pegaSemanaTrabalhada, _ := reader.ReadString('\n')
		pegaSemanaInt := strings.TrimSpace(pegaSemanaTrabalhada)
		pegaSemana, err := strconv.Atoi(pegaSemanaInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		if pegaSemana == 0 {
			break
		}

		fmt.Println("vendas brutas:")
		pegaVendasBruto, _ := reader.ReadString('\n')
		pegaVendasint := strings.TrimSpace(pegaVendasBruto)
		pegaVendas, err := strconv.Atoi(pegaVendasint)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		salariosComputados += 1
		total := pegaSemana*200 + ((9 * pegaVendas) / 100) + pegaVendas
		listaTotalSemanaEValor = append(listaTotalSemanaEValor, total)

		for i := 0; i < salariosComputados; i++ {
			if listaTotalSemanaEValor[i] > 200 && listaTotalSemanaEValor[i] < 299 {
				a += 1
			} else if listaTotalSemanaEValor[i] > 300 && listaTotalSemanaEValor[i] < 399 {
				b += 1
			} else if listaTotalSemanaEValor[i] > 400 && listaTotalSemanaEValor[i] < 499 {
				c += 1
			} else if listaTotalSemanaEValor[i] > 500 && listaTotalSemanaEValor[i] < 599 {
				d += 1
			} else if listaTotalSemanaEValor[i] > 600 && listaTotalSemanaEValor[i] < 699 {
				e += 1
			} else if listaTotalSemanaEValor[i] > 700 && listaTotalSemanaEValor[i] < 799 {
				f += 1
			} else if listaTotalSemanaEValor[i] > 800 && listaTotalSemanaEValor[i] < 899 {
				g += 1
			} else if listaTotalSemanaEValor[i] > 900 && listaTotalSemanaEValor[i] < 999 {
				h += 1
			} else if listaTotalSemanaEValor[i] > 1000 {
				acimaK += 1
			}
		}
	}
	for i := 0; i < salariosComputados; i++ {
		if a > 0 {
			fmt.Println("salários entre 200 e 299:", a)
		}
		if b > 0 {
			fmt.Println("salários entre 300 e 399:", b)
		}
		if c > 0 {
			fmt.Println("salários entre 400 e 499:", c)
		}
		if d > 0 {
			fmt.Println("salários entre 500 e 599:", d)
		}
		if e > 0 {
			fmt.Println("salários entre 600 e 699:", e)
		}
		if f > 0 {
			fmt.Println("salários entre 700 e 799:", f)
		}
		if g > 0 {
			fmt.Println("salários entre 800 e 899:", g)
		}
		if h > 0 {
			fmt.Println("salários entre 900 e 999:", h)
		}
		if acimaK > 0 {
			fmt.Println("salários 1000 em diante:", acimaK)
		}
	}
	//fmt.Println(listaTotalSemanaEValor)
}

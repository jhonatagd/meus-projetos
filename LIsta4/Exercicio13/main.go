package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um programa que receba a temperatura média de cada mês do ano e armazene-as em uma lista.
// Após isto, calcule a média anual das temperaturas e mostre todas as temperaturas acima da média anual,
// e em que mês elas ocorreram (mostrar o mês por extenso: 1 – Janeiro, 2 – Fevereiro, . . . ).
func main() {
	reader := bufio.NewReader(os.Stdin)
	var mediaMes [12]float64
	somaDasMedias := 0.0
	for i := 0; i < 12; i++ {
		fmt.Println("Temperatura do", i+1, "mês:")
		pegaMes, _ := reader.ReadString('\n')
		limpaMes := strings.TrimSpace(pegaMes)
		mediaDosMeses, err := strconv.ParseFloat(limpaMes, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			i--
			continue
		}
		mediaMes[i] = mediaDosMeses
		somaDasMedias += mediaMes[i]
	}
	verificaMediaAnual := somaDasMedias / 12
	var mesesAcimaDaMedia []string
	fmt.Println(verificaMediaAnual)
	for i := 0; i < 12; i++ {
		mesUm := "Janeiro"
		mesDois := "Fevereiro"
		mesTres := "Março"
		mesQuatro := "Abril"
		mesCinco := "Maio"
		mesSeis := "Junho"
		mesSete := "Julho"
		mesOito := "Agosto"
		mesNove := "Setembro"
		mesDez := "Outubro"
		mesOnze := "Novembro"
		mesDoze := "Dezembro"
		if mediaMes[i] > verificaMediaAnual {
			if i == 0 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesUm)
			} else if i == 1 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesDois)
			} else if i == 2 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesTres)
			} else if i == 3 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesQuatro)
			} else if i == 4 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesCinco)
			} else if i == 5 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesSeis)
			} else if i == 6 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesSete)
			} else if i == 7 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesOito)
			} else if i == 8 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesNove)
			} else if i == 9 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesDez)
			} else if i == 10 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesOnze)
			} else if i == 11 {
				mesesAcimaDaMedia = append(mesesAcimaDaMedia, mesDoze)
			}
		}
	}

	fmt.Println("temperaturas acima da média anual", mesesAcimaDaMedia)
}

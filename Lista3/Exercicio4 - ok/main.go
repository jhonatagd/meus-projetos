package main

import "fmt"

func main() {
	habA := float64(80000)
	habB := float64(200000)
	taxaAF := float64((habA * 3) / 100)
	taxaBF := float64((habB * 1.5) / 100)
	anos := float64(0)

	for {
		anos = anos + 1.0
		habA = habA + (habA * taxaAF)
		habB = habB + (habB * taxaBF)
		if habA >= habB {
			fmt.Println("Serão necessarios", anos, "anos para que a população do pais A ultrapasse ou iguale a população do Pais B")
			return
		}

	}
	//Supondo que a população de um país A seja da ordem de 80000 habitantes com uma taxa anual
	//de crescimento de 3% e que a população de B seja 200000 habitantes com uma taxa de crescimento
	//de 1.5%. Faça um programa que calcule e escreva o número de anos necessários para que a população
	//do país A ultrapasse ou iguale a população do país B, mantidas as taxas de crescimento.
}

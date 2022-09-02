package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Faça um programa que leia um número indeterminado de valores, correspondentes a notas, encerrando a entrada
//de dados quando for informado um valor igual a -1 (que não deve ser armazenado). Após esta entrada de dados, faça:

func main() {
	reader := bufio.NewReader(os.Stdin)
	var listaValores []int
	//Mostre a quantidade de valores que foram lidos;
	quantidadeDeNumeros := 0
	for {
		fmt.Println("Numero indefinido:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaANumeroInt := strings.TrimSpace(pegaNumero)
		numero, err := strconv.Atoi(limpaANumeroInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else if numero == -1 {
			break
		}
		listaValores = append(listaValores, numero)
		quantidadeDeNumeros = quantidadeDeNumeros + 1
	}
	//Exiba todos os valores na ordem em que foram informados, um ao lado do outro;
	fmt.Println("Quantidade de numeros digitados", quantidadeDeNumeros)
	fmt.Println("Valores na ordem em que foram informados", listaValores)
	//Exiba todos os valores na ordem inversa à que foram informados, um abaixo do outro;
	fmt.Println("valores na ordem inversa:")
	for i := len(listaValores) - 1; i >= 0; i-- {
		fmt.Println(listaValores[i])
	}

	somaDosNumeros := 0

	//Calcule e mostre a soma dos valores;

	for i := 0; i < quantidadeDeNumeros; i++ {
		somaDosNumeros = somaDosNumeros + listaValores[i]
	}
	fmt.Println("Soma dos numeros", somaDosNumeros)

	//Calcule e mostre a média dos valores;
	mediaDosNumeros := somaDosNumeros / quantidadeDeNumeros
	fmt.Println("Media dos numeros", mediaDosNumeros)

	//Calcule e mostre a quantidade de valores acima da média calculada;
	quantidadeAcimaDaMedia := 0
	for i := 0; i < quantidadeDeNumeros; i++ {
		if listaValores[i] > mediaDosNumeros {
			quantidadeAcimaDaMedia += 1
		}
	}
	fmt.Println("Quantidade de valores acima da média calculada:", quantidadeAcimaDaMedia)

	//Calcule e mostre a quantidade de valores abaixo de sete;
	abaixode7 := 0
	for i := 0; i < quantidadeDeNumeros; i++ {
		if listaValores[i] < 7 {
			abaixode7 += 1
		}
	}
	fmt.Println("Quantidade de valores abaixo de sete:", abaixode7)

	//Encerre o programa com uma mensagem;
	fmt.Println("Finalmente fim.")

}

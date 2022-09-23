package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Faça um programa que carregue uma lista com os modelos de cinco carros (exemplo de modelos: FUSCA, GOL, VECTRA etc).
//Carregue uma outra lista com o consumo desses carros
// isto é, quantos quilômetros cada um desses carros faz com um litro de combustível.
//Calcule e mostre:
//O modelo do carro mais econômico;

func main() {
	reader := bufio.NewReader(os.Stdin)
	var todosOsMouses []string
	contaFor := 0
	problemasEsfera := 0.0
	problemasLimpeza := 0.0
	problemasCaboOuConector := 0.0
	problemasQuebradoOuInutilizado := 0.0
	for {
		fmt.Println("Situação do mouse: 1- esfera / limpeza / troca do cabo ou conector / quebrado ou inutilizado")
		problemaMouse, err := reader.ReadString('\n')
		pegaProblemaMouse := strings.TrimSpace(problemaMouse)
		pegaProblemaMouse = strings.ToUpper(pegaProblemaMouse)
		if err != nil {
			fmt.Println("O que você digitou não é um problema")
			continue
		}
		todosOsMouses = append(todosOsMouses, pegaProblemaMouse)
		contaFor += 1
		if pegaProblemaMouse == "0" {
			break
		}
	}

	for i := 0; i < contaFor; i++ {
		if todosOsMouses[i] == "ESFERA" {
			problemasEsfera += 1
		} else if todosOsMouses[i] == "LIMPEZA" {
			problemasLimpeza += 1
		} else if todosOsMouses[i] == "TROCA DO CABO OU CONECTOR" {
			problemasCaboOuConector += 1
		} else if todosOsMouses[i] == "QUEBRADO OU INUTILIZADO" {
			problemasQuebradoOuInutilizado += 1
		}

	}
	total := problemasEsfera + problemasLimpeza + problemasCaboOuConector + problemasQuebradoOuInutilizado

	fmt.Println("Situação                        Quantidade              Percentual")
	fmt.Println("1- necessita da esfera                 ", problemasEsfera, "    ", fmt.Sprintf("%.2f", (problemasEsfera*100)/total))
	fmt.Println("2- necessita de limpeza                ", problemasLimpeza, "   ", fmt.Sprintf("%.2f", (problemasLimpeza*100)/total))
	fmt.Println("3- necessita troca do cabo ou conector ", problemasCaboOuConector, "   ", fmt.Sprintf("%.2f", (problemasCaboOuConector*100)/total))
	fmt.Println("4- quebrado ou inutilizado             ", problemasQuebradoOuInutilizado, "   ", fmt.Sprintf("%.2f", (problemasQuebradoOuInutilizado*100)/total))

}

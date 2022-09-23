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
	nomeTarefa := " "
	somaPontos := 0.0

	for {
		fmt.Println("Digite o nome da tarefa:")
		pegaNome, err := reader.ReadString('\n')
		pegaPerguntaTarefa := strings.TrimSpace(pegaNome)
		pegaPerguntaTarefa = strings.ToUpper(pegaPerguntaTarefa)
		if err != nil {
			fmt.Println("O que você digitou não é um nome")
			continue
		}
		nomeTarefa = pegaPerguntaTarefa
		break
	}
	var listaPontuacao []float64
	//atualizaComplexidadeTarefa := 0

	for {
		fmt.Println("Nivel de complexidade da tarefa: 1-(P) 2-(M) 3-(G) 5 -(XG) 8-(GG)")
		pegaComplexidade, _ := reader.ReadString('\n')
		limpaComplexidade := strings.TrimSpace(pegaComplexidade)
		complexidade, err := strconv.ParseFloat(limpaComplexidade, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		listaPontuacao = append(listaPontuacao, complexidade)
		somaPontos += complexidade
		break
	}

	//atualizaTrabalhoTarefa := 0

	for {
		fmt.Println("Nivel de trabalho da tarefa: 1-(P) 2-(M) 3-(G) 5 -(XG) 8-(GG)")
		pegaTrabalho, _ := reader.ReadString('\n')
		limpaTrabalho := strings.TrimSpace(pegaTrabalho)
		trabalho, err := strconv.ParseFloat(limpaTrabalho, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		listaPontuacao = append(listaPontuacao, trabalho)
		somaPontos += trabalho
		break
	}
	//atualizaRiscosTarefa := 0

	for {
		fmt.Println("Nivel de risco da tarefa: 1-(P) 2-(M) 3-(G) 5 -(XG) 8-(GG)")
		pegaRisco, _ := reader.ReadString('\n')
		limpaRisco := strings.TrimSpace(pegaRisco)
		riscos, err := strconv.ParseFloat(limpaRisco, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		listaPontuacao = append(listaPontuacao, riscos)
		somaPontos += riscos
		break
	}

	//atualizaIncertezaTarefa := 0

	for {
		fmt.Println("Nivel de incerteza da tarefa: 1-(P) 2-(M) 3-(G) 5 -(XG) 8-(GG)")
		pegaIncerteza, _ := reader.ReadString('\n')
		limpaIncerteza := strings.TrimSpace(pegaIncerteza)
		incertezas, err := strconv.ParseFloat(limpaIncerteza, 64)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		}
		listaPontuacao = append(listaPontuacao, incertezas)
		somaPontos += incertezas
		break
	}
	dificuldadeDaTarefa := " "
	if somaPontos/4 < 2 {
		dificuldadeDaTarefa = "Simples"
	} else if somaPontos/4 > 2 && somaPontos/4 < 4 {
		dificuldadeDaTarefa = "Mediana"
	} else if somaPontos/4 > 4 {
		dificuldadeDaTarefa = "dificil"
	}

	fmt.Println("Tarefa:", nomeTarefa)
	fmt.Println(fmt.Sprintf("%.2f", (listaPontuacao[0]*100)/somaPontos), "%", "de complexidade")
	fmt.Println(fmt.Sprintf("%.2f", (listaPontuacao[1]*100)/somaPontos), "%", "de trabalho")
	fmt.Println(fmt.Sprintf("%.2f", (listaPontuacao[2]*100)/somaPontos), "%", "de riscos")
	fmt.Println(fmt.Sprintf("%.2f", (listaPontuacao[3]*100)/somaPontos), "%", "de incertezas")
	fmt.Println("Tarefa", dificuldadeDaTarefa)
}

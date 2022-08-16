package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// um programa que peça a linha e a coluna
// duas pessoas
// um X e outro O
// a cada vez que selecionar uma linha e tiver preenchido o jogo deve colocar
// como invalido e pedir de novo
// e a cada vez que preencher uma coluna e linha o sistema verifica se o jogo acabou
// e informar quem ganhou
// pedir players
// usar matriz
// sempre que digitar uma posição, imprimir o jogo atual
func main() {
	reader := bufio.NewReader(os.Stdin)
	nomeJogador1 := ""
	nomeJogador2 := ""
	for {
		fmt.Println("Nome do primeiro jogador:")
		pegaNomeJogador1, err := reader.ReadString('\n')
		pegaNomeJogador1String := strings.TrimSpace(pegaNomeJogador1)
		pegaNomeJogador1String = strings.ToUpper(pegaNomeJogador1String)
		if err != nil {
			fmt.Println("O que você digitou não é um nome")
			continue
		} else if pegaNomeJogador1String == "" {
			fmt.Println("O que você digitou não é um nome")
			continue
		}
		nomeJogador1 = pegaNomeJogador1String
		break
	}

	for {
		fmt.Println("Nome do segundo jogador:")
		pegaNomeJogador2, err := reader.ReadString('\n')
		pegaNomeJogador2String := strings.TrimSpace(pegaNomeJogador2)
		pegaNomeJogador2String = strings.ToUpper(pegaNomeJogador2String)
		if err != nil {
			fmt.Println("O que você digitou não é um nome")
			continue
		} else if pegaNomeJogador2String == "" {
			fmt.Println("O que você digitou não é um nome")
			continue
		}
		nomeJogador2 = pegaNomeJogador2String
		break
	}
	var jogoDaVelha [3][3]string

	coluna := 0
	linha := 0
	for i := 0; i < 9; i++ {
		nomeJogadorAtivo := ""
		valorJogadorAtivo := ""

		verificaImparPar := i % 2
		if verificaImparPar == 0 {
			nomeJogadorAtivo = nomeJogador1
			valorJogadorAtivo = "X"
		} else {
			nomeJogadorAtivo = nomeJogador2
			valorJogadorAtivo = "O"
		}
		for {
			fmt.Println(nomeJogadorAtivo, "Escolha a linha de inicio: [1][2][3]")
			pegaLinha, _ := reader.ReadString('\n')
			limpaLinha := strings.TrimSpace(pegaLinha)
			pegaLinhaJogador, err := strconv.Atoi(limpaLinha)
			if err != nil {
				fmt.Println("O que você digitou não é um numero")
				continue
			} else if pegaLinhaJogador > 3 || pegaLinhaJogador < 1 {
				fmt.Println("Não é um numero valido")
				continue
			}
			linha = pegaLinhaJogador - 1
			break
		}
		// [] [] []
		// [] [] []
		// [] [x] []
		for {
			fmt.Println(nomeJogadorAtivo, "Escolha a coluna da linha: [1][2][3]")
			pegaColuna, _ := reader.ReadString('\n')
			limpaColuna := strings.TrimSpace(pegaColuna)
			pegaColunaJogador, err := strconv.Atoi(limpaColuna)
			if err != nil {
				fmt.Println("O que você digitou não é um numero")
				continue
			} else if pegaColunaJogador > 3 || pegaColunaJogador < 1 {
				fmt.Println("Não é um numero valido")
				continue
			}
			coluna = pegaColunaJogador - 1
			break
		}

		if jogoDaVelha[linha][coluna] == "" {
			jogoDaVelha[linha][coluna] = valorJogadorAtivo
			ImprimirTabela(jogoDaVelha)
			recebeString := QuemGanhou2(jogoDaVelha)
			if recebeString == "X" {
				fmt.Println("Jogador", nomeJogador1, "ganhou")
				return
			}
			if recebeString == "O" {
				fmt.Println("Jogador", nomeJogador2, "ganhou")
				return
			}
		} else {
			fmt.Println("Escolha outra posição")
			ImprimirTabela(jogoDaVelha)
			i = i - 1
			continue
		}

	}
	fmt.Println("Deu velha")
}

func ImprimirTabela(matriz [3][3]string) {
	for l := 0; l < 3; l++ {
		desenho := ""
		for c := 0; c < 3; c++ {
			desenho += matriz[l][c] + "|"
		}
		fmt.Println(desenho)
	}

}

func QuemGanhou(matriz [3][3]string) string {
	if matriz[0][0] == "X" && matriz[0][1] == "X" && matriz[0][2] == "X" {
		return "X"
	} else if matriz[0][0] == "X" && matriz[1][1] == "X" && matriz[2][2] == "X" {
		return "X"
	} else if matriz[0][0] == "X" && matriz[1][0] == "X" && matriz[2][0] == "X" {
		return "X"
	} else if matriz[1][0] == "X" && matriz[1][1] == "X" && matriz[1][2] == "X" {
		return "X"
	} else if matriz[2][0] == "X" && matriz[1][1] == "X" && matriz[0][2] == "X" {
		return "X"
	} else if matriz[0][1] == "X" && matriz[1][1] == "X" && matriz[2][1] == "X" {
		return "X"
	} else if matriz[2][0] == "X" && matriz[2][1] == "X" && matriz[2][2] == "X" {
		return "X"
	} else if matriz[0][2] == "X" && matriz[1][2] == "X" && matriz[2][2] == "X" {
		return "X"
	} else if matriz[0][0] == "O" && matriz[0][1] == "O" && matriz[0][2] == "O" {
		return "O"
	} else if matriz[0][0] == "O" && matriz[1][1] == "O" && matriz[2][2] == "O" {
		return "O"
	} else if matriz[0][0] == "O" && matriz[1][0] == "O" && matriz[2][0] == "O" {
		return "O"
	} else if matriz[1][0] == "O" && matriz[1][1] == "O" && matriz[1][2] == "O" {
		return "O"
	} else if matriz[1][0] == "O" && matriz[1][1] == "O" && matriz[0][2] == "O" {
		return "O"
	} else if matriz[0][1] == "O" && matriz[1][1] == "O" && matriz[2][1] == "O" {
		return "O"
	} else if matriz[2][0] == "O" && matriz[2][1] == "O" && matriz[2][2] == "O" {
		return "O"
	} else if matriz[1][2] == "O" && matriz[1][2] == "O" && matriz[2][2] == "O" {
		return "O"
	} else {
		return ""
	}
}

func QuemGanhou2(matriz [3][3]string) string {
	if matriz[0][0] == matriz[0][1] && matriz[0][0] == matriz[0][2] {
		return matriz[0][0]
	} else if matriz[0][0] == matriz[1][1] && matriz[0][0] == matriz[2][2] {
		return matriz[0][0]
	} else if matriz[0][0] == matriz[1][0] && matriz[0][0] == matriz[2][0] {
		return matriz[0][0]
	} else if matriz[1][0] == matriz[1][1] && matriz[1][0] == matriz[1][2] {
		return matriz[1][0]
	} else if matriz[2][0] == matriz[1][1] && matriz[2][0] == matriz[0][2] {
		return matriz[2][0]
	} else if matriz[0][1] == matriz[1][1] && matriz[0][1] == matriz[2][1] {
		return matriz[0][1]
	} else if matriz[2][0] == matriz[2][1] && matriz[0][1] == matriz[2][2] {
		return matriz[2][0]
	} else if matriz[0][2] == matriz[1][2] && matriz[0][1] == matriz[2][2] {
		return matriz[0][2]
	} else {
		return ""
	}
}

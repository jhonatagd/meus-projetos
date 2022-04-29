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

	for {
		fmt.Println("Informe o aluno: ")
		aluno, _ := reader.ReadString('\n')
		aluno = strings.TrimSpace(aluno)

		fmt.Println("Informe a materia - Português (P) ou Matemática (M): ")
		materia, _ := reader.ReadString('\n')
		materia = strings.TrimSpace(materia)
		materia = strings.ToUpper(materia)

		if materia != "P" && materia != "M" {
			fmt.Println("O que você digitou não é um valor válido")
			break
		}

		fmt.Println("Informe uma nota de 0 a 10: ")
		nota, _ := reader.ReadString('\n')
		nota = strings.TrimSpace(nota)
		notaInt, err := strconv.Atoi(nota)
		if err != nil {
			fmt.Println("O que você digitou não é um valor válido")
			break
		}

		if notaInt < 0 || notaInt > 10 {
			fmt.Println("O que você digitou não é um valor válido")
			break
		}

		fmt.Println("Informe o bimestre (1B, 2B, 3B ou 4B): ")
		bimestre, _ := reader.ReadString('\n')
		bimestre = strings.TrimSpace(bimestre)
		bimestre = strings.ToUpper(bimestre)

		if bimestre != "1B" && bimestre != "2B" && bimestre != "3B" && bimestre != "4B" {
			fmt.Println("O que você digitou não é um valor válido")
			break
		}

		fileName := "Lista3/ArmazenarInformacao/DB.txt"
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Erro ao tentar abrir arquivo", err)
			break
		}

		textoParaSalvar := fmt.Sprintf("%s-%s-%d-%s", aluno, materia, notaInt, bimestre)

		_, err = file.WriteString(textoParaSalvar)
		if err != nil {
			fmt.Println("Erro ao tentar escrever no arquivo")
			break
		}

		file.Close()
	}
}

// 2
// 2 ,3 ,4, 12

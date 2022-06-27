package main

//fazer um programa que pergunte a cada um dos clientes da academia seu código, sua altura e seu peso.
// O final da digitação de dados deve ser dada quando o usuário digitar 0 (zero) no campo código.
// Ao encerrar o programa também deve ser informados os códigos e valores do clente mais alto,
//do mais baixo, do mais gordinho e do mais magro, além da média das alturas e dos pesos dos clientes
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	guardaMaisAlto := 0
	guardaCodigoMaisAlto := 0
	guardaMaisBaixo := 2147483647
	guardaCodigoMaisBaixo := 0
	guardaMaisMagro := 2147483647
	guardaCodigoMaisMagro := 0
	guardaMaisGordinho := 0
	guardaCodigoMaisGordinho := 0
	mediaAltura := 0
	mediaPeso := 0
	divisor := 0
	for {
		fmt.Println("Seu codigo:")
		pegaCodigo, _ := reader.ReadString('\n')
		limpaCodigo := strings.TrimSpace(pegaCodigo)
		codigo, err := strconv.Atoi(limpaCodigo)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
		} else if codigo == 0 {
			fmt.Println("Codigo encerrado.")
			break
		}
		fmt.Println("Sua altura:")
		pegaAltura, _ := reader.ReadString('\n')
		limpaAltura := strings.TrimSpace(pegaAltura)
		altura, err := strconv.Atoi(limpaAltura)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
		}
		fmt.Println("Seu Peso:")
		pegaPeso, _ := reader.ReadString('\n')
		limpaPeso := strings.TrimSpace(pegaPeso)
		peso, err := strconv.Atoi(limpaPeso)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
		}
		divisor = divisor + 1
		mediaAltura = (mediaAltura + altura) / divisor
		mediaPeso = (mediaPeso + peso) / divisor

		if altura > guardaMaisAlto {
			guardaMaisAlto = altura
			guardaCodigoMaisAlto = codigo
		}
		if peso > guardaMaisGordinho {
			guardaMaisGordinho = peso
			guardaCodigoMaisGordinho = codigo
		}
		if altura < guardaMaisBaixo {
			guardaMaisBaixo = altura
			guardaCodigoMaisBaixo = codigo
		}
		if peso < guardaMaisMagro {
			guardaMaisMagro = peso
			guardaCodigoMaisMagro = codigo

		}

	}
	fmt.Println("Codigo do mais baixo:", guardaCodigoMaisBaixo, ", altura:", guardaMaisBaixo)
	fmt.Println("Codigo do mais alto:", guardaCodigoMaisAlto, ", altura:", guardaMaisAlto)
	fmt.Println("Codigo do mais magro:", guardaCodigoMaisMagro, ", peso:", guardaMaisMagro)
	fmt.Println("Codigo do mais gordinho:", guardaCodigoMaisGordinho, ", peso:", guardaMaisGordinho)
	fmt.Printf("Media dos pesos: %v Media das alturas: %v", mediaPeso, mediaAltura)

}

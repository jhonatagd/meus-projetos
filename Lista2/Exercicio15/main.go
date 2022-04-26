package main

//  peça os 3 lados de um triângulo
// O programa deverá informar se os valores podem ser um triângulo. Indique, caso os lados formem um triângulo,
// se o mesmo é: equilátero, isósceles ou escaleno.
// Três lados formam um triângulo quando a soma de quaisquer dois lados for maior que o terceiro;
// Triângulo Equilátero: três lados iguais;
// Triângulo Isósceles: quaisquer dois lados iguais;
// Triângulo Escaleno: três lados diferentes;
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Valor do Primeiro lado do triangulo:")
	ladoUm, _ := reader.ReadString('\n')
	limpaLadoum := strings.TrimSpace(ladoUm)
	ladoUmFloat, err := strconv.Atoi(limpaLadoum)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	fmt.Println("Valor do segundo lado do triangulo:")
	ladoDois, _ := reader.ReadString('\n')
	limpaLadoDois := strings.TrimSpace(ladoDois)
	ladoDoisFloat, err := strconv.Atoi(limpaLadoDois)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	fmt.Println("Valor do terceiro lado do triangulo:")
	ladoTres, _ := reader.ReadString('\n')
	limpaLadoTres := strings.TrimSpace(ladoTres)
	ladoTresFloat, err := strconv.Atoi(limpaLadoTres)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	if ladoUmFloat == ladoDoisFloat && ladoUmFloat == ladoTresFloat {
		fmt.Println("Triangulo equilátero")
	} else if (ladoUmFloat == ladoDoisFloat && ladoUmFloat != ladoTresFloat) || (ladoUmFloat == ladoTresFloat && ladoUmFloat != ladoDoisFloat) || (ladoDoisFloat == ladoTresFloat && ladoDoisFloat != ladoUmFloat) {
		fmt.Println("Triangulo Isósceles")
	} else {
		fmt.Println("Triangulo escaleno")
	}
}

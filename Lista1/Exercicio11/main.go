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
	fmt.Println("numero inteiro um: ")
	pegaIntUm, _ := reader.ReadString('\n')
	limpaIntUm := strings.TrimSpace(pegaIntUm)
	intUmFloat, err := strconv.Atoi(limpaIntUm)
	if err != nil {
		fmt.Println("O que você digitou não é um numero inteiro")
	}
	fmt.Println("numero inteiro dois: ")
	pegaIntdois, _ := reader.ReadString('\n')
	limpaIntdois := strings.TrimSpace(pegaIntdois)
	intdoisFloat, err := strconv.Atoi(limpaIntdois)
	if err != nil {
		fmt.Println("O que você digitou não é um numero inteiro")
	}
	fmt.Println("numero natural: ")
	pegaReal, _ := reader.ReadString('\n')
	limpaReal := strings.TrimSpace(pegaReal)
	pegaRealFloat, err := strconv.ParseFloat(limpaReal, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um numero real")
	}
	a := (intUmFloat * 2.0) * (intdoisFloat / 2.0)
	b := float64(intUmFloat*3) + pegaRealFloat
	c := pegaRealFloat * pegaRealFloat * pegaRealFloat

	fmt.Println("Resultado a: ")
	fmt.Println(a)
	fmt.Println("Resultado b: ")
	fmt.Println(b)
	fmt.Println("Resultado c: ")
	fmt.Println(c)

}

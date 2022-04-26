package main

//Faça um Programa que leia um número inteiro menor que 1000 e imprima a quantidade de centenas, dezenas e unidades do mesmo.
//Observando os termos no plural a colocação do "e", da vírgula entre outros. Exemplo:
//326 = 3 centenas, 2 dezenas e 6 unidades
//12 = 1 dezena e 2 unidades Testar com: 326, 300, 100, 320, 310,305, 301, 101, 311, 111, 25, 20, 10, 21, 11, 1, 7 e 16
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite um numero menor que 1000:")
	numero, _ := reader.ReadString('\n')
	limpaN := strings.TrimSpace(numero)
	numeroInt, err := strconv.Atoi(limpaN)
	if err != nil {
		fmt.Println("O que você digitou não é um numero")
	}
	// pegando a unidade
	unidade := numeroInt % 10
	// tentando eliminar a unidade
	numeroInt = (numeroInt - unidade) / 10
	// pegando a ddezena
	dezena := numeroInt % 10
	// pegando o meu numero tirando a dezena e deixando a centena
	numeroInt = (numeroInt - dezena) / 10
	centena := (numeroInt)

	if numeroInt <= 1000 {
		fmt.Println(centena, "centena,", dezena, "dezena, e ", unidade, "unidade.")
	}
}

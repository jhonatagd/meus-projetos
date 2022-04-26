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
	fmt.Println("Valor do saque:")
	saque, _ := reader.ReadString('\n')
	limpaSaque := strings.TrimSpace(saque)
	numero, err := strconv.Atoi(limpaSaque)
	if err != nil {
		fmt.Println("O que voce digitou não é um valor")
	} else if numero < 10 {
		fmt.Println("Saque minimo 10 Reais")
		return
	} else if numero > 600 {
		fmt.Println("Saque maximo 600 Reais")
		return
	}

	cem := numero / 100
	numero = numero - (cem * 100)
	cinquenta := numero / 50
	numero = numero - (cinquenta * 50)
	dez := numero / 10
	numero = numero - (dez * 10)
	cinco := numero / 5
	numero = numero - (cinco * 5)
	um := numero

	if numero > 10 && numero < 600 {
		fmt.Println(cem, "Nota(s) de R$100,00", cinquenta, "Nota(s) de R$50,00", dez, "Nota(s) de R$10,00", cinco, "Nota(s) de R$5,00", um, "Nota(s) de R$1,00")
	}
}

// programa deverá perguntar ao usuário a valor do saque
// depois informar quantas notas de cada valor serão fornecidas.
// notas disponíveis serão as de 1, 5, 10, 50 e 100 reais.
// valor mínimo é de 10 reais e o máximo de 600 reais.
// O programa não deve se preocupar com a quantidade de notas existentes na máquina.

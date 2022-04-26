package main

// calcule as raízes de uma equação do segundo grau, na forma ax2 + bx + c.
// O programa deverá pedir os valores de a, b e c e fazer as consistências
// se informar o valor de A = a 0, a equação não é do segundo grau e o programa não deve fazer pedir os demais valores, sendo encerrado;
// Se o delta calculado for negativo, a equação não possui raizes reais. Informe ao usuário e encerre o programa;
// Se o delta calculado for igual a zero a equação possui apenas uma raiz real; informe-a ao usuário;
// Se o delta for positivo, a equação possui duas raiz reais; informe-as ao usuário;
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Valor de A: ")
	a, _ := reader.ReadString('\n')
	limpaA := strings.TrimSpace(a)
	aFloat, err := strconv.ParseFloat(limpaA, 64)
	if aFloat == 0 {
		fmt.Println("Se A = 0, não é equação de segundo grau.")
		return
	}

	fmt.Println("Valor de B: ")
	b, _ := reader.ReadString('\n')
	limpaB := strings.TrimSpace(b)
	bFloat, err := strconv.ParseFloat(limpaB, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	fmt.Println("Valor de C: ")
	c, _ := reader.ReadString('\n')
	limpaC := strings.TrimSpace(c)
	cFloat, err := strconv.ParseFloat(limpaC, 64)
	if err != nil {
		fmt.Println("O que você digitou não é um valor")
	}
	delta := (bFloat * bFloat) - (4 * (aFloat * cFloat))
	if delta < 0 {
		fmt.Println("Não possui raizes reais ")
		return
	} else if delta == 0 {
		fmt.Println("Possui apenas uma raiz real", (-bFloat)/2*aFloat)
	} else if delta > 0 {
		raizDeDelta := math.Sqrt(delta)
		fmt.Println("Equação possui duas raizes, Positiva:", ((-bFloat + raizDeDelta) / (2 * aFloat)), "E negativa: ", ((-bFloat - raizDeDelta) / (2 * aFloat)))
	}
}

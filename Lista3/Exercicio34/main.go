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
	pegaNumerosPrimos := 0
	for {
		fmt.Printf("Um número inteiro:")
		pegaNumero, _ := reader.ReadString('\n')
		limpaNumero := strings.TrimSpace(pegaNumero)
		numeroInteiro, err := strconv.Atoi(limpaNumero)
		if err != nil {
			fmt.Println("O que você digitou não é um numero")
			continue
		}
		pegaNumerosPrimos = numeroInteiro
		break
	}
	if pegaNumerosPrimos %1 == 0 && pegaNumerosPrimos %pegaNumerosPrimos== 0 {
		fmt.Println("É um numero primo ")
	}
}
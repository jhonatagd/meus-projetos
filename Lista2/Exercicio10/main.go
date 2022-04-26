package main

//pergunte em que turno você estuda.
//Peça para digitar M-matutino ou V-Vespertino ou N- Noturno.
// Imprima a mensagem "Bom Dia!", "Boa Tarde!" ou "Boa Noite!" ou "Valor Inválido!", conforme o caso.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("em que turno você estuda (M-matutino ou V-Vespertino ou N- Noturno):")
	pegaLetra, err := reader.ReadString('\n')
	pegaLetraString := strings.TrimSpace(pegaLetra)
	pegaLetraString = strings.ToUpper(pegaLetraString)
	if err != nil {
		fmt.Println("O que você digitou não é uma letra")
	}
	manha := "M"
	vespertino := "V"
	noturno := "N"

	if pegaLetraString == manha {
		fmt.Println("Bom Dia!")
	} else if pegaLetraString == vespertino {
		fmt.Println("Boa Tarde!")
	} else if pegaLetraString == noturno {
		fmt.Println("Boa Noite!")
	} else {
		fmt.Println("Valor Inválido!")
	}
}

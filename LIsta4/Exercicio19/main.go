package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func main() {
	reader := bufio.NewReader(os.Stdin)
	melhorSistema := 0
	listaWindowsServe := 0
	listaUnix := 0
	listaLinux := 0
	listaNetware := 0
	listaMacOS := 0
	listaOutro := 0

	for {
		fmt.Println("Qual o melhor Sistema Operacional para uso em servidores?")
		fmt.Println("1- Windows Server 2- Unix 3- Linux 4- Netware 5- Mac OS 6- Outro")
		pegaResposta, _ := reader.ReadString('\n')
		pegaRespostaInt := strings.TrimSpace(pegaResposta)
		pegaMelhorSistema, err := strconv.Atoi(pegaRespostaInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else if pegaMelhorSistema == 0 {
			break
		} else if pegaMelhorSistema > 6 {
			fmt.Println("Valor invalido")
		}
		melhorSistema = pegaMelhorSistema
		total := 0
		if melhorSistema == 1 {
			listaWindowsServe += 1
			total += 1
		} else if melhorSistema == 2 {
			listaUnix += 1
			total += 1
		} else if melhorSistema == 3 {
			listaLinux += 1
			total += 1
		} else if melhorSistema == 4 {
			listaNetware += 1
			total += 1
		} else if melhorSistema == 5 {
			listaMacOS += 1
			total += 1
		} else if melhorSistema == 6 {
			listaOutro += 1
			total += 1
		}
		//var listaPorcentagem {}int

	}
	fmt.Println("Sistema Operacional     Votos   %")
	fmt.Println("Windows Server", listaWindowsServe)
	fmt.Println("Unix", listaUnix)
	fmt.Println("Linux", listaLinux)
	fmt.Println("Netware", listaNetware)
	fmt.Println("Mac OS", listaMacOS)
	fmt.Println("Outro", listaOutro)

}
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	var listaSO [6]float64

	for {
		fmt.Println("Qual o melhor Sistema Operacional para uso em servidores?")
		fmt.Println("1- Windows Server 2- Unix 3- Linux 4- Netware 5- Mac OS 6- Outro")
		pegaResposta, _ := reader.ReadString('\n')
		pegaRespostaInt := strings.TrimSpace(pegaResposta)
		pegaMelhorSistema, err := strconv.Atoi(pegaRespostaInt)
		if err != nil {
			fmt.Println("O que você digitou não é um número")
			continue
		} else if pegaMelhorSistema == 0 {
			break
		} else if pegaMelhorSistema > 6 {
			fmt.Println("Valor invalido")
			continue
		}
		listaSO[pegaMelhorSistema-1] += 1

	}
	total := 0.0
	for i := 0; i < len(listaSO); i++ {
		total += listaSO[i]
	}

	fmt.Println("Sistema Operacional     Votos   %")
	fmt.Println("Windows Server", listaSO[0], fmt.Sprintf("%.2f", (listaSO[0]*100)/total))
	fmt.Println("Unix", listaSO[1], fmt.Sprintf("%.2f", (listaSO[1]*100)/total))
	fmt.Println("Linux", listaSO[2], fmt.Sprintf("%.2f", (listaSO[2]*100)/total))
	fmt.Println("Netware", listaSO[3], fmt.Sprintf("%.2f", (listaSO[3]*100)/total))
	fmt.Println("Mac OS", listaSO[4], fmt.Sprintf("%.2f", (listaSO[4]*100)/total))
	fmt.Println("Outro", listaSO[5], fmt.Sprintf("%.2f", (listaSO[5]*100)/total))

}

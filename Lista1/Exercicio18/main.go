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
	fmt.Println(" tamanho de um arquivo para download (em MB): ")
	mB, _ := reader.ReadString('\n')
	limpaMB := strings.TrimSpace(mB)
	mBFloat, err := strconv.Atoi(limpaMB)
	if err != nil {
		fmt.Println("O que você digitou não é um MB")
	}
	fmt.Println("velocidade de um link de Internet (em Mbps): ")
	mBPS, _ := reader.ReadString('\n')
	limpaMBPS := strings.TrimSpace(mBPS)
	mBFPSloat, err := strconv.Atoi(limpaMBPS)
	if err != nil {
		fmt.Println("O que você digitou não é um MB")
	}
	segundos := mBFloat / mBFPSloat
	minutos := segundos / 60
	seg := segundos % 60

	fmt.Println("Tempo estimado para download", minutos, "e", seg, "segundos")

}

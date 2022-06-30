package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Nome de usuario: ")
		nome, err := reader.ReadString('\n')
		nomeUsuario := strings.TrimSpace(nome)
		nomeUsuario = strings.ToUpper(nomeUsuario)
		if err != nil {
			fmt.Println("O que você digitou não é um valor válido")
		}
		fmt.Println("Senha do usuario: ")
		senha, err := reader.ReadString('\n')
		senhaUsuario := strings.TrimSpace(senha)
		senhaUsuario = strings.ToUpper(senhaUsuario)
		if err != nil {
			fmt.Println("O que você digitou não é um valor válido")
		}

		if nomeUsuario == senhaUsuario {
			fmt.Println("erro, escreva as informações novamente.")
		} else {
			fmt.Println("Senha e usuario ok.")
			return
		}
	}

	// Faça um programa que leia um nome de usuário e a sua senha e não aceite a senha igual ao nome do usuário
	// mostrando uma mensagem de erro e voltando a pedir as informações.
}

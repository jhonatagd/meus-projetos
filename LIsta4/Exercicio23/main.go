package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {
	arquivoUsuario, err := os.Open("usuarios.txt")
	if err != nil {
		fmt.Println("Erro ao tentar abrir arquivo")
		return
	}

	defer arquivoUsuario.Close()

	reader := bufio.NewReader(arquivoUsuario)
	//buf := make([]byte, 16)

	for {
		bytes, _, err := reader.ReadLine()

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				continue
			}

			break
		}

		//linha := string(bytes)
		fmt.Println(binary.Size(bytes), " bytes")

		//Criar um arquivo "relatório.txt"
		//Escrever no arquivo de acordo com o formato exigido
	}

	// reader := bufio.NewReader(os.Stdin)
	// var listaDosUsuario []string
	// var listaDoTamanho []float64
	// //	var listaContaUsuarios []int
	// contaUsuarios := 0
	// for {
	// 	fmt.Println("Nome do usuario, minimo 15 caracteres:")
	// 	nomeUsuario, err := reader.ReadString('\n')
	// 	pegaNomeUsuario := strings.TrimSpace(nomeUsuario)
	// 	pegaNomeUsuario = strings.ToUpper(pegaNomeUsuario)
	// 	if err != nil {
	// 		fmt.Println("O que você digitou não é um usuario")
	// 		continue
	// 	} else if pegaNomeUsuario == "" {
	// 		break
	// 	} else if len(pegaNomeUsuario) < 15 {
	// 		fmt.Println("Usuario invalido, nome muito pequeno")
	// 		continue
	// 	}

	// 	listaDosUsuario = append(listaDosUsuario, pegaNomeUsuario)

	// 	fmt.Println("Espaço utilizado:")
	// 	espacousado, _ := reader.ReadString('\n')
	// 	pegaEspacoUtilizado := strings.TrimSpace(espacousado)
	// 	espacoutilizado, err := strconv.ParseFloat(pegaEspacoUtilizado, 64)
	// 	if err != nil {
	// 		fmt.Println("O que você digitou não é um número")
	// 		continue
	// 	}
	// 	listaDoTamanho = append(listaDoTamanho, espacoutilizado)
	// 	//	listaContaUsuarios = append(listaContaUsuarios, +1)
	// 	contaUsuarios += 1
	// }

	// somaEspacoOcupado := 0.0
	// for i := 0; i < contaUsuarios; i++ {
	// 	somaEspacoOcupado = somaEspacoOcupado + listaDoTamanho[i]
	// }

	// espacoMedioOcupado := somaEspacoOcupado / float64(contaUsuarios)

	// test := 1
	// fmt.Println("Nr. Usuário     Espaço utilizado     %", " do uso")
	// for i := 0; i < contaUsuarios; i++ {
	// 	fmt.Println(test, listaDosUsuario[i], listaDoTamanho[i], "MB ", fmt.Sprintf("%.2f", (listaDoTamanho[i]*100)/float64(somaEspacoOcupado)), "%")
	// 	test = test + 1
	// }
	// fmt.Println("Espaço total ocupado:", somaEspacoOcupado)
	// fmt.Println("Espaço médio ocupado:", espacoMedioOcupado)
}

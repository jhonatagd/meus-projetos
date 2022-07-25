package main

import "fmt"

// Lista/Array/Fila/Pilha, armazenar varios valores dentro de uma variável
func main() {
	listInteiros := []int{}  //lista/arrays de inteiros
	var listaInteiros2 []int //lista/arrays de inteiros
	var listaString []string //lista/arrays de string
	var listaFloat []float64 //lista/arrays de float
	var listaBool []bool     //lista/arrays de boolean

	listaStringComValor := []string{"valor1", "2", "valor3"} //Criar listas com valores
	listaIntComValor := []int{1, 2, 3}                       //Criar listas com valores

	//listas com limite
	var listaStringComLimite [3]string
	listaStringComLimiteEValor := [3]string{"1", "teste", "123"}

	//Acessar posição de cada item na lista
	listaPreenchida := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(listaPreenchida[0])                      // -> 0
	fmt.Println(listaPreenchida[1])                      // -> 1
	fmt.Println(listaPreenchida[2])                      // -> 2
	fmt.Println(listaPreenchida[3])                      // -> 3
	fmt.Println(listaPreenchida[4])                      // -> 4
	fmt.Println(listaPreenchida[5])                      // -> 5
	fmt.Println(listaPreenchida[6])                      // -> 6
	fmt.Println(listaPreenchida[7])                      // -> 7
	fmt.Println(listaPreenchida[8])                      // -> 8
	fmt.Println(listaPreenchida[9])                      // -> 9
	fmt.Println(listaPreenchida[10])                     // -> Erro
	fmt.Println(len(listaPreenchida))                    // -> 10 que é o tamanho da lista
	fmt.Println(listaPreenchida[len(listaPreenchida)-1]) // -> Pegar o ultimo valor de uma lista

	for i := 0; i < len(listaPreenchida); i++ {
		fmt.Println(listaPreenchida[i]) // 0 1 2 3 4 5 6 7 8 9
	}

	for i := len(listaPreenchida) - 1; i <= 0; i-- {
		fmt.Println(listaPreenchida[i]) // 9 8 7 6 5 4 3 2 1 0
	}

	for i := 0; i < len(listaPreenchida); i++ { //Percorrer a mesma lista ao mesmo tempo
		for j := 0; j < len(listaPreenchida); j++ {
			if listaPreenchida[i] < listaPreenchida[j] {
				// valorMenor = i
				// valorMenor = listaPreenchida[i]
			}
		}
	}

	//Adicionar valor na lista
	listaPreenchida = append(listaPreenchida, 10) // 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	fmt.Println(len(listaPreenchida)) // 11

	var listaComValores []int // vazia
	for i := 0; i < 5; i++ {
		digitarValor := 10 
		listaComValores = append(listaComValores, digitarValor + i)
	}
	// 0 -> 10
	// 1 -> 10 , 11
	// 2 -> 10 , 11, 12
	// 3 -> 10 , 11, 12, 13
	// 4 -> 10 , 11, 12, 13, 14

	//lista[posicão] -> Pegar o valor da lista. Vai retornar o valor daquela posição da lista
	//len(lista) -> Retornar o tamanho da sua lista
	//Sempre que percorrer uma lista em um for, usar o len da lista

	listaPreenchida2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(listaPreenchida2[0])                       // -> 9
	fmt.Println(listaPreenchida2[1])                       // -> 8
	fmt.Println(listaPreenchida2[2])                       // -> 7
	fmt.Println(listaPreenchida2[3])                       // -> 6
	fmt.Println(listaPreenchida2[4])                       // -> 5
	fmt.Println(listaPreenchida2[5])                       // -> 4
	fmt.Println(listaPreenchida2[6])                       // -> 3
	fmt.Println(listaPreenchida2[7])                       // -> 2
	fmt.Println(listaPreenchida2[8])                       // -> 1
	fmt.Println(listaPreenchida2[9])                       // -> 0
	fmt.Println(listaPreenchida2[10])                      // -> Erro
	fmt.Println(len(listaPreenchida2))                     // -> 10 que é o tamanho da lista
	fmt.Println(listaPreenchida2[len(listaPreenchida2)-1]) // -> Pegar o ultimo valor de uma lista

	for i := 0; i < len(listaPreenchida2); i++ {
		fmt.Println(listaPreenchida2[i]) // 9 8 7 6 5 4 3 2 1 0
	}

	for i := len(listaPreenchida2); i <= 0; i-- {
		fmt.Println(listaPreenchida2[i]) // 0 1 2 3 4 5 6 7 8 9
	}
}


// Aluno1 [5],[6],[5],[7],[4]
// Aluno2 [10],[9],[10],[10],[8]
// Aluno3 [2],[2],[1],[4],[3]

var matriz [3][5]int

matriz[1][3] => 10 

matriz[3][3]

[][][]
[][][]
[][][]
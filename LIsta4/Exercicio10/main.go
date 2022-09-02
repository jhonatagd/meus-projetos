package main

import "fmt"

//Faça um Programa que leia dois vetores com 10 elementos cada. Gere um terceiro vetor de 20
//elementos, cujos valores deverão ser compostos pelos elementos intercalados dos dois outros vetores.
func main() {
	//reader := bufio.NewReader(os.Stdin)
	vetorA := [7]string{"A", "A", "A", "A", "A", "A", "A"}
	vetorB := [7]string{"B", "B", "B", "B", "B", "B", "B"}

	/*for i := 0; i < 7; i++ {
		fmt.Println(i+1, "Elemento da lista A:")
		pegaElementoA, err := reader.ReadString('\n')
		pegaElementoStringA := strings.TrimSpace(pegaElementoA)
		pegaElementoStringA = strings.ToUpper(pegaElementoStringA)
		if err != nil {
			fmt.Println("O que você digitou não é um elemento")
		}
		vetorA[i] = pegaElementoStringA
	}
	for j := 0; j < 7; j++ {
		fmt.Println(j+1, "Elemento da lista B:")
		pegaElementoB, err := reader.ReadString('\n')
		pegaElementoStringB := strings.TrimSpace(pegaElementoB)
		pegaElementoStringB = strings.ToUpper(pegaElementoStringB)
		if err != nil {
			fmt.Println("O que você digitou não é um elemento")
		}
		vetorB[j] = pegaElementoStringB
	}
	*/
	// var intercalaElementos [14]string

	//for m := 0; m <= 14; m++ {
	//	if m == 0 || m == 2 || m == 4 || m == 6 || m == 8 || m == 10 || m == 12 {
	//		intercalaElementos[m] = vetorA[m]

	//	} else if m == 1 || m == 3 || m == 5 || m == 7 || m == 9 || m == 11 || m == 13 {
	//		intercalaElementos[m] = vetorB[m]
	//	}
	//}

	//var intercalaElementos [14]string

	//for i := 0; i < 7; i++ {

	//	for m := 1; m <= 14; m++ {
	//	if m == 1 || m == 3 || m == 5 || m == 7 || m == 9 || m == 11 || m == 13 {
	///			intercalaElementos[m] = vetorA[i]

	//	} else if m == 2 || m == 4 || m == 6 || m == 8 || m == 10 || m == 12 || m == 14 {
	//		intercalaElementos[m] = vetorB[i]
	//	}
	//}

	//}
	var intercalaElementos2 []string
	for m := 0; m < 7; m++ {
		intercalaElementos2 = append(intercalaElementos2, vetorA[m])
		intercalaElementos2 = append(intercalaElementos2, vetorB[m])
	}

	//for m := 0; m <= 14; m++ {
	//	for i := 0; i < 12; i++ {
	//		for p := 0; p <= 6; p++ {
	//			if m == i {
	//				intercalaElementos[i] = vetorA[p]
	//			}
	//			i += 1
	//		}
	//	}
	//}
	//	if m == 2 {
	//		intercalaElementos[m] = vetorA[1]
	//	}
	//	if m == 4 {
	//		intercalaElementos[m] = vetorA[2]
	//	}
	//	if m == 6 {
	//		intercalaElementos[m] = vetorA[3]
	//	}
	//	if m == 8 {
	//		intercalaElementos[m] = vetorA[4]
	//	}
	//	if m == 10 {
	//		intercalaElementos[m] = vetorA[5]
	//}
	//if m == 12 {
	//	intercalaElementos[m] = vetorA[6]
	//}

	// outro vetor
	//	for m := 0; m <= 14; m++ {
	//		for i := 1; i < 13; i++ {
	//			for p := 0; p <= 6; p++ {
	//				if m == i {
	//					intercalaElementos[i] = vetorB[p]
	//				}
	//				i += 1
	//			}
	//		}
	//	}
	//	if m == 1 {
	//		intercalaElementos[m] = vetorB[0]
	//	}
	//	if m == 3 {
	//		intercalaElementos[m] = vetorB[1]
	//	}
	//	if m == 5 {
	//		intercalaElementos[m] = vetorB[2]
	//	}
	//	if m == 7 {
	//		intercalaElementos[m] = vetorB[3]
	//	}
	//	if m == 9 {
	//		intercalaElementos[m] = vetorB[4]
	//	}
	//	if m == 11 {
	//		intercalaElementos[m] = vetorB[5]
	//	}
	//	if m == 13 {
	//		intercalaElementos[m] = vetorB[6]
	//	}
	//
	//	}
	fmt.Println(intercalaElementos2)
}

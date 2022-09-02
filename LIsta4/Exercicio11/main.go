package main

import "fmt"

//Altere o programa anterior, intercalando 3 vetores de 10 elementos cada
func main() {
	//reader := bufio.NewReader(os.Stdin)
	vetorA := [7]string{"A", "A", "A", "A", "A", "A", "A"}
	vetorB := [7]string{"B", "B", "B", "B", "B", "B", "B"}
	vetorC := [7]string{"C", "C", "C", "C", "C", "C", "C"}
	var intercalaElementos2 []string

	for m := 0; m < 7; m++ {
		intercalaElementos2 = append(intercalaElementos2, vetorA[m])
		intercalaElementos2 = append(intercalaElementos2, vetorB[m])
		intercalaElementos2 = append(intercalaElementos2, vetorC[m])
	}
	fmt.Println(intercalaElementos2)
}

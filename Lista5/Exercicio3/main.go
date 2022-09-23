package main

import "fmt"

func main() {
	fmt.Println(soma3(1, 1, 1))
	fmt.Println(soma3(2, 2, 2))
	fmt.Println(soma3(3, 3, 3))
}

func soma3(a int, b int, c int) int {
	return a + b + c

}

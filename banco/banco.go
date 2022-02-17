package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

type contaDeTeste struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoMarquito := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDoJp := ContaCorrente{"Jp", 220, 111220, 200}
	contaDoDiegao := ContaCorrente{"Diegao", 221, 111221, 201}
	contaDoAudi := ContaCorrente{"Audi", 222, 111222, 202}
	contaDoGustavo := ContaCorrente{"Gustavo", 223, 111223, 203}
	fmt.Println(fmt.Sprintf("%v %v %v %v %v", contaDoMarquito, contaDoJp, contaDoDiegao, contaDoAudi, contaDoGustavo))

}

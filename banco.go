package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoMarquito := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDoJp := ContaCorrente{"Jp", 222, 111222, 200}
	contaDoDiegao := ContaCorrente{"Diegao", 222, 111222, 200}
	contaDoAudi := ContaCorrente{"Audi", 222, 111222, 200}
	contaDoGustavo := ContaCorrente{"Diegao", 222, 111222, 200}

	fmt.Println(fmt.Sprintf("%v %v %v %v %v", contaDoMarquito, contaDoJp, contaDoDiegao, contaDoAudi, contaDoGustavo))

}

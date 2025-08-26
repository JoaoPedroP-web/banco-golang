package main

import (
	"fmt"

	"github.com/Joao/banco/contas"
)

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}

func main() {
	contaJoao := contas.ContaPoupanca{}
	contaJoao.Depositar(5000)
	pagarBoleto(&contaJoao, 4950)


	contaMaria := contas.ContaCorrente{}
	contaMaria.Depositar(5000)
	pagarBoleto(&contaMaria, 4950)
	

	fmt.Println(contaJoao.ObterSaldo(), contaMaria.ObterSaldo())
}
package contas

import "github.com/Joao/banco/clientes"

type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia int
	NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := c.saldo > valorDoSaque && valorDoSaque > 0
	if podeSacar {
		c.saldo = c.saldo - valorDoSaque
		return "Saque efetuado com sucesso. O valor do seu saldo atual é: R$ ", c.saldo
	} else {
		return "saldo insuficiente. O valor do seu saldo atual é: R$ ", c.saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito efetuado com sucesso. O valor do seu saldo atual é: R$",c.saldo
	} else {
		return "Deposite um valor acima de zero. O valor do seu saldo atual é: R$",c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
# banco-golang
Estudo de OOP em Golang aplicado a operações bancárias.

O projeto tem como objetivo realizar uma simulação de operações bancárias por meio das funções “Sacar”, “Depositar” “ObterSaldo” e “Transferir”. Essas funções estão aplicadas em dois arquivos Go diferentes, um para cada tipo de conta (“contaCorrente.go” e “contaPoupanca.go”), enquanto a execução das operações é feita pelo arquivo “main.go”.

OBJETIVO:
•	Montar de maneira estruturalmente cautelosa operações de banco realistas utilizando a linguagem Golang.
•	Cada função possui elementos condicionais que verificam a possibilidade de operação e execução da função.
•	Dar ao usuário a possibilidade de realizar cálculos bancários pessoais por meio de uma aplicação Golang.

TECNOLOGIA UTILIZADA:
GOLANG (https://golang.org)

ESTRUTURA:
O repositório está organizado em duas pastas separadas: “clientes” e “contas”. A pasta “clientes”, cada pasta guarda arquivos referentes a sua classe, separadamente.

Pasta “clientes”: 
Arquivo “cliente.go” – código responsável por guardar a Struct “Titular” com informações chave específicas do cliente (Nome, CPF e Profissão)

Pasta “contas”: 
Arquivo “contaCorrente.go” – código responsável por conter Struct de uma conta corrente para as informações principais dela (Titular, que é importada do código “cliente.go”, Numero de agência, número da conta e saldo) e as seguintes funções para a conta:
Sacar – realiza, se possível, o saque de um valor na conta.
Depositar – deposita, se possível, um valor na conta.
Transferir – Transfere, se possível, um valor a uma conta destino
ObterSaldo – Retorna o saldo atual da conta.

Arquivo “contaPoupanca.go” – código responsável por conter:
Uma Struct de uma conta corrente para as informações principais dela (Titular, que é importada do código “cliente.go”, Numero de agência, número da conta, operação e saldo) e as seguintes funções para a conta:
Sacar – realiza, se possível, o saque de um valor na conta.
Depositar – deposita, se possível, um valor na conta.
ObterSaldo – Retorna o saldo atual da conta.

main.go – 
Função “PagarBoleto” – uma função que incorpora a função “Sacar”, mas o diferencial é que ela permite ao usuário declarar um parâmetro que representa o valor do boleto para depois realizar as operações originais da função sacar. Basicamente um simulador de pagar boletos, tal qual em um banco de verdade.
Código fonte principal para a execução das funções.

COMO FUNCIONA:
•	O programa permite a criação de informações pessoais de titulares e contas separadas em pastas diferentes.
•	O usuário utiliza uma lógica de OOP para a criação de contas poupanças e/ou contas conrrentes.
•	Depois de serem criadas, o usuário pode realizar operações financeiras de acordo com as funções criadas, que representam operações bancárias.

AUTOR:
João Pedro Pereira Requejo

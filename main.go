// Programa de gerenciamento bancário básico em Go.
//
// Este programa define uma estrutura ContaCorrente para simular operações bancárias simples, como
// saque, depósito e transferência entre contas. As funções implementadas permitem realizar essas 
// operações e também imprimir o status das contas envolvidas. O código demonstra conceitos como 
// métodos, ponteiros e estruturas.
//
// Funcionalidades:
// - Saque: Deduz um valor do saldo da conta, caso haja saldo suficiente.
// - Depósito: Adiciona um valor ao saldo da conta, desde que o valor seja positivo.
// - Transferência: Move um valor entre duas contas correntes, verificando o saldo disponível.


package main

import "fmt"

// ContaCorrente representa uma conta bancária com informações do titular,
// número da agência, número da conta e saldo disponível.
type ContaCorrente struct {
	Titular       string  // Nome do titular da conta
	NumeroAgencia int     // Número da agência
	NumeroConta   int     // Número da conta
	Saldo         float64 // Saldo disponível na conta
}

// Sacar realiza a operação de saque, verificando se o valor é positivo
// e se há saldo suficiente para a transação. Retorna uma string indicando o status do saque.
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

// Depositar adiciona o valor especificado ao saldo da conta se ele for positivo.
// Retorna uma mensagem de sucesso ou erro, e o novo saldo da conta.
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", c.Saldo
	} else {
		return "Valor do depósito precisa ser maior que zero", c.Saldo
	}
}

// Transferir realiza uma transferência de saldo entre duas contas correntes, 
// verificando se o valor é positivo e se o saldo é suficiente para realizar a transferência.
// Retorna um booleano indicando o status da operação.
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

// Função principal do programa, que inicializa contas correntes e realiza operações
// como transferências, comparações entre contas e impressões de status.
func main() {
	contaDoJoao := ContaCorrente{Titular: "João", NumeroAgencia: 111, NumeroConta: 333444, Saldo: 450}
	contaDaBeatriz := ContaCorrente{"Beatriz", 222, 111222, 300}
	contaDaBeatriz2 := ContaCorrente{"Beatriz", 221, 111222, 300}

	fmt.Println(contaDoJoao)
	fmt.Println(contaDaBeatriz == contaDaBeatriz2)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.Titular = "Cris"
	contaDaCris2.Saldo = 500

	contaDaSilvia := ContaCorrente{Titular: "Silvia", Saldo: 305}
	contaDoGustavo := ContaCorrente{Titular: "Gustavo", Saldo: 200}

	// Realiza transferência da conta do Gustavo para a conta da Silvia
	status := contaDoGustavo.Transferir(195, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}

package contas

import (
	"errors"
	"fmt"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
)

type Contas struct {
	ContaId         int
	NumeroDocumento int
	Saldo           float64
	transacoes      []int
}

var ContasRegistradas map[int]Contas
var m chan int

func IsContaExiste(id int) error {
	inicializarRegistros()
	if _, ok := ContasRegistradas[id]; ok {
		return nil
	}
	return errors.New("Conta não existente na plataforma!")
}

func RegistrarConta(numeroDocumento int) (int, error) {
	inicializarRegistros()
	var novaConta Contas
	novaConta.Saldo = 0
	novaConta.NumeroDocumento = numeroDocumento
	gerarContaId()
	novoId := <-m
	if IsContaExiste(novoId) == nil {
		return 0, errors.New("Erro na criação da conta! Conta já existente na plataforma")
	}
	novaConta.ContaId = novoId
	ContasRegistradas[novoId] = novaConta
	return novoId, nil

}
func inicializarRegistros() {
	if ContasRegistradas == nil {
		ContasRegistradas = make(map[int]Contas)
	}
}
func gerarContaId() {
	if m == nil {
		logs.RegistrarLogInformativo("Iniciando estrutura para geração de conta!")
		m = make(chan int, 1)
	}

	m <- len(ContasRegistradas) + 1
}

func VincularTransacao(contaId int, transacaoId int) error {
	if IsContaExiste(contaId) != nil {
		return errors.New("Erro no registro da transação! conta não existe na plataforma")
	}
	conta := ContasRegistradas[contaId]
	conta.transacoes = append(conta.transacoes, transacaoId)
	ContasRegistradas[contaId] = conta
	return nil
}

func ConsultarConta(contaId int) (Contas, error) {
	if IsContaExiste(contaId) != nil {
		var contaVazia Contas
		return contaVazia, errors.New("Erro na consulta da conta! conta não registrada na plataforma")
	}
	return ContasRegistradas[contaId], nil
}

func ModificarSaldo(contaId int, valor float64, operationId int) error {
	if IsContaExiste(contaId) != nil {
		return errors.New("Erro na consulta da conta! conta não registrada na plataforma")
	}
	conta := ContasRegistradas[contaId]
	if operationId >= 1 && operationId <= 3 {
		if valor < 0 {
			valor = valor * -1
		}
		if conta.Saldo < valor {
			return errors.New(fmt.Sprintf("Erro ao abater saldo. Saldo da conta é menor do que o valor da transação. Saldo da conta: %.2f, Valor da transação: %.2f", conta.Saldo, valor))
		}
		conta.Saldo -= valor
	} else {
		conta.Saldo += valor
	}
	ContasRegistradas[contaId] = conta
	return nil
}

func ConvertConta(conta ContaDTO) Contas {
	var contas Contas
	contas.NumeroDocumento = conta.Document_number
	contas.Saldo = conta.Saldo
	contas.ContaId = conta.Account_id
	return contas
}

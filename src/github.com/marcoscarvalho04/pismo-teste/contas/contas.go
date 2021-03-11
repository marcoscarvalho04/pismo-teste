package contas

import (
	"errors"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
)

type Contas struct {
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

package contas

import (
	"errors"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/transacoes"
)

type Contas struct {
	NumeroDocumento int
	Saldo           float64
	transacoes      []transacoes.TransacoesModel
}

var ContasRegistradas map[int]Contas
var m chan int
var contador int

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
	if IsContaExiste(novoId) != nil {
		return 0, errors.New("Erro na criação da conta!")
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
		m = make(chan int, 1)
		contador = 0
	}
	m <- contador + 1
}

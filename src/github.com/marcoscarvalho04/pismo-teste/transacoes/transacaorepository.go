package transacoes

import (
	"errors"

	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"strconv"
	"time"
)

/*
Pacote: transacoes.

Data de criação: 10/03/2021

Criador: Marcos Siqueira

Breve Descrição: Pacote criado para manipular informações de transação dentro da plataforma.
Ele é responsável por criar operações de adição de transação à conta.
*/

type TransacoesModel struct {
	ContaId     int
	OperacaoId  int
	Valor       float64
	Data        time.Time
	TransacaoId int
}

type Transacoes interface {
	ConsultarTransacao(int) (TransacoesModel, error)
	RegistrarTransacao(TransacoesModel) error
}

var m chan int
var contador int

var TransacoesRegistradas map[int]TransacoesModel

func ConsultarTransacao(id int) (TransacoesModel, error) {
	if Transacao, ok := TransacoesRegistradas[id]; ok {
		return Transacao, nil
	}
	error := errors.New("Transação não existente no sistema!")
	var vazio TransacoesModel
	return vazio, error
}
func RegistrarTransacao(novaTransacao TransacoesModel) (int, error) {
	if m == nil {
		m = make(chan int, 1)
		contador = 0
	}
	if TransacoesRegistradas == nil {
		TransacoesRegistradas = make(map[int]TransacoesModel)
	}
	err := checarEAbaterSaldo(novaTransacao)
	if err != nil {
		return 0, err
	}
	gerarIdTransacao()
	idTransacao := <-m
	if _, ok := TransacoesRegistradas[idTransacao]; ok {
		return 0, errors.New("Não foi possível gerar o id da transação. Último ID gerado: " + strconv.Itoa(idTransacao))

	}
	novaTransacao.TransacaoId = idTransacao
	TransacoesRegistradas[idTransacao] = novaTransacao
	contas.VincularTransacao(novaTransacao.ContaId, idTransacao)
	return idTransacao, nil

}

func gerarIdTransacao() {
	if m == nil {
		m = make(chan int, 1)
	}
	m <- len(TransacoesRegistradas) + 1
}

func ConverterTransacao(transacao TransacaoDTO) TransacoesModel {
	var transacaoModel TransacoesModel
	transacaoModel.ContaId = transacao.Account_id
	transacaoModel.Data = time.Now()
	transacaoModel.OperacaoId = transacao.Operation_type_id
	transacaoModel.Valor = transacao.Amount
	transacaoModel.TransacaoId = transacao.TransactionId
	return transacaoModel
}

func checarEAbaterSaldo(transacao TransacoesModel) error {
	err := contas.ModificarSaldo(transacao.ContaId, transacao.Valor, transacao.OperacaoId)
	if err != nil {
		return err
	}

	return nil
}

package transacoes

import (
	"errors"
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
	ContaId    int
	OperacaoId int
	Valor      float64
	Data       time.Time
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
func RegistrarTransacao(novaTransacao TransacoesModel) error {
	if m == nil {
		m = make(chan int, 1)
		contador = 0
	}
	gerarIdTransacao()
	idTransacao := <-m
	if _, ok := TransacoesRegistradas[idTransacao]; ok {
		return errors.New("Não foi possível gerar o id da transação. Último ID gerado: " + strconv.Itoa(idTransacao))

	}
	TransacoesRegistradas[idTransacao] = novaTransacao
	return nil

}

func gerarIdTransacao() {
	if m == nil {
		m = make(chan int, 1)
		contador = 0
	}
	m <- contador + 1
}

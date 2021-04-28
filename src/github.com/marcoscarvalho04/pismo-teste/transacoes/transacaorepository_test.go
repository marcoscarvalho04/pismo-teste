package transacoes

import (
	"fmt"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"testing"
	"time"
)

const ERRO_VALIDAR_CAMPO string = "Erro ao validar o campo: %v, obtido: %v, esperado: %v"

func TestRegistrarTransacaoComSucesso(t *testing.T) {
	conta, errCadastrarConta := contas.RegistrarConta(12345)
	if errCadastrarConta != nil {
		t.Errorf("Erro ao cadastrar conta: %v", errCadastrarConta.Error())
		return
	}

	var transacao TransacoesModel
	transacao.ContaId = conta
	transacao.Data = time.Now()
	transacao.Valor = 10
	transacao.OperacaoId = 4

	transacaoId, err := RegistrarTransacao(transacao)
	if err != nil {
		t.Errorf(fmt.Sprintf("Erro ao registrar transação: %v", err.Error()))
		return
	}
	if transacaoId == 0 {
		t.Errorf("Erro ao registrar transação. Id esperado: %v, obtido: %v", transacaoId, 0)
		return
	}
	transacaoRegistrada, errConsultarTransacao := ConsultarTransacao(transacaoId)
	if errConsultarTransacao != nil {
		t.Errorf(fmt.Sprintf("Erro ao consultar transação: %v", errConsultarTransacao.Error()))
		return
	}
	if transacaoRegistrada.ContaId != transacao.ContaId {
		t.Errorf(fmt.Sprintf(ERRO_VALIDAR_CAMPO, "contaId", transacao.ContaId, transacaoRegistrada.ContaId))
		return
	}
	if transacaoRegistrada.Data != transacao.Data {
		t.Errorf(fmt.Sprintf(ERRO_VALIDAR_CAMPO, "Data", transacao.Data, transacaoRegistrada.Data))
		return
	}
	if transacaoRegistrada.OperacaoId != transacao.OperacaoId {
		t.Errorf(fmt.Sprintf(ERRO_VALIDAR_CAMPO, "OperacaoId", transacao.OperacaoId, transacaoRegistrada.OperacaoId))
		return
	}
	if transacaoRegistrada.Valor != transacao.Valor {
		t.Errorf(fmt.Sprintf(ERRO_VALIDAR_CAMPO, "Valor", transacaoRegistrada.Valor, transacao.Valor))
		return
	}
	t.Logf("TestRegistrarTransacaoComSucesso passou!")

}

func TestRegistrarTransacaoComContaInvalida(t *testing.T) {
	var transacao TransacoesModel
	transacao.ContaId = 30
	transacao.Data = time.Now()
	transacao.OperacaoId = 4
	transacao.Valor = 10
	_, err := RegistrarTransacao(transacao)
	if err == nil {
		t.Errorf(fmt.Sprintf("Erro esperado e não ocorrido"))
		return
	}
	t.Logf("TestRegistrarTransacaoComContaInvalida passou!")
}

func TestConsultarTransacaoInexistente(t *testing.T) {
	_, err := ConsultarTransacao(5)
	if err == nil {
		t.Errorf(fmt.Sprintf("Exceção esperada, mas não obtida"))
		return
	}
	t.Logf("TestConsultarTransacaoInexistente passou!")

}

func TestConverterTransacaoDTO(t *testing.T) {
	var transacaoDTO TransacaoDTO
	transacaoDTO.Account_id = 1
	transacaoDTO.Amount = 10
	transacaoDTO.Operation_type_id = 4
	transacao := ConverterTransacao(transacaoDTO)
	if transacao.ContaId != transacaoDTO.Account_id {
		t.Errorf(fmt.Sprintf(ERRO_VALIDAR_CAMPO, "contaId", transacao.Valor, transacaoDTO.Amount))
		return
	}
	if transacao.Valor != transacaoDTO.Amount {
		t.Errorf(fmt.Sprintf(ERRO_VALIDAR_CAMPO, "valor", transacao.Valor, transacaoDTO.Amount))
		return
	}
	if transacao.OperacaoId != transacaoDTO.Operation_type_id {
		t.Errorf(fmt.Sprintf(ERRO_VALIDAR_CAMPO, "operationId", transacao.OperacaoId, transacaoDTO.Operation_type_id))
		return
	}

	t.Logf("TestConverterTransacaoDTO passou!")

}

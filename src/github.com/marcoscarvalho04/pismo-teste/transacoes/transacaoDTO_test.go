package transacoes

import (
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/testutil"
	"testing"
	"time"
)

const NUMERO_DOCUMENTO int = 12345
const ERRO_VALIDACAO_CAMPO string = "Erro ao validar campo: %v, esperado: %v, obtido: %v"

func TestConvertDTOSucesso(t *testing.T) {
	var transacao TransacoesModel
	transacao.ContaId = 1
	transacao.Data = time.Now()
	transacao.OperacaoId = 1
	transacao.Valor = -10

	transacaoDTO := ConverterDTO(transacao)
	if transacaoDTO.Account_id != transacao.ContaId {
		t.Errorf(ERRO_VALIDACAO_CAMPO, "contaId", transacao.ContaId, transacaoDTO.Account_id)
		return
	}
	if transacaoDTO.Amount != transacao.Valor {
		t.Errorf(ERRO_VALIDACAO_CAMPO, "valor", transacaoDTO.Amount, transacao.Valor)
		return
	}
	if transacaoDTO.Operation_type_id != transacao.OperacaoId {
		t.Errorf(ERRO_VALIDACAO_CAMPO, "operationId", transacaoDTO.Operation_type_id, transacao.OperacaoId)
		return
	}
	if transacaoDTO.TransactionId != transacao.TransacaoId {
		t.Errorf(ERRO_VALIDACAO_CAMPO, "transacationID", transacaoDTO.TransactionId, transacao.TransacaoId)
		return
	}
	t.Logf("TestParseTransactionComSucesso passou!")
}

func TestParseTransactionSucesso(t *testing.T) {
	jsonEntrada := testutil.FormatarNovaEntradaTransacao(1, 2, -10)
	transacao, err := ParseTransaction([]byte(jsonEntrada))
	if err != nil {
		t.Errorf("Erro ao gerar transação formatada: %v", err.Error())
		return
	}
	if transacao.Account_id != 1 {
		t.Errorf(ERRO_VALIDACAO_CAMPO, "ContaId", transacao.Account_id, 1)
		return
	}
	if transacao.Amount != -10 {
		t.Errorf(ERRO_VALIDACAO_CAMPO, "valor", transacao.Amount, -10)
		return
	}
	if transacao.Operation_type_id != 2 {
		t.Errorf(ERRO_VALIDACAO_CAMPO, "operationId", transacao.Operation_type_id, 2)
		return
	}
	t.Logf("TestParseTransactionSucesso passou!")
}

func TestParseJsonInvalido(t *testing.T) {
	_, err := ParseTransaction([]byte(""))
	if err == nil {
		t.Errorf("Esperado erro, exceção não ocorrida")
		return
	}
	t.Logf("TestParseJsonInvalido passou!")

}

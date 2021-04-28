package transacoes

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

const MENSAGEM_ERRO_INVALIDA string = "mensagem de erro invalida, esperado: %v, obtido: %v"

func gerarMassaParaTransacao(contaId int, valor float64, operationId int) (TransacaoDTO, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	var transacao TransacaoDTO
	transacao.Account_id = contaId
	transacao.Amount = valor
	transacao.Operation_type_id = operationId
	return transacao, w
}
func validarMensagemEErro(err error, mensagem string, t *testing.T, mensagemSucesso string) {
	if err == nil {
		t.Errorf(fmt.Sprintf("Erro ao validar transação. Esperado erro, não obtido!"))
		return
	}
	if err.Error() != mensagem {
		t.Errorf(MENSAGEM_ERRO_INVALIDA, mensagem, err.Error())
		return
	}
	t.Logf(mensagemSucesso)
	return

}
func TestValidarValorZerado(t *testing.T) {
	transacao, w := gerarMassaParaTransacao(1, 0, 4)
	err := ValidarTransacaoRecebida(transacao, w)
	validarMensagemEErro(err, VALOR_ZERADO_NAO_PERMITIDO, t, "TestVlaidarValorZerado passou!")
}

func TestValidarContaInexistente(t *testing.T) {
	transacao, w := gerarMassaParaTransacao(1000, 10, 4)
	err := ValidarTransacaoRecebida(transacao, w)
	validarMensagemEErro(err, CONTA_NAO_EXISTE, t, "TestvalidarContaInexistente passou!")
}
func TestValidarOperacaoNaoSuportada(t *testing.T) {
	transacao, w := gerarMassaParaTransacao(1, 10, -1)
	err := ValidarTransacaoRecebida(transacao, w)
	validarMensagemEErro(err, OPERACAO_NAO_SUPORTADA, t, "TestValidarOperacaoNaoSuportada passou!")
}
func TestValidarOperacaoSaqueCompraInvalido(t *testing.T) {
	transacao, w := gerarMassaParaTransacao(1, 10, 3)
	err := ValidarTransacaoRecebida(transacao, w)
	validarMensagemEErro(err, TRANSACAO_SAQUE_COMPRA_VALOR_NAO_PERMITIDO, t, "TestValidarOperacaoSaqueCompraInvalido")
}
func TestValidarPagamentoNaoPermitido(t *testing.T) {
	transacao, w := gerarMassaParaTransacao(1, -200, 4)
	err := ValidarTransacaoRecebida(transacao, w)
	validarMensagemEErro(err, PAGAMENTO_VALOR_NAO_PERMITIDO, t, "TestValidarPagamentoNaoPermitido passou!")
}

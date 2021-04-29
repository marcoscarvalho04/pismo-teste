package requisicoes

import (
	"encoding/json"
	"errors"
	"fmt"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/services"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/testutil"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/transacoes"
	"strconv"
	"strings"
	"testing"
	"time"
)

const JSON_ENTRADA_CONTAS string = "{\"document_number\": 12345}"

const ERRO_STATUS_CODE string = "Status code inválido. Esperado: %v, obtido: %v"
const DOCUMENT_NUMBER int = 12345

func TestResponderCriarContaSucesso(t *testing.T) {
	statusCode, response := testutil.FazerRequisicaoParaURL("POST", "/accounts", JSON_ENTRADA_CONTAS, ResponderCriarConta, nil)
	if statusCode != 201 {
		t.Errorf(ERRO_STATUS_CODE, 201, statusCode)
		return
	}
	contaId := response[23]
	mensagemResposta := "Conta criada com o ID: " + string(contaId)
	if response != mensagemResposta {
		t.Errorf("Erro ao criar conta, esperado id: %v, obtido: %v", 1, contaId)
		return
	}
	t.Logf("Método: TestResponderCriarContaSucesso passou!")
}

func TestResponderCriarContaSemBody(t *testing.T) {
	statusCode, response := testutil.FazerRequisicaoParaURL("POST", "/accounts", "", ResponderCriarConta, nil)
	if statusCode != 400 {
		t.Errorf(ERRO_STATUS_CODE, 400, statusCode)
		return
	}
	if response != "Impossível ler o JSON de cadastro da conta" {
		t.Errorf("Erro ao ler o body da requisição. Esperado: Impossível ler o JSON de cadastro da conta, obtido: %v", response)
		return
	}
	t.Logf("TesteResponderCriarContaSemBody passou!")
}

func TestConsultarContaComSucesso(t *testing.T) {
	statusCodeCriacaoConta, responseCriarConta := testutil.FazerRequisicaoParaURL("POST", "/accounts", JSON_ENTRADA_CONTAS, ResponderCriarConta, nil)
	if statusCodeCriacaoConta != 201 {
		t.Errorf(ERRO_STATUS_CODE, 201, statusCodeCriacaoConta)
		return
	}
	contaId, errConverterConta := strconv.Atoi(string(responseCriarConta[23:]))
	if errConverterConta != nil {
		t.Errorf("Erro ao criar conta: %v", errConverterConta.Error())
		return
	}
	params := make(map[string]string)
	params["id"] = strconv.Itoa(contaId)
	statusCode, response := testutil.FazerRequisicaoParaURL("GET", "/accounts/"+strconv.Itoa(contaId), "", ResponderConsultarConta, params)
	if statusCode != 200 {
		t.Errorf(ERRO_STATUS_CODE, 200, statusCode)
		t.Logf(response)
		return
	}
	contaRecebida, errParse := contas.ParseDTO(response)
	if errParse != nil {
		t.Errorf("Erro ao fazer parse da conta recebida: %v", errParse.Error())
		return
	}
	if contaRecebida.Account_id != contaId {
		t.Errorf("Erro ao fazer consulta da conta. Id esperado: %v, obtido: %v", contaId, contaRecebida.Account_id)
		return
	}
	if contaRecebida.Document_number != DOCUMENT_NUMBER {
		t.Errorf("Erro ao fazer consulta da conta. Document number esperado: %v, obtido: %v", contaRecebida.Document_number, DOCUMENT_NUMBER)
		return
	}
	if contaRecebida.Saldo > 0 {
		t.Errorf("Erro ao fazer consulta da conta. Saldo esperado deve ser 0, obtido: %v", contaRecebida.Saldo)
		return
	}
	t.Logf("TestConsultarContaComSucesso passou!")

}

func TestConsultarContaSemParametro(t *testing.T) {
	statusCode, _ := testutil.FazerRequisicaoParaURL("GET", "/accounts", "", ResponderConsultarConta, nil)
	statusCodeEsperado := 400
	if statusCode != statusCodeEsperado {
		t.Errorf(ERRO_STATUS_CODE, statusCodeEsperado, statusCode)
		return
	}
}

func TestCriarTransacaoComSucesso(t *testing.T) {
	novaConta, err := criarContaParaTeste(t)
	if err != nil {
		t.Errorf("Erro ao criar conta nova para teste: %v", err.Error())
		return
	}
	jsonEntradaTransacao := testutil.FormatarNovaEntradaTransacao(novaConta.ContaId, 4, 100)
	statusCode, response := testutil.FazerRequisicaoParaURL("POST", "/transactions", jsonEntradaTransacao, ResponderCriarTransacao, nil)
	if statusCode != 201 {
		t.Errorf(ERRO_STATUS_CODE, 201, statusCode)
		return
	}
	novaConta, errConsultarConta := contas.ConsultarConta(novaConta.ContaId)
	if errConsultarConta != nil {
		t.Errorf(errConsultarConta.Error())
		return
	}
	if novaConta.Saldo != 100 {
		t.Errorf("Erro ao validar saldo da conta. Esperado: %v, obtido: %v", 100, novaConta.Saldo)
		return
	}

	t.Logf("Transação registrada para a conta: %v ", string(response[37]))

}

func TestCriarTransacaoSemSaldo(t *testing.T) {
	novaConta, err := criarContaParaTeste(t)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	jsonEntradaTransacao := testutil.FormatarNovaEntradaTransacao(novaConta.ContaId, 3, -100)
	statusCode, response := testutil.FazerRequisicaoParaURL("POST", "/transactions", jsonEntradaTransacao, ResponderCriarTransacao, nil)
	if statusCode != 400 {
		t.Errorf(ERRO_STATUS_CODE, 400, statusCode)
		return
	}
	if !strings.Contains(response, "Erro ao abater saldo. Saldo da conta é menor do que o valor da transação.") {
		t.Errorf("Mensagem de retorno não esperada: %v", response)
		return
	}

}
func TestCriarTransacaoInvalidaSaque(t *testing.T) {
	err := criarERegistrarTransacao(t, 3, 100, transacoes.TRANSACAO_SAQUE_COMPRA_VALOR_NAO_PERMITIDO)
	if err == nil {
		t.Logf("TestCriarTransacaoInvalidaSaque passou!")
	}
}

func TestCriarTransacaoInvalidaCompraParcelada(t *testing.T) {
	err := criarERegistrarTransacao(t, 2, 100, transacoes.TRANSACAO_SAQUE_COMPRA_VALOR_NAO_PERMITIDO)
	if err == nil {
		t.Logf("TestCriarTransacaoInvalidaCompraParcelada passou!")
	}
}
func TestCriarTransacaoInvalidaCompra(t *testing.T) {
	err := criarERegistrarTransacao(t, 1, 100, transacoes.TRANSACAO_SAQUE_COMPRA_VALOR_NAO_PERMITIDO)
	if err == nil {
		t.Logf("TestCriarTransacaoInvalidaCompra passou!")
	}
}
func TestCriarPagamentoInvalido(t *testing.T) {
	err := criarERegistrarTransacao(t, 4, -100, transacoes.PAGAMENTO_VALOR_NAO_PERMITIDO)
	if err == nil {
		t.Logf("TestCriarPagamentoInvalido passou!")
	}
}
func TestCriarTransacaoValorZerado(t *testing.T) {
	err := criarERegistrarTransacao(t, 3, 0, transacoes.VALOR_ZERADO_NAO_PERMITIDO)
	if err == nil {
		t.Logf("TestCriarTransacaoValorZerado passou!")
	}
}

func TestTransacaoPagamentoSaque(t *testing.T) {
	novaConta, err := criarContaParaTeste(t)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	var transacaoPagamento transacoes.TransacoesModel
	transacaoPagamento.ContaId = novaConta.ContaId
	transacaoPagamento.Data = time.Now()
	transacaoPagamento.Valor = 200
	transacaoPagamento.OperacaoId = 4
	_, errTransacaoPagamento := transacoes.RegistrarTransacao(transacaoPagamento)
	if errTransacaoPagamento != nil {
		t.Errorf(fmt.Sprintf("Erro ao registrar transação de pagamento: %v", errTransacaoPagamento.Error()))
		return
	}
	novaConta, errConsultarConta := contas.ConsultarConta(novaConta.ContaId)
	if errConsultarConta != nil {
		t.Errorf(fmt.Sprintf("Erro ao consultar conta com novo saldo: %v", errConsultarConta.Error()))
		return
	}
	if novaConta.Saldo != 200 {
		t.Errorf("Saldo da conta diferente do esperado. Esperado: %v, obtido: %v", 200, novaConta.Saldo)
		return
	}
	var transacaoSaque transacoes.TransacoesModel
	transacaoSaque.ContaId = novaConta.ContaId
	transacaoSaque.Data = time.Now()
	transacaoSaque.OperacaoId = 3
	transacaoSaque.Valor = -200

	_, errTransacaoSaque := transacoes.RegistrarTransacao(transacaoSaque)
	if errTransacaoSaque != nil {
		t.Errorf("Erro ao registrar saque: %v", errTransacaoSaque.Error())
		return
	}
	novaConta, errConsultarConta = contas.ConsultarConta(novaConta.ContaId)
	if errConsultarConta != nil {
		t.Errorf("Erro ao consultar conta: %v", errConsultarConta.Error())
		return
	}
	if novaConta.Saldo != 0 {
		t.Errorf("Saldo da conta diferente do esperado. Esperado: %v, obtido: %v", 0, novaConta.Saldo)
		return
	}
	t.Logf("TestTransacaoPagamentoSaque passou!")
}

func TestCriarTransacaoParaContaInexistente(t *testing.T) {
	var novaTransacao transacoes.TransacoesModel
	novaTransacao.ContaId = 500
	novaTransacao.OperacaoId = 1
	novaTransacao.Valor = -100
	novaTransacao.Data = time.Now()

	transacaoId, err := transacoes.RegistrarTransacao(novaTransacao)
	if err == nil {
		t.Errorf("Esperado erro, mas não houve exceção.")
		return
	}
	if transacaoId > 0 {
		t.Errorf("TransacaoId invalido! Esperado: %v, obtido %v", 0, transacaoId)
		return
	}
	t.Logf("TestCriarTransacaoParaContaInexistente passou!")

}

func criarERegistrarTransacao(t *testing.T, operationId int, valor float64, mensagemRetorno string) error {
	novaConta, err := criarContaParaTeste(t)
	if err != nil {
		t.Errorf(err.Error())
		return err
	}
	jsonEntradaTransacao := testutil.FormatarNovaEntradaTransacao(novaConta.ContaId, operationId, valor)
	statusCode, response := testutil.FazerRequisicaoParaURL("POST", "/transactions", jsonEntradaTransacao, ResponderCriarTransacao, nil)
	if statusCode != 400 {
		t.Errorf(ERRO_STATUS_CODE, 400, statusCode)
		return errors.New(fmt.Sprintf(ERRO_STATUS_CODE, 400, statusCode))
	}
	if response != mensagemRetorno {
		t.Errorf("Mensagem de retorno não esperada.\n Esperada: %v, \n Obtida %v ", mensagemRetorno, response)
		return errors.New(fmt.Sprintf("Mensagem de retorno não esperada.\n Esperada: %v, \n Obtida %v ", mensagemRetorno, response))
	}
	return nil

}

func criarContaParaTeste(t *testing.T) (contas.Contas, error) {
	var conta contas.Contas
	statusCodeCriacaoConta, responseCriarConta := testutil.FazerRequisicaoParaURL("POST", "/accounts", JSON_ENTRADA_CONTAS, ResponderCriarConta, nil)
	if statusCodeCriacaoConta != 201 {
		t.Errorf(ERRO_STATUS_CODE, 201, statusCodeCriacaoConta)
		return conta, errors.New(fmt.Sprint(ERRO_STATUS_CODE, 201, statusCodeCriacaoConta))
	}
	contaId, errConverterConta := strconv.Atoi(string(responseCriarConta[23:]))
	if errConverterConta != nil {
		t.Errorf("Erro ao criar conta: %v", errConverterConta.Error())
		return conta, errors.New(fmt.Sprintf("Erro ao criar conta: %v", errConverterConta.Error()))
	}
	params := make(map[string]string)
	params["id"] = strconv.Itoa(contaId)
	statusCode, response := testutil.FazerRequisicaoParaURL("GET", "/accounts/"+strconv.Itoa(contaId), "", ResponderConsultarConta, params)
	if statusCode != 200 {
		t.Errorf(ERRO_STATUS_CODE, 200, statusCode)
		return conta, errors.New(fmt.Sprintf(ERRO_STATUS_CODE, 200, statusCode))
	}
	contaRecebida, errParse := contas.ParseDTO(response)
	if errParse != nil {
		t.Errorf("Erro ao fazer parse da conta recebida: %v", errParse.Error())
		return conta, errors.New(fmt.Sprintf("Erro ao fazer parse da conta recebida: %v", errParse.Error()))
	}
	conta = contas.ConvertConta(contaRecebida)
	return conta, nil
}

func TestConsultarTransacoesContaSucesso(t *testing.T) {
	conta, errCriarConta := contas.RegistrarConta(DOCUMENT_NUMBER)
	if errCriarConta != nil {
		t.Errorf("Erro ao criar conta: %v", errCriarConta.Error())
		return
	}
	var registrarNovaTransacao transacoes.TransacoesModel
	registrarNovaTransacao.ContaId = conta
	registrarNovaTransacao.Data = time.Now()
	registrarNovaTransacao.OperacaoId = 4
	registrarNovaTransacao.Valor = 100
	transacaoId, errRegistrarTransacao := transacoes.RegistrarTransacao(registrarNovaTransacao)
	if errRegistrarTransacao != nil {
		t.Errorf("Erro ao registrar transação: %v", errRegistrarTransacao.Error())
		return
	}
	params := make(map[string]string)
	params["id"] = strconv.Itoa(conta)
	statusCode, response := testutil.FazerRequisicaoParaURL("GET", "/transactions/accounts/{id}", "", ResponderConsultarTransacaoPorConta, params)
	if statusCode != 200 {
		t.Errorf(ERRO_STATUS_CODE, 200, statusCode)
		return
	}
	var transacaoRetornada services.ContaTransacaoDTO
	err := json.Unmarshal([]byte(response), &transacaoRetornada)
	if err != nil {
		t.Errorf("Erro ao deserializar o json: %v", err.Error())
		return
	}
	if transacaoRetornada.ContaId != conta {
		t.Errorf("Erro ao consultar conta. Esperado: %v, obtido: %v", conta, transacaoRetornada.ContaId)
		return
	}
	if len(transacaoRetornada.Transacoes) != 1 {
		t.Errorf("Erro ao verificar transações vinculadas. Esperado: %v, obtido: %v", 1, len(transacaoRetornada.Transacoes))
		return
	}
	if transacaoRetornada.Transacoes[0].TransactionId != transacaoId {
		t.Errorf("Erro ao verificar transação vinculada. Id de transação esperado: %v, obtido: %v", transacaoId, transacaoRetornada.Transacoes[0].TransactionId)
		return
	}
	t.Logf("TestConsultarTransacoesContaSucesso")

}

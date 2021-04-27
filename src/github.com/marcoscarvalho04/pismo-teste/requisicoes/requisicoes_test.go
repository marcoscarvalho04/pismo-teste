package requisicoes

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

const JSON_ENTRADA_CONTAS string = "{\"document_number\": 12345}"
const JSON_ENTRADA_TRANSACAO string = "{\"account_id\": %d,\"operation_type_id\": %d,\"amount\": %.2f}"

const ERRO_STATUS_CODE string = "Status code inválido. Esperado: %v, obtido: %v"
const DOCUMENT_NUMBER int = 12345

type handler func(http.ResponseWriter, *http.Request)

func fazerRequisicaoParaURL(method string, url string, body string, handler handler, params map[string]string) (statusCode int, responseString string) {
	var req *http.Request

	if len(body) == 0 {
		req = httptest.NewRequest(method, url, nil)
	} else {
		req = httptest.NewRequest(method, url, strings.NewReader(body))
	}
	if params != nil {
		req = mux.SetURLVars(req, params)
	}

	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	response, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		statusCode = 0
		responseString = "Erro"
		return statusCode, responseString
	}
	responseString = string(response)
	statusCode = resp.StatusCode
	return statusCode, responseString
}
func TestResponderCriarContaSucesso(t *testing.T) {
	statusCode, response := fazerRequisicaoParaURL("POST", "/accounts", JSON_ENTRADA_CONTAS, ResponderCriarConta, nil)
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
	statusCode, response := fazerRequisicaoParaURL("POST", "/accounts", "", ResponderCriarConta, nil)
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
	statusCodeCriacaoConta, responseCriarConta := fazerRequisicaoParaURL("POST", "/accounts", JSON_ENTRADA_CONTAS, ResponderCriarConta, nil)
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
	statusCode, response := fazerRequisicaoParaURL("GET", "/accounts/"+strconv.Itoa(contaId), "", ResponderConsultarConta, params)
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
	if contaRecebida.ContaId != contaId {
		t.Errorf("Erro ao fazer consulta da conta. Id esperado: %v, obtido: %v", contaId, contaRecebida.ContaId)
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
	statusCode, _ := fazerRequisicaoParaURL("GET", "/accounts", "", ResponderConsultarConta, nil)
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
	jsonEntradaTransacao := formatarNovaEntradaTransacao(novaConta.ContaId, 4, 100)
	statusCode, response := fazerRequisicaoParaURL("POST", "/transactions", jsonEntradaTransacao, ResponderCriarTransacao, nil)
	if statusCode != 200 {
		t.Errorf(ERRO_STATUS_CODE, 200, statusCode)
		return
	}
	t.Logf("resultado: " + response)

}

func criarContaParaTeste(t *testing.T) (contas.Contas, error) {
	var conta contas.Contas
	statusCodeCriacaoConta, responseCriarConta := fazerRequisicaoParaURL("POST", "/accounts", JSON_ENTRADA_CONTAS, ResponderCriarConta, nil)
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
	statusCode, response := fazerRequisicaoParaURL("GET", "/accounts/"+strconv.Itoa(contaId), "", ResponderConsultarConta, params)
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

func formatarNovaEntradaTransacao(contaId int, operationId int, valor float64) string {
	return fmt.Sprintf(JSON_ENTRADA_TRANSACAO, contaId, operationId, valor)
}

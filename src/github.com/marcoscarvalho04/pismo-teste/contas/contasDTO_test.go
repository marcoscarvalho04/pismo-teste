package contas

import (
	"strconv"
	"testing"
)

const ERRO_CONVERT_DTO string = "TestSucessoConvertDTO falhou no campo %v,  esperado: %v \n resultado: %v "
const SUCESSO_CONVERT_DTO string = "TestSucessoConvertDTO passou"

func TestSucessoConvertDTO(t *testing.T) {
	var conta Contas
	conta.ContaId = 12
	conta.NumeroDocumento = 3265
	conta.Saldo = 0
	contaDTO := ConvertDTO(conta)
	if contaDTO.Account_id != conta.ContaId {
		t.Errorf(ERRO_CONVERT_DTO, "contaId", strconv.Itoa(conta.ContaId), strconv.Itoa(contaDTO.Account_id))
	}
	if contaDTO.Document_number != conta.NumeroDocumento {
		t.Errorf(ERRO_CONVERT_DTO, "Document number", strconv.Itoa(conta.NumeroDocumento), strconv.Itoa(contaDTO.Document_number))
	}
	if contaDTO.Saldo != conta.Saldo {
		t.Errorf(ERRO_CONVERT_DTO, "Saldo", strconv.Itoa(int(contaDTO.Saldo)), strconv.Itoa(int(conta.Saldo)))
	}
	t.Logf(SUCESSO_CONVERT_DTO)
}

func TestSuccessParseDTO(t *testing.T) {
	json := "{\"document_number\": 1234213123 }"
	result, err := ParseDTO(json)
	if err != nil {
		t.Errorf("Erro ao fazer parse do JSON: " + err.Error())
	}
	if result.Document_number != 1234213123 {
		t.Errorf("Erro ao validar número do documento. Esperado %v, obtido: %v", 1234213123, result.Document_number)
	}
	if result.Account_id != 0 {
		t.Errorf("ContaId invalido, esperado: 0, resultado: " + strconv.Itoa(result.Account_id))
	}
	if result.Saldo != 0 {
		t.Errorf("Saldo invalidok, esperado: 0, resultado: %v", result.Saldo)
	}
	t.Logf("Passou na validação. Esperado: %v, obtido: %v", 1234213123, result.Document_number)
}

func TestMalformedJson(t *testing.T) {
	json := "{document_number\": 1234}"
	_, err := ParseDTO(json)
	if err == nil {
		t.Errorf("Erro. Esperando erro, obtido não erro")
	}
	t.Logf("TestMalformedJson passou com sucesso")
}

package contas

import (
	"testing"
)

var numeroDocumento int = 12345

const ERRO_PARSE_CONTA string = "Erro ao fazer o parse de contaDTO no campo %v, contaId eperado: %v, obtido: %v"

func TestContaNaoExistente(t *testing.T) {

	contaId := 1
	err := IsContaExiste(contaId)
	if err == nil {
		t.Errorf("Erro, esperado que a variavel err estivesse preenchida, veio vazia.")
		return
	}
	conta, errNovaConta := RegistrarConta(numeroDocumento)
	if errNovaConta != nil {
		t.Errorf("Erro ao registrar conta de número do documento: %v", numeroDocumento)
		return
	}
	if conta == 0 {
		t.Errorf("ContaId esperado: %v, resultado: %v", 0, conta)
	}
	err = IsContaExiste(conta)
	if err != nil {
		t.Errorf("Erro, esperado que a variável err fosse vazia, veio com erro: %v", err.Error())
		return
	}
	t.Logf("TestContaNaoExistente passou")

}

func TestContaExistente(t *testing.T) {
	conta, err := RegistrarConta(numeroDocumento)
	if err != nil {
		t.Errorf("Erro ao registrar conta de número do documento: %v", numeroDocumento)
		return
	}
	if conta == 0 {
		t.Errorf("Erro ao registrar conta, id deveria ser 2, mas retornou %v", conta)
		return
	}
	t.Logf("TestContaExistente passou, id da conta gerada: %v", conta)
}

func TestConsultarConta(t *testing.T) {
	conta, err := RegistrarConta(numeroDocumento)
	if err != nil {
		t.Errorf("Erro ao registrar conta: %v ", err.Error())
		return
	}
	contaRegistrada, errConsultaConta := ConsultarConta(conta)
	if errConsultaConta != nil {
		t.Errorf("Erro ao consultar conta: %v", errConsultaConta.Error())
		return
	}
	if contaRegistrada.ContaId != conta {
		t.Errorf("Erro pesquisando informações da conta. Id esperado: %v, obtido: %v", conta, contaRegistrada.ContaId)
		return
	}
	if contaRegistrada.NumeroDocumento != numeroDocumento {
		t.Errorf("Erro pesquisando informações da conta. Número documento esperado: %v, obtido: %v", contaRegistrada.NumeroDocumento, numeroDocumento)
		return
	}
	t.Logf("Conta Registrada com sucesso! contaId: %v", contaRegistrada.ContaId)

}
func TestConvertDTOConta(t *testing.T) {
	conta, err := RegistrarConta(numeroDocumento)
	if err != nil {
		t.Errorf("Erro ao registrar conta: %v", err.Error())
		return
	}
	contaRegistrada, errConsultaConta := ConsultarConta(conta)
	if errConsultaConta != nil {
		t.Errorf("Erro ao consultar conta: %v", errConsultaConta.Error())
		return
	}
	contaDTO := ConvertDTO(contaRegistrada)
	if contaDTO.Account_id != contaRegistrada.ContaId {
		t.Errorf(ERRO_PARSE_CONTA, "contaId", contaRegistrada.ContaId, contaDTO.Account_id)
		return
	}
	if contaDTO.Document_number != contaRegistrada.NumeroDocumento {
		t.Errorf(ERRO_PARSE_CONTA, "document number", contaRegistrada.NumeroDocumento, contaDTO.Document_number)
		return
	}
	if contaDTO.Saldo != contaRegistrada.Saldo {
		t.Errorf(ERRO_PARSE_CONTA, "Saldo", contaDTO.Saldo, contaRegistrada.Saldo)
		return
	}
	t.Logf("ConvertDTO em contas passou! ")
}

func TestContaModificarSaldoSucesso(t *testing.T) {
	conta, err := RegistrarConta(numeroDocumento)
	if err != nil {
		t.Errorf("Erro ao registrar conta: %v", err.Error())
		return
	}
	contaRegistrada, errConsultarConta := ConsultarConta(conta)
	if errConsultarConta != nil {
		t.Errorf("Erro ao consultar conta: %v", errConsultarConta.Error())
		return
	}
	ModificarSaldo(contaRegistrada.ContaId, 100, 4)
	contaRegistrada, errConsultarConta = ConsultarConta(conta)
	if errConsultarConta != nil {
		t.Errorf("Erro ao consultar conta: %v", errConsultarConta.Error())
		return
	}
	if contaRegistrada.Saldo != 100 {
		t.Errorf("Erro ao validar saldo da conta. Esperado: %v, Obtido: %v", 100, contaRegistrada.Saldo)
		return
	}
	t.Logf("TestContaModificarSaldoSucesso passou!")
}

func TestModificarContaSemSaldo(t *testing.T) {
	conta, err := RegistrarConta(numeroDocumento)
	if err != nil {
		t.Errorf("Erro ao registrar conta: %v", err.Error())
	}
	errModificarSaldo := ModificarSaldo(conta, 100, 3)
	if errModificarSaldo == nil {
		t.Errorf("Erro. Esperado erro e não encontrado")
		return
	}
	t.Logf("TestModificarContaSemSaldo passou!")
}

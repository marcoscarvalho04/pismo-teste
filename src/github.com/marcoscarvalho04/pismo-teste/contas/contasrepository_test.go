package contas

import "testing"

var numeroDocumento int = 12345

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
	}
	t.Logf("TestContaNaoExistente passou")

}

func TestContaExistente(t *testing.T) {
	conta, err := RegistrarConta(numeroDocumento)
	if err != nil {
		t.Errorf("Erro ao registrar conta de número do documento: %v", numeroDocumento)
	}
	if conta == 0 {
		t.Errorf("Erro ao registrar conta, id deveria ser 2, mas retornou %v", conta)
	}
	t.Logf("TestContaExistente passou, id da conta gerada: %v", conta)
}

func TesteConsultarConta(t *testing.T) {

}

package transacoes

import (
	"errors"
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/requisicoesutil"
)

type Validar func(TransacaoDTO) error

// Constantes de retorno
const CONTA_NAO_EXISTE string = "Conta não existe no sistema!"
const OPERACAO_NAO_SUPORTADA string = "Transação não registrada. Operação não suportada pelo sistema"
const TRANSACAO_SAQUE_COMPRA_VALOR_NAO_PERMITIDO string = "Transação não registrada. Operação de compra e  saque devem ser obrigatoriamente com valor negativo "
const PAGAMENTO_VALOR_NAO_PERMITIDO string = "valor negativo não é permitido para as transações do tipo pagamento!"

// Fim contantes de retorno

func ValidarTransacaoRecebida(transacao TransacaoDTO, response http.ResponseWriter) error {
	passos := configurarPassosValidacao(transacao, response)
	for _, value := range passos {
		err := value(transacao)
		if err != nil {
			return err
		}

	}
	return nil
}

func configurarPassosValidacao(transacao TransacaoDTO, response http.ResponseWriter) []Validar {
	var passos []Validar
	passos = append(passos, validarConta(transacao, response))
	passos = append(passos, validarOperacao(transacao, response))
	passos = append(passos, validarOperacaoValor(transacao, response))
	passos = append(passos, validarValorPagamento(transacao, response))
	return passos
}

func validarConta(transacao TransacaoDTO, response http.ResponseWriter) Validar {
	return func(transacao TransacaoDTO) error {
		err := contas.IsContaExiste(transacao.Account_id)
		if err != nil {
			requisicoesutil.RetornarComRegistroInexistente(CONTA_NAO_EXISTE, response)
			return err
		}
		return nil
	}
}

func validarOperacao(transacao TransacaoDTO, response http.ResponseWriter) Validar {
	return func(transacao TransacaoDTO) error {
		if transacao.Operation_type_id > 4 && transacao.Operation_type_id < 1 {
			requisicoesutil.RetornarComBadRequest(OPERACAO_NAO_SUPORTADA, response)
			return errors.New(OPERACAO_NAO_SUPORTADA)
		}
		return nil
	}
}

func validarOperacaoValor(transacao TransacaoDTO, response http.ResponseWriter) Validar {
	return func(transacao TransacaoDTO) error {
		if (transacao.Operation_type_id == 1 || transacao.Operation_type_id == 2 || transacao.Operation_type_id == 3) && (transacao.Amount > 0) {
			requisicoesutil.RetornarComBadRequest(TRANSACAO_SAQUE_COMPRA_VALOR_NAO_PERMITIDO, response)
			return errors.New(TRANSACAO_SAQUE_COMPRA_VALOR_NAO_PERMITIDO)
		}
		return nil
	}
}

func validarValorPagamento(transacao TransacaoDTO, response http.ResponseWriter) Validar {
	return func(transacao TransacaoDTO) error {
		if transacao.Operation_type_id == 4 && transacao.Amount < 0 {
			return errors.New(PAGAMENTO_VALOR_NAO_PERMITIDO)
		}
		return nil
	}
}

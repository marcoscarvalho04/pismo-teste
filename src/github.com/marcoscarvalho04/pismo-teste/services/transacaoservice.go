package services

import (
	"fmt"
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/requisicoesutil"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/transacoes"
	"strings"
)

const TRANSACAO_CRIADA string = "Transação registrada para a conta: %v"
const ERRO_GERAR_TRANSACAO string = "Erro ao gerar transacao para a conta: %v "

func RegistrarTransacaoService(transacao transacoes.TransacaoDTO, response http.ResponseWriter) {
	err := transacoes.ValidarTransacaoRecebida(transacao, response)
	if err != nil {
		logs.RegistrarLogErro(err.Error())
		return
	} else {
		_, errTransacao := transacoes.RegistrarTransacao(transacoes.ConverterTransacao(transacao))
		if errTransacao != nil && strings.Contains(errTransacao.Error(), "Erro ao abater saldo. Saldo da conta é menor do que o valor da transação.") {
			requisicoesutil.RetornarComBadRequest(errTransacao.Error(), response)
			return
		}
		if errTransacao != nil {
			requisicoesutil.RetornarComInternalErrorServer(errTransacao.Error(), response)
			return
		}

		requisicoesutil.RetornarRegistroCriado(fmt.Sprintf(TRANSACAO_CRIADA, transacao.Account_id), response)

	}
}

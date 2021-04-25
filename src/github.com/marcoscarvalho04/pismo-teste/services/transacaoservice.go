package services

import (
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/requisicoesutil"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/transacoes"
)

const TRANSACAO_CRIADA string = "Transação registrada para a conta: "
const ERRO_GERAR_TRANSACAO string = "Erro ao gerar transacao para a conta: "

func RegistrarTransacaoService(transacao transacoes.TransacaoDTO, response http.ResponseWriter) {
	err := transacoes.ValidarTransacaoRecebida(transacao, response)
	if err != nil {
		logs.RegistrarLogErro(err.Error())
	} else {
		_, errTransacao := transacoes.RegistrarTransacao(transacoes.ConverterTransacao(transacao))
		if errTransacao != nil {
			requisicoesutil.RetornarRegistroCriado(TRANSACAO_CRIADA+string(transacao.Account_id), response)
		} else {
			requisicoesutil.RetornarComInternalErrorServer(ERRO_GERAR_TRANSACAO+string(transacao.Account_id), response)
		}
	}
}

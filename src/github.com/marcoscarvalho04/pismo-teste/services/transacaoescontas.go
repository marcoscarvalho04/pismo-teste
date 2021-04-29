package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/requisicoesutil"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/transacoes"
)

type ContaTransacaoDTO struct {
	Conta      contas.ContaDTO           `json: conta`
	Transacoes []transacoes.TransacaoDTO `json: transacoes`
}

func ConsultarTransacaoContaService(contaId int, response http.ResponseWriter) {

	var consultarContaTransacao ContaTransacaoDTO
	errConsultarConta := contas.IsContaExiste(contaId)

	if errConsultarConta != nil {
		requisicoesutil.RetornarComRegistroInexistente(fmt.Sprintf("ContaId: %v não existente na plataforma", contaId), response)
		return
	}
	contaConsultada, errConsultarConta := contas.ConsultarConta(contaId)
	if errConsultarConta != nil {
		requisicoesutil.RetornarComInternalErrorServer(fmt.Sprintf("Erro ao pesquisar conta: %v", errConsultarConta.Error()), response)
		return
	}
	transacoesDaConta := contaConsultada.Transacoes
	consultarContaTransacao.Conta = contas.ConvertDTO(contaConsultada)
	var todasTransacoes []transacoes.TransacaoDTO
	for _, value := range transacoesDaConta {
		transacaoConsultada, errConsultarTransacao := transacoes.ConsultarTransacao(value)
		if errConsultarTransacao == nil {
			todasTransacoes = append(todasTransacoes, transacoes.ConverterDTO(transacaoConsultada))
		} else {
			requisicoesutil.RetornarComBadRequest(fmt.Sprintf("Erro ao consultar transação: %v", errConsultarTransacao.Error()), response)
			return
		}
	}
	consultarContaTransacao.Transacoes = todasTransacoes

	convertido, errConverterJson := json.Marshal(consultarContaTransacao)
	if errConverterJson != nil {
		requisicoesutil.RetornarComInternalErrorServer(fmt.Sprintf("Erro ao fazer o parse do JSON: %v", errConverterJson.Error()), response)
		return
	}
	requisicoesutil.RetornarComStatusOK(string(convertido), response)
}

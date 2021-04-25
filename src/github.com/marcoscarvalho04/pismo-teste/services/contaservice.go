package services

import (
	"encoding/json"
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/requisicoesutil"
	"strconv"
)

type ContaReq struct {
	Document_number string `json:document_number`
}
type ContaConsulta struct {
	ContaId         int `json:account_id`
	Document_number int `json:document_number`
}

const ERRO_GERA_ID_CONTA string = "Erro ao gerar ID para a conta!"
const CONTA_GERADA_SUCESSO string = "Conta criada com o ID: "
const CONTA_NAO_EXISTE string = "Conta não existe no sistema!"

func RegistrarContaService(response http.ResponseWriter, conta contas.Contas) {

	_, errRegistro := contas.RegistrarConta(conta.NumeroDocumento)
	if errRegistro != nil {
		requisicoesutil.RetornarComInternalErrorServer(ERRO_GERA_ID_CONTA, response)
		return
	}
	requisicoesutil.RetornarRegistroCriado(CONTA_GERADA_SUCESSO+strconv.Itoa(conta.ContaId), response)

}

func ConsultarContaService(response http.ResponseWriter, contaId int) {
	conta, errConsultaConta := contas.ConsultarConta(contaId)
	if errConsultaConta != nil {
		requisicoesutil.RetornarComRegistroInexistente(CONTA_NAO_EXISTE, response)
		return
	}
	var contaConsultaRetorno contas.ContaDTO
	contaConsultaRetorno.ContaId = contaId
	contaConsultaRetorno.Document_number = conta.NumeroDocumento
	retorno, errParseJSON := json.Marshal(contaConsultaRetorno)
	if errParseJSON != nil {
		requisicoesutil.RetornarComInternalErrorServer("Não foi possível serializar o JSON de retorno!", response)
		return
	}
	requisicoesutil.RetornarComStatusOK(string(retorno), response)

}

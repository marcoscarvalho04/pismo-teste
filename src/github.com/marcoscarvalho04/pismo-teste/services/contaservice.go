package services

import (
	"encoding/json"
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
	"strconv"
)

type ContaReq struct {
	Document_number string `json:document_number`
}

func RegistrarContaService(response http.ResponseWriter, request *http.Request) {

	logs.RegistrarLogInformativo("STEP 1 - lendo requisição recebida")
	conteudo := make([]byte, request.ContentLength, request.ContentLength)
	request.Body.Read(conteudo)
	logs.RegistrarLogInformativo("STEP 2 - Requisição recebida com sucesso!")
	logs.RegistrarLogInformativo("STEP 2 - Fazendo parse do conteúdo recebido!")
	logs.RegistrarLogInformativo("Requisição recebida: ")
	logs.RegistrarLogInformativo(string(conteudo))
	var req ContaReq
	err := json.Unmarshal(conteudo, &req)
	if err != nil {
		logs.RegistrarLogErro("Não foi possível fazer o parse do conteúdo recebido")
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Não foi possível fazer o parse do conteúdo recebido : " + err.Error()))
		return
	}
	logs.RegistrarLogInformativo("STEP 3 - Parse concluido! Iniciando registro da conta")
	logs.RegistrarLogInformativo("Número de documento recebido: " + req.Document_number)
	document_number, errParse := strconv.Atoi(req.Document_number)
	if errParse != nil {
		logs.RegistrarLogErro("Erro ao fazer a conversão do número de documento! ")
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Erro ao fazer a conversão do número de documento!  " + errParse.Error()))
		return
	}
	novoId, errRegistro := contas.RegistrarConta(document_number)
	if errRegistro != nil {
		logs.RegistrarLogErro("Erro ao fazer registro da conta!")
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Erro ao fazer registro da conta!" + errRegistro.Error()))
		return
	}
	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("Conta criada com o ID: " + strconv.Itoa(novoId)))

}

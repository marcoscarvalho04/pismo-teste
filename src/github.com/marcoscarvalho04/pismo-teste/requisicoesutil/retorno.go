package requisicoesutil

import (
	"net/http"
)

func RetornarComBadRequest(msg string, response http.ResponseWriter) {
	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte(msg))
}
func RetornarComRegistroInexistente(msg string, response http.ResponseWriter) {
	response.WriteHeader(http.StatusNotFound)
	response.Write([]byte(msg))
}

func RetornarRegistroCriado(msg string, response http.ResponseWriter) {
	response.WriteHeader(http.StatusCreated)
	response.Write([]byte(msg))
}

func RetornarComInternalErrorServer(msg string, response http.ResponseWriter) {
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(msg))
}

func RetornarComStatusOK(msg string, response http.ResponseWriter) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(msg))
}

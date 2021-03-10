package requisicoes

import (
	"net/http"

	"github.com/gorilla/github.com/marcoscarvalho04/pismo-teste/logs"
)

func responderCriarTransacao(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Transação recebida, iniciando tratamento. ")

}

func responderCriarConta(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Pedido de criação de conta recebido. Iniciando tratamento")

}
func responderConsultarConta(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Pedido de consulta de conta recebido. Iniciando tratamento")

}

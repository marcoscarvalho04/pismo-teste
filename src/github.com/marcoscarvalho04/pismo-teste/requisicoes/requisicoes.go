package requisicoes

import (
	"net/http"

	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/services"
)

/*
Pacote: requisicoes.

Data de criação: 10/03/2021

Criador: Marcos Siqueira

Breve Descrição: Funcionalidade que redireciona e chama a interface de serviços para
que esta interface faça o que seja preciso para responder corretamente à requisição.
*/

func ResponderCriarTransacao(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Transação recebida, iniciando tratamento.")
	services.RegistrarTransacaoService(response, request)

}

func ResponderCriarConta(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Pedido de criação de conta recebido. Iniciando tratamento")
	services.RegistrarContaService(response, request)

}
func ResponderConsultarConta(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Pedido de consulta de conta recebido. Iniciando tratamento")
	services.ConsultarContaService(response, request)

}

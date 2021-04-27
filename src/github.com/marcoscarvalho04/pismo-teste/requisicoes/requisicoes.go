package requisicoes

import (
	"net/http"
	"strconv"

	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/requisicoesutil"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/services"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/transacoes"

	"github.com/gorilla/mux"
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
	conteudo := make([]byte, request.ContentLength, request.ContentLength)
	request.Body.Read(conteudo)
	transacao, err := transacoes.ParseTransaction(conteudo)
	if err != nil {
		requisicoesutil.RetornarComBadRequest("Não foi possível ler o JSON enviado! Verifique sua requisição.", response)
	}
	services.RegistrarTransacaoService(transacao, response)

}

func ResponderCriarConta(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Pedido de criação de conta recebido. Iniciando tratamento")
	conteudo := make([]byte, request.ContentLength, request.ContentLength)
	request.Body.Read(conteudo)
	conta, err := contas.ParseDTO(string(conteudo))
	if err != nil {
		requisicoesutil.RetornarComBadRequest("Impossível ler o JSON de cadastro da conta", response)
	} else {
		services.RegistrarContaService(response, contas.ConvertConta(conta))
	}

}
func ResponderConsultarConta(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Pedido de consulta de conta recebido. Iniciando tratamento")
	vars := mux.Vars(request)
	chave := vars["id"]
	if chave == "" {
		requisicoesutil.RetornarComBadRequest("Parâmetro não repassado! Erro ao consultar conta.", response)
	} else {
		contaId, err := strconv.Atoi(chave)
		if err != nil {
			requisicoesutil.RetornarComBadRequest("Não foi possível converter id em numérico", response)
		} else {
			services.ConsultarContaService(response, contaId)
		}

	}

}

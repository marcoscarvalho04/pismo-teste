package main

import (
	"log"
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/constantes"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/requisicoes"

	"github.com/gorilla/mux"
)

/*
Pacote: main.

Data de criação: 10/03/2021

Criador: Marcos Siqueira

Breve Descrição: Pacote principal, criado para rotearmos as requisições que vem do servidor e
fazer as chamadas para os respectivos métodos para tratamento do que for necessário.

*/

func main() {
	manipularRequisicoes()
}

func manipularRequisicoes() {
	router := mux.NewRouter().StrictSlash(true)
	logs.RegistrarLogInformativo("STEP 1 - Criar rotas para os métodos requisitados pelo documento funcional")
	router.HandleFunc("/accounts", requisicoes.ResponderCriarConta).Methods("POST")
	router.HandleFunc("/accounts/{id}", requisicoes.ResponderConsultarConta).Methods("GET")
	router.HandleFunc("/transactions", requisicoes.ResponderCriarTransacao).Methods("POST")
	router.HandleFunc("/transactions/accounts/{id}", requisicoes.ResponderConsultarTransacaoPorConta).Methods("GET")
	logs.RegistrarLogInformativo("STEP 2 - Inicializando configurações do sistema")
	constantes.ColetarInformacoesSistema()
	logs.RegistrarLogInformativo("STEP 3 - Iniciando servidor de registro de contas e transações na porta: " + constantes.Constantes.Config.Porta)
	log.Fatal(http.ListenAndServe(":"+constantes.Constantes.Config.Porta, router))

}

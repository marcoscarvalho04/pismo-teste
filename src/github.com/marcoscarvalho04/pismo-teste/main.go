package main

import (
	"github.com/gorilla/mux"
	"github.com/marcoscarvalho04/pismo-teste/logs"
)

/*
Pacote: main.

Data de criação: 10/03/2021

Criador: Marcos Siqueira

Breve Descrição: Pacote principal, criado para rotearmos as requisições que vem do servidor e
fazer as chamadas para os respectivos métodos para tratamento do que for necessário.

*/

func main() {
	logs.RegistrarLogInformativo("Iniciando sistema...")
	manipularRequisicoes()
}

func manipularRequisicoes() {
	router := mux.NewRouter().StrictSlash(true)

}

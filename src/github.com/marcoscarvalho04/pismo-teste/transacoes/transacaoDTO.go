package transacoes

import (
	"encoding/json"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
)

type TransacaoDTO struct {
	Account_id        int     `json:account_id`
	Operation_type_id int     `json:operation_type_id`
	Amount            float64 `json:amount`
}

/*
Arquivo: TransacaocaoDTO

Função: Fazer a transformação de uma requisição bruta em um DTO amigável para ida e vinda de informações relevantes.

Data da criação: 22/04/2021

Autor: Marcos Carvalho

*/
func ParseTransaction(conteudo []byte) (TransacaoDTO, error) {
	var transacao TransacaoDTO
	err := json.Unmarshal(conteudo, &transacao)
	if err != nil {
		return transacao, err
	}
	logs.RegistrarLogInformativo("Requisição recebida com sucesso! iniciando processamento da transação!")
	return transacao, nil
}

func ConverterDTO(transacao TransacoesModel) TransacaoDTO {
	var transacaoDTO TransacaoDTO
	transacaoDTO.Account_id = transacao.ContaId
	transacaoDTO.Amount = transacao.Valor
	transacaoDTO.Operation_type_id = transacao.OperacaoId
	return transacaoDTO
}

package contas

import (
	"encoding/json"
)

type ContaDTO struct {
	Account_id      int     `json:account_id`
	Document_number int     `json:document_number`
	Saldo           float64 `json:saldo`
}

func ConvertDTO(conta Contas) ContaDTO {
	var contaDTO ContaDTO
	contaDTO.Account_id = conta.ContaId
	contaDTO.Document_number = conta.NumeroDocumento
	contaDTO.Saldo = conta.Saldo
	return contaDTO
}

func ParseDTO(msg string) (ContaDTO, error) {
	var conta ContaDTO
	err := json.Unmarshal([]byte(msg), &conta)
	if err != nil {
		return conta, err
	}
	return conta, nil
}

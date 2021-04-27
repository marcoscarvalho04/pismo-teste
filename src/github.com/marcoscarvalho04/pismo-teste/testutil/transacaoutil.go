package testutil

import "fmt"

const JSON_ENTRADA_TRANSACAO string = "{\"account_id\": %d,\"operation_type_id\": %d,\"amount\": %.2f}"

func FormatarNovaEntradaTransacao(contaId int, operationId int, valor float64) string {
	return fmt.Sprintf(JSON_ENTRADA_TRANSACAO, contaId, operationId, valor)
}

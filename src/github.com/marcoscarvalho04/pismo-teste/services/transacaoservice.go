package services

import (
	"encoding/json"
	"net/http"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/contas"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/transacoes"
	"strconv"
	"time"
)

type TransacaoReq struct {
	Account_id        int     `json:account_id`
	Operation_type_id int     `json:operation_type_id`
	Amount            float32 `json:amount`
}

func RegistrarTransacaoService(response http.ResponseWriter, request *http.Request) {
	logs.RegistrarLogInformativo("Service - Iniciando tratamento de requisição do registro de informação da transação")
	logs.RegistrarLogInformativo("STEP 1 - Ler requisição")
	conteudo := make([]byte, request.ContentLength, request.ContentLength)
	request.Body.Read(conteudo)
	logs.RegistrarLogInformativo("Requisição lida")
	logs.RegistrarLogInformativo(string(conteudo))
	logs.RegistrarLogInformativo("STEP 2 - Fazendo conversão de JSON para a estrutura em que ele deveria vir")
	var req TransacaoReq
	err := json.Unmarshal(conteudo, &req)
	if err != nil {
		logs.RegistrarLogErro("Erro ao fazer a conversão do JSON recebido da transação!")
		logs.RegistrarLogErro(err.Error())
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Erro ao ler o JSON enviado: " + err.Error()))
	} else {
		logs.RegistrarLogInformativo("STEP 3 - JSON convertido. Validando informações antes de inserir a transação")
		logs.RegistrarLogInformativo("STEP 3.1 - Verificando se a conta já está registrada")
		err := contas.IsContaExiste(req.Account_id)
		if err != nil {
			logs.RegistrarLogErro("Transação não registrada para a conta de ID: " + strconv.Itoa(req.Account_id))
			logs.RegistrarLogErro("Motivo: Conta não existente na plataforma!")
			response.WriteHeader(http.StatusNotFound)
			response.Write([]byte("Transação não registrada. Conta não encontrada."))
			return
		}
		logs.RegistrarLogInformativo("STEP 3.2 - Verificando a consistência da informação enviada")
		logs.RegistrarLogInformativo("STEP 3.2 - Se saque ou compra, valor deve ser negativo")
		logs.RegistrarLogInformativo("STEP 3.2 - Se pagamento, positivo")
		logs.RegistrarLogInformativo("STEP 3.2 - Se operation_type_id for maior do que 4 ou menor do que 1, operação inválida")
		if req.Operation_type_id > 4 || req.Operation_type_id < 1 { // deve melhorar se houver mais transações
			logs.RegistrarLogErro("Transação não registrada para a conta de ID: " + strconv.Itoa(req.Account_id))
			logs.RegistrarLogErro("Motivo: Operação não mapeada/encontrada no sistema!")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Transação não registrada. Operação não encontrada."))
			return
		}
		if (req.Operation_type_id == 1 || req.Operation_type_id == 2 || req.Operation_type_id == 3) && (req.Amount > 0) {
			logs.RegistrarLogErro("Transação não registrada para a conta de ID: " + strconv.Itoa(req.Account_id))
			logs.RegistrarLogErro("Motivo: valor positivo não é permitido para as transações do tipo: saque e compra !")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Transação não registrada. valor positivo não é permitido para as transações do tipo: saque e compra !"))
			return
		}
		if req.Operation_type_id == 4 && req.Amount < 0 {
			logs.RegistrarLogErro("Transação não registrada para a conta de ID: " + strconv.Itoa(req.Account_id))
			logs.RegistrarLogErro("Motivo: valor negativo não é permitido para as transações do tipo pagamento!")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("Transação não registrada. Motivo: valor negativo não é permitido para as transações do tipo pagamento!"))
			return
		}
		// deve melhorar se houver mais transações
		// validações finalizadas, registrando transação para vincular àquela conta.
		var registrarTransacao transacoes.TransacoesModel
		registrarTransacao.ContaId = req.Account_id
		registrarTransacao.Data = time.Now()
		registrarTransacao.OperacaoId = req.Operation_type_id
		registrarTransacao.Valor = float64(req.Amount)

		transacaoId, errTransacao := transacoes.RegistrarTransacao(registrarTransacao)
		if errTransacao != nil {
			logs.RegistrarLogErro("Transação não registrada para a conta de ID: " + strconv.Itoa(req.Account_id))
			logs.RegistrarLogErro("Motivo: Erro ao registrar a transação: " + err.Error())
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("Erro ao registrar a transação: " + err.Error()))
			return
		}
		contas.VincularTransacao(req.Account_id, transacaoId)
		response.WriteHeader(http.StatusCreated)
		response.Write([]byte("Transação criada sob o ID: " + strconv.Itoa(transacaoId)))

	}
}

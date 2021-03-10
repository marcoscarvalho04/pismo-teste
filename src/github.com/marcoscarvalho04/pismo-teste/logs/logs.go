package logs

import (
	"fmt"
)

/*
Pacote: Logs.

Data de criação: 10/03/2021

Criador: Marcos Siqueira

Breve Descrição: Pacote criado para registro de logs na aplicação. Tem por funcionalidade padronizar
a exibição de logs no sistema como um todo, tornando-o homogêneo neste sentido.
*/

func RegistrarLogInformativo(mensagem string) {
	fmt.Println("[INFO] " + mensagem)
}
func RegistrarLogErro(mensagem string) {
	fmt.Println("[ERROR] - " + mensagem)
}

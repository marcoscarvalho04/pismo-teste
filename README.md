# Teste da Pismo - Marcos Siqueira

Esta é a documentação feita para registrar transações e contas para o teste de ingresso na pismo. 

## Pre-requisitos

Antes de iniciar, tenha certeza de que cumpre os requisitos abaixo: 


* Tenha instalada a versão  `<golang/1.14.1>` ou maior no seu computador. 
* tenha instalado algum software de teste para API's, como exemplo do postman.

## Instalando o pismo_teste

Para instalar este projeto nomeado  de pismo_teste, siga esses passos:

```
1. Tenha feito o clone completo do código neste repositório, da branch master. 
2. Após o clone deste repositório, não será necessário setar as variáveis de ambiente GOPATH, uma vez que ele já está adequado para utilização de módulos, facilitando a sua instalação. 
3. Com a versão da linguagem golang mencionada na seção anterior devidamente instalada, juntamente com o software de teste para consumo de API's, será necessário, através do prompt de comando do seu sistema operacional, navegar até á raiz do projeto. Após isso, na pasta <github.com/marcoscarvalho04/pismo-teste> execute o comando <go run main.go> para dar início ao sistema. 
```

## Estrutura 

Abaixo, segue a estrutura de pacotes e sua devida explicação para o projeto.

```
src
--github.com
---marcoscarvalho04
----pismo-teste
-----constantes
-----contas
-----logs
-----requisicoes
-----services
-----transacoes
----main.go
```

Onde: 

- **Constantes**: Pacote responsável por fazer a leitura do arquivo de configuração yaml e ter informações mais gerais do projeto. 
- **Contas**: pacote responsável por fazer operações nas contas dos usuários. 
- logs: pacote responsável por logar as informações do sistema de maneira geral e consistente. 
- **Requisições**: pacote intermediário responsável por captar a requisição recebida no projeto principal e redirecionar para o pacote de serviços
- **Services**: pacote de serviços. Faz o tratamento da requisição em si, respondendo com os códigos corretos e consultando os pacotes corretos que precisa para isso. Reside a maior parte das regras de negócio. 
- **Transacoes**: pacote responsável por fazer operações sobre as transações enviadas ao sistema. 

## Restrições e melhorias

Para o projeto, não foram considerados os seguintes procedimentos: 

1) Gravação de dados em banco de dados: por ter sido solicitado algo mais simples e direto, isso foi dispensado, no entanto se necessário pode ser implementado. 

2) O path de consulta de contas não mostra as transações vinculadas ao portador, uma vez que a documentação exije isso. Isto também pode ser implementado, caso solicitado. 



## Contato

Se quiser entrar em contato basta mandar e-mail para: <siqueira.marcos.04@gmail.com>.


package constantes

import (
	"io/ioutil"
	"os"
	"pismo-teste/github.com/marcoscarvalho04/pismo-teste/logs"

	"gopkg.in/yaml.v2"
)

type YamlConstantes struct {
	Informacoes struct {
		Versao    string `yaml:"versao"`
		Descricao string `yaml:"descricao"`
	}
	Config struct {
		Porta string `yaml:"porta"`
	}
}

var Constantes YamlConstantes

func ColetarInformacoesSistema() {

	logs.RegistrarLogInformativo("Fazendo parse do arquivo: config.yaml")
	var nomeArquivo string
	nomeArquivo = "config.yaml"
	yamlFile, err := ioutil.ReadFile(nomeArquivo)
	if err != nil {
		logs.RegistrarLogErro("Erro ao ler o arquivo: config.yaml!")
		os.Exit(1)

	}

	err = yaml.Unmarshal(yamlFile, &Constantes)
	if err != nil {
		logs.RegistrarLogErro("Impossível fazer o parse do arquivo: config.yaml!")
	}
	logs.RegistrarLogInformativo("Constantes de servidor lidas com sucesso!")
	logs.RegistrarLogInformativo("Informações lidas abaixo")
	logs.RegistrarLogInformativo("Versão do sistema: " + Constantes.Informacoes.Versao)
	logs.RegistrarLogInformativo("Descrição do sistema: " + Constantes.Informacoes.Descricao)
}

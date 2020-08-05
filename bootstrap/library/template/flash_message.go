package template

import "os"

type System struct {
	Name    string
	Version string
}

type FlashMessage struct {
	Type    string
	Message string
}

type Parametro struct {
	System       System
	FlashMessage FlashMessage
	Parametro    interface{}
}

// ObterTipoMensagemGravacaoSucesso -
func ObterTipoMensagemGravacaoSucesso() (string, string) {
	return os.Getenv("CSS_MENSAGEM_SUCESSO"), os.Getenv("MENSAGEM_GRAVACAO_SUCESSO")
}

// ObterTipoMensagemGravacaoErro -
func ObterTipoMensagemGravacaoErro() (string, string) {
	return os.Getenv("CSS_MENSAGEM_ERRO"), os.Getenv("MENSAGEM_GRAVACAO_ERRO")
}

// ObterTipoMensagemExclusaoSucesso -
func ObterTipoMensagemExclusaoSucesso() (string, string) {
	return os.Getenv("CSS_MENSAGEM_SUCESSO"), os.Getenv("MENSAGEM_EXCLUSAO_SUCESSO")
}

// ObterTipoMensagemExclusaoErro -
func ObterTipoMensagemExclusaoErro() (string, string) {
	return os.Getenv("CSS_MENSAGEM_ERRO"), os.Getenv("MENSAGEM_EXCLUSAO_ERRO")
}

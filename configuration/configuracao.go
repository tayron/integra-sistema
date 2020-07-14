package configuration

// Configuracao - Armazena configuração do sistema
type Configuracao struct {
	NomeAplicativo     string
	VersaoAplicativo   string
	IPServidor         string
	PortaServidor      string
	EnderecoBancoDados string
	PortaBancoDados    string
	NomeBancoDados     string
	UsuarioBancoDados  string
	SenhaBancoDados    string
}

// GetNomeAplicativo -
func (c Configuracao) GetNomeAplicativo() string {
	return c.NomeAplicativo
}

// GetVersaoAplicativo -
func (c Configuracao) GetVersaoAplicativo() string {
	return c.VersaoAplicativo
}

// GetIPServidor -
func (c Configuracao) GetIPServidor() string {
	return c.IPServidor
}

// GetPortaServidor -
func (c Configuracao) GetPortaServidor() string {
	return c.PortaServidor
}

// GetEnderecoBancoDados -
func (c Configuracao) GetEnderecoBancoDados() string {
	return c.EnderecoBancoDados
}

// GetPortaBancoDados -
func (c Configuracao) GetPortaBancoDados() string {
	return c.PortaBancoDados
}

// GetNomeBancoDados -
func (c Configuracao) GetNomeBancoDados() string {
	return c.NomeBancoDados
}

// GetUsuarioBancoDados -
func (c Configuracao) GetUsuarioBancoDados() string {
	return c.UsuarioBancoDados
}

// GetSenhaBancoDados -
func (c Configuracao) GetSenhaBancoDados() string {
	return c.SenhaBancoDados
}

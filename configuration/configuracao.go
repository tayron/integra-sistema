package configuration

import "os"

// SetConfiguracao -
func SetConfiguracao(chave, valor string) {
	os.Setenv(chave, valor)
}

// GetConfiguracao -
func GetConfiguracao(chave string) string {
	return os.Getenv(chave)
}

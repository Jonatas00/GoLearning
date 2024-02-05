package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TipoDeEndereco Verifica se o endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return cases.Title(language.Und).String(primeiraPalavraEndereco)
	}

	return "Tipo Inválido"
}
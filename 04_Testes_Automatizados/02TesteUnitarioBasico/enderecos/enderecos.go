package enderecos

import "strings"

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
		return primeiraPalavraEndereco
	}

	return "Tipo Inválido"
}
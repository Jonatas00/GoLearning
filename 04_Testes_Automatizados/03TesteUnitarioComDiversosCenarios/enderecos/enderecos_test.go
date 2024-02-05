package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	
	cenariosDeTeste := []cenarioDeTeste {
		{ "Rua ABC", "Rua"},
		{ "Avenida Paulista", "Avenida"},
		{ "Avenida Consolação", "Avenida"},
		{ "Rua Marília", "Rua"},
		{ "Praça das Rosas", "Tipo Inválido"},
		{ "Rodovia Anhanguera", "Rodovia"},
		{ "Estrada Coronel J. Gladiador", "Estrada"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
			 retornoRecebido,
			 cenario.retornoEsperado,
			)
		}
	}

}
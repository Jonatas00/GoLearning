package enderecos_test

import (
	. "introducao-a-testes/enderecos"
	"testing"
)
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

// go test ./...
// go test ./... -v                        <- Relatório mais detalhado
// go test --cover                         <- Verificar cobertura do teste
// go test --coverprofile cobertura.txt    <- cria arquivo de cobertura
// go tool cover --html=cobertura.txt      <- visualiza em formato html

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

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

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
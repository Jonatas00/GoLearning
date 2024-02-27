package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (u *Usuario) validar(etapa string) error {
	var camposVazios []string

	if u.Nome == "" {
		camposVazios = append(camposVazios, "nome")
	}
	if u.Nick == "" {
		camposVazios = append(camposVazios, "nick")
	}
	if u.Email == "" {
		camposVazios = append(camposVazios, "email")
	}

	if etapa == "cadastro" && u.Senha == "" {
		camposVazios = append(camposVazios, "senha")
	}

	if len(camposVazios) > 0 {
		mensagem := "Os seguintes campos estão vazios e não podem estar em branco: "
		for _, campo := range camposVazios {
			mensagem += campo + ", "
		}

		return errors.New(mensagem)
	}

	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}

func (u *Usuario) Preparar(etapa string) error {
	if erro := u.validar(etapa); erro != nil {
		return erro
	}
	u.formatar()
	return nil
}

package models

import (
	"net/http"
	"time"
)

type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

func BuscarUsuario(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go buscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go buscarSeguidores(canalSeguidores, usuarioID, r)
	go buscarSeguindo(canalSeguindo, usuarioID, r)
	go buscarPublicacoes(canalPublicacoes, usuarioID, r)
}

func buscarDadosDoUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {

}

func buscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func buscarSeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func buscarPublicacoes(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {

}

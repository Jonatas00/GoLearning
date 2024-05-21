package router

import (
	"webapp/internal/router/routes"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routes.Configurar(r)
}

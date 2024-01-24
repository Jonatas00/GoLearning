package app

import "github.com/urfave/cli"

// Gerar irá retornar a aplicação de linha de comnado pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	return app
}

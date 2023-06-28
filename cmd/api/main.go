package main

import (
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/infra/webServer"
)

func main() {
	if err := webServer.Run(); err != nil {
		panic(err)
	}
}

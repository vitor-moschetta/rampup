package controllers

import (
	"fmt"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) HelloController(w http.ResponseWriter, r *http.Request) error {
	return web.EncodeJSON(w, fmt.Sprintf("%s, world!", r.URL.Path[1:]), http.StatusOK)
}

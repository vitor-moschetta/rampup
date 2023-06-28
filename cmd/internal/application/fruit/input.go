package fruit

import (
	"net/http"

	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/application/general"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/domain/fruit"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/pkg/validators"
)

type CreateInput struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

func (c *CreateInput) Validate() (output general.Output) {
	if c.Name == "" {
		output.AddError("name is required")
	}
	if !validators.IsAlpha(c.Name) {
		output.AddError("name must be alphabetic characters")
	}
	if c.Quantity <= 0 {
		output.AddError("quantity must be greater than 0")
	}
	if c.Price <= 0 {
		output.AddError("price must be greater than 0")
	}
	if output.HasErrors() {
		output.SetStatusCode(http.StatusBadRequest)
	}
	return
}

func ToFruitModel(input CreateInput, owner string) fruit.Fruit {
	return fruit.NewFruit(
		input.Name,
		input.Quantity,
		input.Price,
		owner)
}

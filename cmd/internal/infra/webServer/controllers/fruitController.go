package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-core/pkg/web"
	fruitApplication "github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/application/fruit"
	fruitDomain "github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/domain/fruit"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/domain/general"
)

// Using CQRS pattern: UseCase is responsible for the business logic (change the state of the data).
// For creating the data, the controller calls the use case, which calls the repository.
// For getting the data, the controller calls the repository directly.
type FruitController struct {
	Repository general.IRepository
	UseCase    *fruitApplication.FruitUseCase
}

func NewFruitController(repository general.IRepository, useCase *fruitApplication.FruitUseCase) *FruitController {
	return &FruitController{
		Repository: repository,
		UseCase:    useCase,
	}
}

// PostFruitsController conteins the logic for creating a fruit. UseCase is called.
func (c *FruitController) PostFruitsController(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	owner := r.Header.Get("owner")
	if owner == "" {
		return web.NewError(http.StatusBadRequest, "owner is required in header")
	}

	input := fruitApplication.CreateInput{}
	if err := web.DecodeJSON(r, &input); err != nil {
		log.Error(ctx, err.Error())
		return web.NewError(http.StatusInternalServerError, "Error decoding request")
	}

	output := c.UseCase.CreateFruit(ctx, input, owner)

	return web.EncodeJSON(w, output, output.StatusCode)
}

// GetFruitController conteins the logic only for getting a fruit. Repository only is called.
func (c *FruitController) GetFruitController(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id := r.URL.Path[len("/fruits/"):]
	if id == "" {
		return web.NewError(http.StatusBadRequest, "id is required")
	}

	item, err := c.Repository.Get(ctx, id)
	if err != nil {
		log.Error(ctx, err.Error())
		return web.NewError(http.StatusInternalServerError, "Internal error.")
	}

	var response fruitDomain.Fruit
	err = json.Unmarshal(item, &response)
	if err != nil {
		log.Error(ctx, err.Error())
		return web.NewError(http.StatusInternalServerError, "Internal error.")
	}

	if response.ID == "" {
		return web.NewError(http.StatusNotFound, "Fruit not found.")
	}

	return web.EncodeJSON(w, response, http.StatusOK)
}

// GetFruitsController conteins the logic only for getting a list of fruits. Repository only is called.
func (c *FruitController) GetFruitsController(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	keys := r.URL.Query()["key"]
	if len(keys) == 0 {
		return web.NewError(http.StatusBadRequest, "key is required")
	}

	items, err := c.Repository.GetAll(ctx, keys)
	if err != nil {
		log.Error(ctx, err.Error())
		return web.NewError(http.StatusInternalServerError, "Internal error.")
	}

	var response []fruitDomain.Fruit
	err = json.Unmarshal(items, &response)
	if err != nil {
		log.Error(ctx, err.Error())
		return web.NewError(http.StatusInternalServerError, "Internal error.")
	}

	return web.EncodeJSON(w, response, http.StatusOK)
}

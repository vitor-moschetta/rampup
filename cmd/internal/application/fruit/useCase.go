package fruit

import (
	"context"
	"log"
	"net/http"

	applicationGeneral "github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/application/general"
	domainGeneral "github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/domain/general"
)

type FruitUseCase struct {
	Repository domainGeneral.IRepository
}

func NewFruitUseCase(repository domainGeneral.IRepository) *FruitUseCase {
	return &FruitUseCase{
		Repository: repository,
	}
}

func (f *FruitUseCase) CreateFruit(ctx context.Context, input CreateInput, owner string) (output applicationGeneral.Output) {
	output = input.Validate()
	if output.HasErrors() {
		log.Println(ctx, output)
		output.SetStatusCode(http.StatusBadRequest)
		return
	}

	entity := ToFruitModel(input, owner)
	err := f.Repository.Save(ctx, entity.ID, entity)
	if err != nil {
		log.Println(ctx, err)
		output.AddError("internal error")
		output.SetStatusCode(http.StatusInternalServerError)
		return
	}

	output.SetStatusCode(http.StatusCreated)
	output.SetData(entity)
	return
}

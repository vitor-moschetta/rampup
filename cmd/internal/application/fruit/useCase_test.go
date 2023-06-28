package fruit

import (
	"context"
	"net/http"
	"testing"

	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/infra/repositories"
	"github.com/stretchr/testify/assert"
)

func Test_With_Fuit_Add_With_Valid_Data(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "Banana",
		Quantity: 10,
		Price:    10.0,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.False(t, output.HasErrors())
	assert.Equal(t, http.StatusCreated, output.StatusCode)
}

func Test_With_Fuit_Add_With_Invalid_Name(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "",
		Quantity: 10,
		Price:    10.0,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.True(t, output.HasErrors())
	assert.Equal(t, http.StatusBadRequest, output.StatusCode)
}

func Test_With_Fuit_Add_With_Invalid_Name2(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "Banana123",
		Quantity: 10,
		Price:    10.0,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.True(t, output.HasErrors())
	assert.Equal(t, http.StatusBadRequest, output.StatusCode)
}

func Test_With_Fuit_Add_With_Invalid_Name3(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "Banana*",
		Quantity: 10,
		Price:    10.0,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.True(t, output.HasErrors())
	assert.Equal(t, http.StatusBadRequest, output.StatusCode)
}

func Test_With_Fuit_Add_With_Invalid_Quantity(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "Banana",
		Quantity: 0,
		Price:    10.0,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.True(t, output.HasErrors())
	assert.Equal(t, http.StatusBadRequest, output.StatusCode)
}

func Test_With_Fuit_Add_With_Invalid_Quantity2(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "Banana",
		Quantity: -1,
		Price:    10.0,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.True(t, output.HasErrors())
	assert.Equal(t, http.StatusBadRequest, output.StatusCode)
}

func Test_With_Fuit_Add_With_Invalid_Price(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "Banana",
		Quantity: 10,
		Price:    0,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.True(t, output.HasErrors())
	assert.Equal(t, http.StatusBadRequest, output.StatusCode)
}

func Test_With_Fuit_Add_With_Invalid_Price2(t *testing.T) {
	// Arrange
	ctx := context.Background()
	repository := repositories.NewMemoryRepository()
	useCase := NewFruitUseCase(repository)
	input := CreateInput{
		Name:     "Banana",
		Quantity: 10,
		Price:    -1,
	}
	// Act
	output := useCase.CreateFruit(ctx, input, "owner")
	// Assert
	assert.True(t, output.HasErrors())
	assert.Equal(t, http.StatusBadRequest, output.StatusCode)
}

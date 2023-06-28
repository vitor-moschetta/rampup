package fruit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_Fruit_Model_Instance(t *testing.T) {
	// Arrange
	// Act
	entity := NewFruit("Fruit", 10, 10.0, "owner")
	// Assert
	assert.Equal(t, entity.Name, "Fruit")
	assert.Equal(t, entity.Quantity, 10)
	assert.Equal(t, entity.Price, float32(10.0))
	assert.Equal(t, entity.Owner, "owner")
}

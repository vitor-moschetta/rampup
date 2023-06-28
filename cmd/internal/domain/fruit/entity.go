package fruit

import (
	"time"

	"github.com/google/uuid"
)

type Fruit struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Quantity        int     `json:"quantity"`
	Price           float32 `json:"price"`
	DateCreated     string  `json:"date_created"`
	DateLastUpdated string  `json:"date_last_updated"`
	Owner           string  `json:"owner"`
}

func NewFruit(name string, quantity int, price float32, owner string) Fruit {
	return Fruit{
		Name:            name,
		Quantity:        quantity,
		Price:           price,
		ID:              uuid.New().String(),
		DateCreated:     time.Now().Format("2006-01-02T15:04:05"),
		DateLastUpdated: time.Now().Format("2006-01-02T15:04:05"),
		Owner:           owner,
	}
}

func (f *Fruit) Update(name string, quantity int, price float32) {
	f.Name = name
	f.Quantity = quantity
	f.Price = price
	f.DateLastUpdated = time.Now().Format("2006-01-02T15:04:05")
}

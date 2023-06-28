package repositories

import (
	"context"
	"encoding/json"
	"log"

	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/domain/general"
)

type MemoryRepository struct {
	Store map[string]interface{}
}

func NewMemoryRepository() general.IRepository {
	return &MemoryRepository{
		Store: make(map[string]interface{}),
	}
}

func (r *MemoryRepository) Save(ctx context.Context, key string, value interface{}) (err error) {
	r.Store[key] = value
	return nil
}

func (r *MemoryRepository) Get(ctx context.Context, key string) (value []byte, err error) {
	item := r.Store[key]

	jsonData, err := json.Marshal(item)
	if err != nil {
		log.Println(err)
		return
	}

	value = jsonData
	return
}

func (r *MemoryRepository) GetAll(ctx context.Context, keys []string) (values []byte, err error) {
	items := []interface{}{}
	for _, key := range keys {
		item := r.Store[key]
		items = append(items, item)
	}

	jsonData, err := json.Marshal(items)
	if err != nil {
		log.Println(err)
		return
	}

	values = jsonData
	return
}

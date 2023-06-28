package repositories

import (
	"context"
	"log"

	"github.com/mercadolibre/fury_go-toolkit-kvs/pkg/kvs"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/domain/general"
)

type KvsRepository struct {
	Store kvs.Client
}

func NewKvsRepository(kvs kvs.Client) general.IRepository {
	return &KvsRepository{
		Store: kvs,
	}
}

func (r *KvsRepository) Save(ctx context.Context, key string, value interface{}) (err error) {
	err = r.Store.Set(ctx, key, value)
	if err != nil {
		log.Println(err)
	}
	return
}

func (r *KvsRepository) Get(ctx context.Context, key string) (value []byte, err error) {
	item, err := r.Store.Get(ctx, key)
	if err != nil {
		log.Println(err)
	}
	value = item.Value
	return
}

func (r *KvsRepository) GetAll(ctx context.Context, keys []string) (values []byte, err error) {
	bulk, err := r.Store.BulkGet(ctx, keys)
	if err != nil {
		log.Println(err)
	}
	for _, item := range bulk.Items {
		values = append(values, item.Value...)
	}
	return
}

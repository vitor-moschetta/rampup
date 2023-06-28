package general

import "context"

type IRepository interface {
	Save(ctx context.Context, key string, value interface{}) (err error)
	Get(ctx context.Context, key string) (value []byte, err error)
	GetAll(ctx context.Context, keys []string) (values []byte, err error)
}

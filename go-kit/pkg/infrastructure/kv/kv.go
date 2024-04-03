package kv

import (
	"context"
	"time"

	//"github.com/redis/go-redis/v9"
	goredis "github.com/redis/go-redis/v9"
)

type IKeyValueStore interface {
	Ping() error
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Remove(ctx context.Context, keys ...string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	Scan(ctx context.Context, pattern string) ([]string, error)
}

//go:generate mockery --name=IRedisStore --output=./redismocks
type IRedisStore interface {
	IKeyValueStore
	GetClient() (string, *goredis.Client)
	Zset(ctx context.Context, key string, value interface{}, score float64) error
	Zget(ctx context.Context, key string, start, stop int64) ([]goredis.Z, error)
	Zdel(ctx context.Context, key string, value interface{}) error
	Zlist(ctx context.Context, key string, start, stop int64) (map[string]interface{}, error)
	DeleteHash(ctx context.Context, key string, fields ...string) error
	Delete(ctx context.Context, key string) error
	ScanHash(ctx context.Context, key string, filterStrategy IFilterStrategy) ([][]byte, error)
	ExistHash(ctx context.Context, key string, fields ...interface{}) (bool, error)
	StoreHash(ctx context.Context, key string, fields ...interface{}) error
	ExecuteLuaScript(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error)
}

package kv

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
)

const Timeout = 1

type redis struct {
	timeout      time.Duration
	globalPrefix string
	client       goredis.UniversalClient
}

type OptsRedis func(*redis)

// New Redis interface with config
func NewRedis(client goredis.UniversalClient, opts ...OptsRedis) (IRedisStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	r := &redis{
		timeout: Timeout * time.Second,
		client:  client,
	}
	for _, opt := range opts {
		opt(r)
	}

	return r, nil
}

func WithGlobalPrefix(prefix string) OptsRedis {
	return func(r *redis) {
		if len(prefix) > 0 && prefix[len(prefix)-1] != ':' {
			prefix = prefix + ":"
		}
		r.globalPrefix = prefix
	}
}

func WithTimeout(timeout time.Duration) OptsRedis {
	return func(r *redis) {
		r.timeout = timeout
	}
}

func (r *redis) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout*time.Second)
	defer cancel()

	if r.client == nil {
		return fmt.Errorf("redis client is nil")
	}

	_, err := r.client.Ping(ctx).Result()
	return err
}

func (r *redis) Get(ctx context.Context, key string, value interface{}) error {
	strValue, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	logger.Info(strValue)

	switch v := value.(type) {
	case *string:
		*v = strValue
		return nil
	default:
		return fmt.Errorf("unsupported type for value")
	}
}

func (r *redis) SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	key = r.globalPrefix + key
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *redis) Set(ctx context.Context, key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	key = r.globalPrefix + key
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *redis) DeleteHash(ctx context.Context, key string, fields ...string) error {
	key = r.globalPrefix + key
	exist, err := r.client.HExists(ctx, key, fmt.Sprintf("%v", fields[0])).Result()
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("key %v not found", fields[0])
	}
	return r.client.HDel(ctx, key, fields...).Err()
}

func (r *redis) Delete(ctx context.Context, key string) error {
	key = r.globalPrefix + key
	return r.client.Del(ctx, key).Err()
}

func (r *redis) StoreHash(ctx context.Context, key string, fields ...interface{}) error {
	key = r.globalPrefix + key
	_, err := r.client.HSet(ctx, key, fields).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *redis) ExistHash(ctx context.Context, key string, fields ...interface{}) (bool, error) {
	key = r.globalPrefix + key
	return r.client.HExists(ctx, key, fmt.Sprintf("%v", fields[0])).Result()
}

func (r *redis) ScanHash(ctx context.Context, key string, filter IFilterStrategy) ([][]byte, error) {
	key = r.globalPrefix + key
	var result [][]byte
	var err error
	var cursor uint64 = 0
	var pattern string

	count, preFilterKey, preFilter := filter.GetFilter()
	if preFilter {
		pattern = preFilterKey
	} else {
		pattern = ""
	}
	iter := r.client.HScan(ctx, key, cursor, pattern, count).Iterator()
	for iter.Next(ctx) {
		// Since HScan returns keys and values in pairs, we use one call to Next for the key
		// and a second call to Next for the value, hence advancing the iterator twice per loop iteration.
		//key := iter.Val()
		if !iter.Next(ctx) {
			break // This should not happen, but we check to avoid an odd number of elements causing a panic.
		}
		value := iter.Val()
		if filter.ShouldInclude([]byte(value)) {
			//fields[key] = value
			result = append(result, []byte(value))
		}
	}
	if err = iter.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *redis) Remove(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *redis) Keys(ctx context.Context, pattern string) ([]string, error) {
	return r.client.Keys(ctx, pattern).Result()
}

func (r *redis) Scan(ctx context.Context, pattern string) ([]string, error) {
	ret := make([]string, 0)
	keys := r.client.Scan(ctx, 0, pattern, -1).Iterator()
	for keys.Next(ctx) {
		err := keys.Err()
		if err != nil {
			return nil, err
		}
		ret = append(ret, keys.Val())
	}
	return ret, nil
}

// Execute lua script
func (r *redis) ExecuteLuaScript(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	return r.client.Eval(ctx, script, keys, args...).Result()
}

// zset
func (r *redis) Zset(ctx context.Context, key string, value interface{}, score float64) error {
	key = r.globalPrefix + key
	return r.client.ZAdd(ctx, key, goredis.Z{Score: score, Member: value}).Err()
}

// zdel
func (r *redis) Zdel(ctx context.Context, key string, value interface{}) error {
	key = r.globalPrefix + key
	z := goredis.Z{Member: value}
	return r.client.ZRem(ctx, key, z).Err()
}

// zget
func (r *redis) Zget(ctx context.Context, key string, start, stop int64) ([]goredis.Z, error) {
	key = r.globalPrefix + key
	return r.client.ZRangeWithScores(ctx, key, start, stop).Result()
}

// zlist by using zscan
func (r *redis) Zlist(ctx context.Context, key string, start, stop int64) (map[string]interface{}, error) {
	ret := make(map[string]interface{}, 0)
	keys := r.client.ZScan(ctx, key, 0, "", -1).Iterator()
	for keys.Next(ctx) {
		err := keys.Err()
		if err != nil {
			return nil, err
		}
		ret[keys.Val()] = keys.Val()
	}
	return ret, nil
}

func (r *redis) GetClient() (string, goredis.UniversalClient) {
	return r.globalPrefix, r.client
}

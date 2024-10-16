// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	kv "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	mock "github.com/stretchr/testify/mock"

	redis "github.com/redis/go-redis/v9"

	time "time"
)

// IRedisStore is an autogenerated mock type for the IRedisStore type
type IRedisStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, key
func (_m *IRedisStore) Delete(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteHash provides a mock function with given fields: ctx, key, fields
func (_m *IRedisStore) DeleteHash(ctx context.Context, key string, fields ...string) error {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteHash")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) error); ok {
		r0 = rf(ctx, key, fields...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecuteLuaScript provides a mock function with given fields: ctx, script, keys, args
func (_m *IRedisStore) ExecuteLuaScript(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, script, keys)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteLuaScript")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, ...interface{}) (interface{}, error)); ok {
		return rf(ctx, script, keys, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, ...interface{}) interface{}); ok {
		r0 = rf(ctx, script, keys, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []string, ...interface{}) error); ok {
		r1 = rf(ctx, script, keys, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistHash provides a mock function with given fields: ctx, key, fields
func (_m *IRedisStore) ExistHash(ctx context.Context, key string, fields ...interface{}) (bool, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, fields...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExistHash")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (bool, error)); ok {
		return rf(ctx, key, fields...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) bool); ok {
		r0 = rf(ctx, key, fields...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, key, fields...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, key, value
func (_m *IRedisStore) Get(ctx context.Context, key string, value interface{}) error {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetClient provides a mock function with given fields:
func (_m *IRedisStore) GetClient() (string, redis.UniversalClient) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetClient")
	}

	var r0 string
	var r1 redis.UniversalClient
	if rf, ok := ret.Get(0).(func() (string, redis.UniversalClient)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() redis.UniversalClient); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(redis.UniversalClient)
		}
	}

	return r0, r1
}

// Keys provides a mock function with given fields: ctx, pattern
func (_m *IRedisStore) Keys(ctx context.Context, pattern string) ([]string, error) {
	ret := _m.Called(ctx, pattern)

	if len(ret) == 0 {
		panic("no return value specified for Keys")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, pattern)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, pattern)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pattern)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields:
func (_m *IRedisStore) Ping() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Remove provides a mock function with given fields: ctx, keys
func (_m *IRedisStore) Remove(ctx context.Context, keys ...string) error {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, keys...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scan provides a mock function with given fields: ctx, pattern
func (_m *IRedisStore) Scan(ctx context.Context, pattern string) ([]string, error) {
	ret := _m.Called(ctx, pattern)

	if len(ret) == 0 {
		panic("no return value specified for Scan")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]string, error)); ok {
		return rf(ctx, pattern)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, pattern)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, pattern)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScanHash provides a mock function with given fields: ctx, key, filterStrategy
func (_m *IRedisStore) ScanHash(ctx context.Context, key string, filterStrategy kv.IFilterStrategy) ([][]byte, error) {
	ret := _m.Called(ctx, key, filterStrategy)

	if len(ret) == 0 {
		panic("no return value specified for ScanHash")
	}

	var r0 [][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, kv.IFilterStrategy) ([][]byte, error)); ok {
		return rf(ctx, key, filterStrategy)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, kv.IFilterStrategy) [][]byte); ok {
		r0 = rf(ctx, key, filterStrategy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, kv.IFilterStrategy) error); ok {
		r1 = rf(ctx, key, filterStrategy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, key, value
func (_m *IRedisStore) Set(ctx context.Context, key string, value interface{}) error {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWithExpiration provides a mock function with given fields: ctx, key, value, expiration
func (_m *IRedisStore) SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	ret := _m.Called(ctx, key, value, expiration)

	if len(ret) == 0 {
		panic("no return value specified for SetWithExpiration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreHash provides a mock function with given fields: ctx, key, fields
func (_m *IRedisStore) StoreHash(ctx context.Context, key string, fields ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, fields...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StoreHash")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) error); ok {
		r0 = rf(ctx, key, fields...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Zdel provides a mock function with given fields: ctx, key, value
func (_m *IRedisStore) Zdel(ctx context.Context, key string, value interface{}) error {
	ret := _m.Called(ctx, key, value)

	if len(ret) == 0 {
		panic("no return value specified for Zdel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) error); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Zget provides a mock function with given fields: ctx, key, start, stop
func (_m *IRedisStore) Zget(ctx context.Context, key string, start int64, stop int64) ([]redis.Z, error) {
	ret := _m.Called(ctx, key, start, stop)

	if len(ret) == 0 {
		panic("no return value specified for Zget")
	}

	var r0 []redis.Z
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) ([]redis.Z, error)); ok {
		return rf(ctx, key, start, stop)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) []redis.Z); ok {
		r0 = rf(ctx, key, start, stop)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]redis.Z)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) error); ok {
		r1 = rf(ctx, key, start, stop)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Zlist provides a mock function with given fields: ctx, key, start, stop
func (_m *IRedisStore) Zlist(ctx context.Context, key string, start int64, stop int64) (map[string]interface{}, error) {
	ret := _m.Called(ctx, key, start, stop)

	if len(ret) == 0 {
		panic("no return value specified for Zlist")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) (map[string]interface{}, error)); ok {
		return rf(ctx, key, start, stop)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64, int64) map[string]interface{}); ok {
		r0 = rf(ctx, key, start, stop)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64, int64) error); ok {
		r1 = rf(ctx, key, start, stop)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Zset provides a mock function with given fields: ctx, key, value, score
func (_m *IRedisStore) Zset(ctx context.Context, key string, value interface{}, score float64) error {
	ret := _m.Called(ctx, key, value, score)

	if len(ret) == 0 {
		panic("no return value specified for Zset")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, float64) error); ok {
		r0 = rf(ctx, key, value, score)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIRedisStore creates a new instance of IRedisStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRedisStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRedisStore {
	mock := &IRedisStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

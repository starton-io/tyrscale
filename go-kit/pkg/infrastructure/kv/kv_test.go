package kv

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	goredis "github.com/redis/go-redis/v9"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv/mocks"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func setupLogger() {
	logger.InitLogger()
}

func TestNewRedis_Success(t *testing.T) {
	setupLogger()
	db, mock := redismock.NewClientMock()
	defer mock.ExpectationsWereMet() // Assert that all expectations were met

	mock.ExpectPing().SetVal("PONG")

	// Test the NewRedis function
	redisStore, err := NewRedis(db, WithGlobalPrefix("test"))
	assert.NotNil(t, redisStore, "Expected redis store to be non-nil")
	assert.Nil(t, err, "Expected error to be nil")

	// Additional checks can be performed here, such as checking the timeout or other configurations
	r, ok := redisStore.(*redis)
	assert.True(t, ok, "Expected IRedisStore to be of type *redis")
	assert.Equal(t, 1*time.Second, r.timeout, "Expected default timeout to be set")
}

func TestNewRedis_Failure(t *testing.T) {
	setupLogger()
	db, mock := redismock.NewClientMock()
	defer mock.ExpectationsWereMet() // Assert that all expectations were met

	mock.ExpectPing().SetErr(goredis.Nil)

	// Test the NewRedis function
	redisStore, err := NewRedis(db, WithGlobalPrefix("test"))
	assert.Nil(t, redisStore, "Expected redis store to be nil due to ping failure")
	assert.NotNil(t, err, "Expected error to be non-nil")
}

func TestRedis_StoreHash(t *testing.T) {
	setupLogger()
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "test:",
	}

	key := "hashKey"
	fields := []interface{}{"field1", "value1", "field2", "value2"}

	// Setup the mock expectation
	mock.ExpectHSet("test:hashKey", fields).SetVal(2) // Assuming 2 fields are set

	// Call the function
	err := r.StoreHash(context.Background(), key, fields...)
	assert.NoError(t, err, "StoreHash should not return an error")

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedis_Ping(t *testing.T) {
	db, mock := redismock.NewClientMock()
	r := &redis{
		client: db,
	}

	mock.ExpectPing().SetVal("PONG")

	err := r.Ping()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedis_Set(t *testing.T) {
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "test:",
	}

	mock.ExpectSet("test:key", "value", 0).SetVal("OK")

	err := r.Set(context.Background(), "key", "value")
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedis_SetWithExpiration(t *testing.T) {
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "test:",
	}

	key := "key"
	fullKey := "test:key" // The key with the global prefix applied
	value := "value"
	expiration := 10 * time.Minute

	// Setup the mock expectation
	mock.ExpectSet(fullKey, value, expiration).SetVal("OK")

	// Call the function
	err := r.SetWithExpiration(context.Background(), key, value, expiration)
	assert.NoError(t, err, "SetWithExpiration should not return an error")

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedis_Get(t *testing.T) {
	setupLogger()
	db, mock := redismock.NewClientMock()
	r := &redis{
		client: db,
	}

	expectedValue := "value"
	mock.ExpectGet("key").SetVal(expectedValue)

	var result string
	err := r.Get(context.Background(), "key", &result)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, result)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedis_ExistHash(t *testing.T) {
	setupLogger()
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "test:",
	}

	key := "hashKey"
	field := "field1"
	fullKey := fmt.Sprintf("test:%s", key)

	// Test case where the hash field exists
	mock.ExpectHExists(fullKey, fmt.Sprintf("%v", field)).SetVal(true)

	exists, err := r.ExistHash(context.Background(), key, field)
	assert.NoError(t, err, "ExistHash should not return an error when the field exists")
	assert.True(t, exists, "Expected the hash field to exist")

	// Test case where the hash field does not exist
	mock.ExpectHExists(fullKey, fmt.Sprintf("%v", field)).SetVal(false)

	exists, err = r.ExistHash(context.Background(), key, field)
	assert.NoError(t, err, "ExistHash should not return an error when the field does not exist")
	assert.False(t, exists, "Expected the hash field to not exist")

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedis_GetClient(t *testing.T) {
	expectedClient := goredis.NewClient(&goredis.Options{})
	expectedPrefix := "testPrefix:"

	r := &redis{
		client:       expectedClient,
		globalPrefix: expectedPrefix,
	}

	prefix, client := r.GetClient()

	assert.Equal(t, expectedPrefix, prefix, "Expected globalPrefix to match")
	assert.Equal(t, expectedClient, client, "Expected *goredis.Client to match")
}

func TestRedis_GetHash(t *testing.T) {
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "test:",
	}

	key := "hashKey"
	fullKey := "test:hashKey"
	filter := &mocks.IFilterStrategy{}

	filter.On("GetFilter").Return(int64(10), "some-pattern*", true)

	// Mocking HScan to return JSON stringified values
	mock.ExpectHScan(fullKey, uint64(0), "some-pattern*", int64(10)).SetVal([]string{
		"key1", `{"include": false, "value": "data1"}`,
	}, uint64(0))

	filter.On("ShouldInclude", []byte(`{"include": false, "value": "data1"}`)).Return(true)

	result, err := r.ScanHash(context.Background(), key, filter)
	assert.NoError(t, err)
	assert.Len(t, result, 1, "Expected one field to be included")
	assert.Equal(t, []byte(`{"include": false, "value": "data1"}`), result[0], "Expected result does not match")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRedis_Remove(t *testing.T) {
	db, mock := redismock.NewClientMock()

	// Assuming `Remove` is a method of a struct that includes a Redis client
	r := &redis{
		client: db,
	}

	// Setup expectations
	keys := []string{"key1", "key2"}
	mock.ExpectDel(keys...).SetVal(int64(len(keys)))

	// Call the method
	err := r.Remove(context.Background(), keys...)
	assert.NoError(t, err)

	// Assert expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %s", err)
	}
}

func TestRedis_Zset(t *testing.T) {
	ctx := context.TODO()
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "prefix_",
	}

	key := "key"
	value := "value"
	score := 1.0
	prefixedKey := "prefix_key"

	mock.ExpectZAdd(prefixedKey, goredis.Z{Score: score, Member: value}).SetVal(1)

	err := r.Zset(ctx, key, value, score)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %s", err)
	}
}

func TestRedis_Zdel(t *testing.T) {
	ctx := context.TODO()
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "prefix_",
	}

	key := "key"
	value := "value"
	prefixedKey := "prefix_key"

	mock.ExpectZRem(prefixedKey, goredis.Z{Member: value}).SetVal(1)

	err := r.Zdel(ctx, key, value)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %s", err)
	}
}

func TestRedis_Zget(t *testing.T) {
	ctx := context.TODO()
	db, mock := redismock.NewClientMock()
	r := &redis{
		client:       db,
		globalPrefix: "prefix_",
	}

	key := "key"
	start := int64(0)
	stop := int64(1)
	prefixedKey := "prefix_key"
	expectedResult := []goredis.Z{
		{Score: 1.0, Member: "value1"},
		{Score: 2.0, Member: "value2"},
	}

	//mock.ExpectZRange(prefixedKey, start, stop).SetVal(expectedResult)
	mock.ExpectZRangeWithScores(prefixedKey, start, stop).SetVal(expectedResult)

	result, err := r.Zget(ctx, key, start, stop)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %s", err)
	}
}

func TestParamsFilter_ShouldInclude(t *testing.T) {
	tests := []struct {
		name           string
		pf             ParamsFilter
		value          string
		expectedResult bool
	}{
		{
			name: "Match criteria met",
			pf: ParamsFilter{
				MatchCriteria: map[string]string{
					"key1": "value1",
				},
				EnablePrefilter: false,
			},
			value:          `{"key1":"value1"}`,
			expectedResult: true,
		},
		{
			name: "Match criteria not met",
			pf: ParamsFilter{
				MatchCriteria: map[string]string{
					"key1": "value1",
				},
				EnablePrefilter: false,
			},
			value:          `{"key1":"value2"}`,
			expectedResult: false,
		},
		{
			name: "Prefilter enabled",
			pf: ParamsFilter{
				EnablePrefilter: true,
			},
			value:          `{}`,
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pf.ShouldInclude([]byte(tt.value)); got != tt.expectedResult {
				t.Errorf("ParamsFilter.ShouldInclude() = %v, want %v", got, tt.expectedResult)
			}
		})
	}
}

func TestParamsFilter_GetFilter(t *testing.T) {
	tests := []struct {
		name              string
		pf                ParamsFilter
		expectedCount     int64
		expectedPattern   string
		expectedPrefilter bool
	}{
		{
			name: "Get filter values",
			pf: ParamsFilter{
				Count:            10,
				PrefilterPattern: "pattern",
				EnablePrefilter:  true,
			},
			expectedCount:     10,
			expectedPattern:   "pattern",
			expectedPrefilter: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotPattern, gotPrefilter := tt.pf.GetFilter()
			if gotCount != tt.expectedCount || gotPattern != tt.expectedPattern || gotPrefilter != tt.expectedPrefilter {
				t.Errorf("ParamsFilter.GetFilter() got = (%v, %v, %v), want (%v, %v, %v)", gotCount, gotPattern, gotPrefilter, tt.expectedCount, tt.expectedPattern, tt.expectedPrefilter)
			}
		})
	}
}

func TestRedis_DeleteHash(t *testing.T) {
	var ctx = context.TODO()
	db, mock := redismock.NewClientMock()

	// Assuming globalPrefix and initialization of redis struct
	r := &redis{
		client:       db,
		globalPrefix: "app:",
	}

	key := "testhash"
	fullKey := "app:testhash"
	field := "field1"

	// Mocking HExists to return true, meaning the field exists
	mock.ExpectHExists(fullKey, field).SetVal(true)

	// Mocking HDel to succeed
	mock.ExpectHDel(fullKey, field).SetVal(1)

	// Test deleting an existing field
	err := r.DeleteHash(ctx, key, field)
	if err != nil {
		t.Errorf("Error while deleting hash: %v", err)
	}

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %s", err)
	}

	// Test for non-existing field
	mock.ExpectHExists(fullKey, field).SetVal(false)

	err = r.DeleteHash(ctx, key, field)
	if err == nil || err.Error() != fmt.Sprintf("key %v not found", field) {
		t.Errorf("Expected error for non-existing field, got %v", err)
	}

	// Test for Redis command error
	mock.ExpectHExists(fullKey, field).SetErr(errors.New("redis error"))

	err = r.DeleteHash(ctx, key, field)
	if err == nil || err.Error() != "redis error" {
		t.Errorf("Expected redis error, got %v", err)
	}
}

package ratelimiter

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisStoreInit(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer rdb.Close()

	redisStore := &redisStore{}
	config := RateLimiterConfig{rdb: rdb}
	redisStore.init("test-key", 100, time.Minute, config)

	assert.Equal(t, int64(100), redisStore.limit)
	assert.Equal(t, REDIS_STORE_PREFIX+"test-key", redisStore.key)
	assert.Equal(t, time.Minute, redisStore.duration)
	assert.Equal(t, rdb, redisStore.rdb)
	assert.NotNil(t, redisStore.ctx)
	mock.ExpectationsWereMet()
}

func TestRedisStoreGetStatusSuccess(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer rdb.Close()

	redisStore := &redisStore{
		limit: 100,
		rdb:   rdb,
		ctx:   context.Background(),
		key:   REDIS_STORE_PREFIX + "test-key",
	}

	mock.ExpectGet(redisStore.key).SetVal("50")

	count, isWithinLimit, err := redisStore.getStatus()
	assert.NoError(t, err)
	assert.Equal(t, int64(50), count)
	assert.True(t, isWithinLimit)
	mock.ExpectationsWereMet()
}

func TestRedisStoreGetStatusError(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer rdb.Close()

	redisStore := &redisStore{
		limit: 100,
		rdb:   rdb,
		ctx:   context.Background(),
		key:   REDIS_STORE_PREFIX + "test-key",
	}

	mock.ExpectGet(redisStore.key).SetErr(redis.Nil)

	_, _, err := redisStore.getStatus()
	assert.Error(t, err)
	mock.ExpectationsWereMet()
}

func TestRedisStoreIncrementAndCheckFirstTime(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer rdb.Close()

	redisStore := &redisStore{
		limit:    100,
		rdb:      rdb,
		ctx:      context.Background(),
		key:      REDIS_STORE_PREFIX + "test-key",
		duration: time.Minute,
	}

	mock.ExpectIncr(redisStore.key).SetVal(1)
	mock.ExpectExpire(redisStore.key, time.Minute).SetVal(true)

	isWithinLimit, err := redisStore.incrementAndCheck()
	assert.NoError(t, err)
	assert.True(t, isWithinLimit)
	mock.ExpectationsWereMet()
}

func TestRedisStoreIncrementAndCheckError(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer rdb.Close()

	redisStore := &redisStore{
		limit: 100,
		rdb:   rdb,
		ctx:   context.Background(),
		key:   REDIS_STORE_PREFIX + "test-key",
	}

	mock.ExpectIncr(redisStore.key).SetErr(redis.Nil)

	isWithinLimit, err := redisStore.incrementAndCheck()
	assert.Error(t, err)
	assert.False(t, isWithinLimit)
	mock.ExpectationsWereMet()
}

func TestRedisStoreIncrementAndCheckOverLimit(t *testing.T) {
	rdb, mock := redismock.NewClientMock()
	defer rdb.Close()

	redisStore := &redisStore{
		limit: 100,
		rdb:   rdb,
		ctx:   context.Background(),
		key:   REDIS_STORE_PREFIX + "test-key",
	}

	mock.ExpectIncr(redisStore.key).SetVal(101)

	isWithinLimit, err := redisStore.incrementAndCheck()
	assert.NoError(t, err)
	assert.False(t, isWithinLimit)
	mock.ExpectationsWereMet()
}

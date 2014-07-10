package cacher

import (
	"log"

	"github.com/alphazero/Go-Redis"
)

func NewRedisCache() *RedisCache {
	spec := redis.DefaultSpec().Db(7)
	client, err := redis.NewSynchClientWithSpec(spec)

	if err != nil {
		log.Println("failed to create the client", err)
	}
	return &RedisCache{
		Client: client,
	}
}

type RedisCache struct {
	Client redis.Client
}

func (rc RedisCache) Get(key string) (*Item, error) {

	value, err := rc.Client.Get(key)
	if err != nil {
		return nil, err
	}

	return &Item{
		Key:   key,
		Value: []byte(value),
	}, nil
}

func (rc RedisCache) Set(key string, value []byte) (err error) {
	return rc.Client.Set(key, value)
}

func (rc RedisCache) Flush() (err error) {
	return rc.Client.Flushdb()
}

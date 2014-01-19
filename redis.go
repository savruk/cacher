package cacher

import (
	"log"

	"github.com/alphazero/Go-Redis"
)

func NewRedisEngine() *RedisEngine {
	spec := redis.DefaultSpec().Db(7)
	client, err := redis.NewSynchClientWithSpec(spec)

	if err != nil {
		log.Println("failed to create the client", err)
	}
	return &RedisEngine{
		Client: client,
	}
}

type RedisEngine struct {
	Client redis.Client
}

func (rc RedisEngine) Get(key string) (*Item, error) {

	value, err := rc.Client.Get(key)
	if err != nil {
		return nil, err
	}

	return &Item{
		Key:   key,
		Value: []byte(value),
	}, nil
}

func (rc RedisEngine) Set(key string, value []byte) (err error) {
	return rc.Client.Set(key, value)
}

func (rc RedisEngine) Flush() (err error) {
	return rc.Client.Flushdb()
}

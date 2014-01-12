package cacher

import (
	"github.com/alphazero/Go-Redis"
	"log"
)

func NewRedisEngine() *RedisEngine {
	spec := redis.DefaultSpec().Db(13)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
	}
	return &RedisEngine{
		Client: client,
	}
}

type RedisEngine struct {
	Client redis.Client
}

func (rc *RedisEngine) Get(key string) (*Item, error) {
	value, err := rc.Client.Get(key)
	if err != nil {
		return nil, err
	}
	return &Item{
		Key:   key,
		Value: []byte(value),
	}, nil
}

func (rc *RedisEngine) Set(key string, value []byte) (err error) {
	return rc.Client.Set(key, value)
}

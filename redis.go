package cacher

import (
	"log"

	"menteslibres.net/gosexy/redis"
)

func NewRedisEngine() *RedisEngine {
	client := redis.New()
	err := client.Connect("127.0.0.1", 6379)

	if err != nil {
		log.Println("failed to create the client", err)
	}
	return &RedisEngine{
		Client: client,
	}
}

type RedisEngine struct {
	Client *redis.Client
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
	_, err = rc.Client.Set(key, value)
	return err
}

func (rc *RedisEngine) Flush() (err error) {
	message, err := rc.Client.FlushDB()
	log.Printf("%s %s", "FLUSHDB", message)
	return err
}

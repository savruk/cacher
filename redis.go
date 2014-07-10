package cacher

import (
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

func NewRedisCache() *RedisCache {
	log.Printf("%s - %s ", os.Getenv("WERCKER_REDIS_HOST"), os.Getenv("WERCKER_REDIS_PORT"))
	client, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Println("failed to create the client", err)
	}
	return &RedisCache{
		Client: client,
	}
}

type RedisCache struct {
	Client redis.Conn
}

func (rc RedisCache) Get(key string) (*Item, error) {

	value, err := redis.String(rc.Client.Do("GET", key))

	if err != nil {
		return nil, err
	}

	return &Item{
		Key:   key,
		Value: []byte(value),
	}, nil
}

func (rc RedisCache) Set(key string, value []byte) (err error) {
	_, err = rc.Client.Do("SET", key, value)
	return
}

func (rc RedisCache) Flush() (err error) {
	return
	// return rc.Client.Flushdb()
}

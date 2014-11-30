package cacher

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

type RedisCache struct {
	Client redis.Conn
}

func NewRedisCache(servers ...Server) *RedisCache {

	client, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", servers[0].Host, servers[0].Port))

	if err != nil {
		log.Println("failed to create the client", err)
	}
	return &RedisCache{
		Client: client,
	}
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
	_, err = rc.Client.Do("FLUSHALL")
	return
}

package cacher

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func NewMemcacheEngine(servers ...Servers) *MemcacheEngine {
	srvs := prepareServers(servers)
	return &MemcacheEngine{
		Client: memcache.New(srvs...),
	}
}

type MemcacheEngine struct {
	Client *memcache.Client
}

func (mc *MemcacheEngine) Get(key string) (*Item, error) {
	item, err := mc.Client.Get(key)
	if err != nil {
		return nil, err
	}
	return &Item{
		Key:        item.Key,
		Value:      item.Value,
		Object:     item.Object,
		Flags:      item.Flags,
		Expiration: item.Expiration,
	}, nil
}

func (mc *MemcacheEngine) Set(key string, value []byte) (err error) {
	return mc.Client.Set(&memcache.Item{Key: key, Value: value})
}

func (mc *MemcacheEngine) Flush() (err error) {
	log.Panicln("Not supported")
	return nil
}

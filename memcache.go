package cacher

import "github.com/bradfitz/gomemcache/memcache"

type Memcache struct {
	Client *memcache.Client
}

func NewMemcache(servers ...Server) *Memcache {
	srvs := prepareServers(servers)
	return &Memcache{
		Client: memcache.New(srvs...),
	}
}

func (mc *Memcache) Get(key string) (*Item, error) {
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

func (mc *Memcache) Set(key string, value []byte) (err error) {
	return mc.Client.Set(&memcache.Item{Key: key, Value: value})
}

func (mc *Memcache) Flush() (err error) {
	return mc.Client.FlushAll()
}

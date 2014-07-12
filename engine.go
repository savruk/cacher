package cacher

import "strings"

type Servers struct {
	Host string
	Port string
}

func prepareServers(servers []Servers) []string {
	var srvs []string
	for _, value := range servers {
		srvs = append(srvs, strings.Join([]string{value.Host, value.Port}, ":"))
	}
	return srvs
}

type Item struct {
	Key        string
	Value      []byte
	Object     interface{}
	Flags      uint32
	Expiration int32
	casid      uint64
}

type Engine interface {
	Get(key string) (*Item, error)
	Set(key string, value []byte) error
	Flush() error
}

package cacher

import (
	"testing"
)

func Test_Set_MemCache(t *testing.T) {
	engine := NewMemcacheEngine(Servers{
		"127.0.0.1", "11211",
	})
	cacher := Cacher{engine}
	err := cacher.Set("hello", []byte("world"))

	if err != nil {
		t.Error(err)
	}
}

func Test_Get_MemCache(t *testing.T) {
	engine := NewMemcacheEngine(Servers{
		"127.0.0.1", "11211",
	})
	cacher := Cacher{engine}
	item, err := cacher.Get("hello")
	if err != nil {
		t.Error(err)
	}

	if string(item.Value) != "world" {
		t.Error("Cache writing failed")
	}

}

func Test_Set_FileCache(t *testing.T) {

	engine := NewFilecacheEngine("__cache__")
	cacher := Cacher{engine}
	err := cacher.Set("hello", []byte("world"))

	if err != nil {
		t.Error(err)
	}

}

func Test_Get_FileCache(t *testing.T) {
	engine := NewFilecacheEngine("__cache__")
	cacher := Cacher{engine}
	item, err := cacher.Get("hello")
	if err != nil {
		t.Error(err)
	}

	if string(item.Value) != "world" {
		t.Error("Cache writing failed")
	}

}

func Test_Set_RedisCache(t *testing.T) {
	engine := NewRedisEngine()
	cacher := Cacher{engine}
	err := cacher.Set("hello", []byte("world"))
	if err != nil {
		t.Error(err)
	}
}

func Test_Get_RedisCache(t *testing.T) {
	engine := NewRedisEngine()
	cacher := Cacher{engine}
	item, err := cacher.Get("hello")
	if err != nil {
		t.Error(err)
	}
}

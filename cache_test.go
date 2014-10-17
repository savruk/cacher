package cacher

import (
	"bytes"
	"os"
	"testing"
)

func TestSetMemCache(t *testing.T) {
	cacher := NewMemcache(Servers{
		"127.0.0.1", "11211",
	})
	err := cacher.Set("hello", []byte("world"))

	if err != nil {
		t.Error(err)
	}
}

func TestGetMemCache(t *testing.T) {
	cacheValue := []byte("world")
	cacher := NewMemcache(Servers{
		"127.0.0.1", "11211",
	})
	item, err := cacher.Get("hello")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(item.Value, cacheValue) {
		t.Error("Cache writing failed")
	}

}

func TestSetFileCache(t *testing.T) {
	cacher := NewFileCache("__cache__")
	err := cacher.Set("hello", []byte("world"))

	if err != nil {
		t.Error(err)
	}

}

func TestGetFileCache(t *testing.T) {
	cacheValue := []byte("world")
	cacher := NewFileCache("__cache__")
	item, err := cacher.Get("hello")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(item.Value, cacheValue) {
		t.Error("Cache writing failed")
	}

}

func TestSetRedisCache(t *testing.T) {
	cacher := NewRedisCache(Servers{
		os.Getenv("WERCKER_REDIS_HOST"), "6379",
	})
	err := cacher.Set("hello", []byte("world"))
	if err != nil {
		t.Error(err)
	}
}

func TestGetRedisCache(t *testing.T) {
	cacheValue := []byte("world")
	cacher := NewRedisCache(Servers{
		os.Getenv("WERCKER_REDIS_HOST"), "6379",
	})
	item, err := cacher.Get("hello")
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(item.Value, cacheValue) {
		t.Error("Cache writing failed")
	}
}

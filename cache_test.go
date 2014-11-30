package cacher

import (
	"bytes"
	"os"
	"testing"
)

var (
	CACHE_KEY   = "hello"
	CACHE_VALUE = []byte("world")
)

func TestSetMemCache(t *testing.T) {
	cacher := NewMemcache(Server{
		"127.0.0.1", "11211",
	})
	err := cacher.Set(CACHE_KEY, CACHE_VALUE)

	if err != nil {
		t.Error(err)
	}
}

func TestGetMemCache(t *testing.T) {
	cacher := NewMemcache(
		Server{"127.0.0.1", "11211"},
		Server{"127.0.0.1", "11212"},
	)
	err := cacher.Set(CACHE_KEY, CACHE_VALUE)

	if err != nil {
		t.Error(err)
	}
	item, err := cacher.Get(CACHE_KEY)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(item.Value, CACHE_VALUE) {
		t.Error("Cache writing failed")
	}
}

func TestFlushMemCache(t *testing.T) {
	cacher := NewMemcache(Server{
		"127.0.0.1", "11211",
	})
	err := cacher.Set(CACHE_KEY, CACHE_VALUE)

	if err != nil {
		t.Error(err)
	}
	if err = cacher.Flush(); err != nil {
		t.Error(err)
	}
	item, err := cacher.Get(CACHE_KEY)
	if item != nil {
		t.Error(err)
	}

}

func TestSetFileCache(t *testing.T) {
	cacher := NewFileCache("__cache__")
	err := cacher.Set(CACHE_KEY, CACHE_VALUE)

	if err != nil {
		t.Error(err)
	}

}

func TestGetFileCache(t *testing.T) {
	cacher := NewFileCache("__cache__")
	item, err := cacher.Get(CACHE_KEY)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(item.Value, CACHE_VALUE) {
		t.Error("Cache writing failed")
	}

}

func TestFlushFileCache(t *testing.T) {
	cacher := NewFileCache("__cache__")
	err := cacher.Set(CACHE_KEY, CACHE_VALUE)
	if err != nil {
		t.Error(err)
	}
	if err = cacher.Flush(); err != nil {
		t.Error(err)
	}
	item, err := cacher.Get(CACHE_KEY)
	if item != nil {
		t.Error("Cache flush failed")
	}
}

func TestSetRedisCache(t *testing.T) {
	cacher := NewRedisCache(Server{
		os.Getenv("WERCKER_REDIS_HOST"), "6379",
	})
	err := cacher.Set(CACHE_KEY, CACHE_VALUE)
	if err != nil {
		t.Error(err)
	}
}

func TestGetRedisCache(t *testing.T) {
	cacher := NewRedisCache(Server{
		os.Getenv("WERCKER_REDIS_HOST"), "6379",
	})
	item, err := cacher.Get(CACHE_KEY)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(item.Value, CACHE_VALUE) {
		t.Error("Cache writing failed")
	}
}

func TestFlushRedisCache(t *testing.T) {
	cacher := NewRedisCache(Server{
		os.Getenv("WERCKER_REDIS_HOST"), "6379",
	})
	err := cacher.Set(CACHE_KEY, CACHE_VALUE)
	if err != nil {
		t.Error(err)
	}
	if err = cacher.Flush(); err != nil {
		t.Error(err)
	}
	item, err := cacher.Get(CACHE_KEY)
	if item != nil {
		t.Error("Cache flush failed")
	}
}

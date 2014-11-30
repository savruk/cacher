package cacher

import (
	"bytes"
	"os"
	"testing"
)

func BenchmarkGetMemCache(b *testing.B) {
	cacheValue := []byte("world")
	for n := 0; n < b.N; n++ {
		cacher := NewMemcache(Server{
			"127.0.0.1", "11211",
		})
		item, err := cacher.Get("hello")
		if err != nil {
			b.Error(err)
		}

		if !bytes.Equal(item.Value, cacheValue) {
			b.Error("Cache writing failed")
		}
	}
}

func BenchmarkGetFileCache(b *testing.B) {
	cacheValue := []byte("world")
	for n := 0; n < b.N; n++ {
		cacher := NewFileCache("__cache__")
		item, err := cacher.Get("hello")
		if err != nil {
			b.Error(err)
		}
		if !bytes.Equal(item.Value, cacheValue) {
			b.Error("Cache writing failed")
		}
	}
}

func BenchmarkGetRedisCache(b *testing.B) {
	cacheValue := []byte("world")
	for n := 0; n < b.N; n++ {
		cacher := NewRedisCache(Server{
			os.Getenv("WERCKER_REDIS_HOST"), "6379",
		})
		item, err := cacher.Get("hello")
		if err != nil {
			b.Error(err)
		}
		if !bytes.Equal(item.Value, cacheValue) {
			b.Error("Cache writing failed")
		}
	}
}

// func BenchmarkSetFileCache(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		cacher := NewFileCache("__cache__")
// 		err := cacher.Set("hello", []byte("world"))

// 		if err != nil {
// 			b.Error(err)
// 		}
// 	}
// }

// func BenchmarkSetRedisCache(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		cacher := NewRedisCache(Servers{
// 			os.Getenv("WERCKER_REDIS_HOST"), "6379",
// 		})
// 		err := cacher.Set("hello", []byte("world"))
// 		if err != nil {
// 			b.Error(err)
// 		}
// 	}
// }

// func BenchmarkSetMemCache(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		cacher := NewMemcache(Servers{
// 			"127.0.0.1", "11211",
// 		})
// 		err := cacher.Set("hello", []byte("world"))

// 		if err != nil {
// 			b.Error(err)
// 		}
// 	}
// }

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

	if err != nil{
		t.Error(err)
	}
}

func Test_Get_MemCache(t *testing.T) {
	engine := NewMemcacheEngine(Servers{
		"127.0.0.1", "11211",
	})
	cacher := Cacher{engine}
	item, err := cacher.Get("hello")
	if err != nil{
		t.Error(err)
	}

	if string(item.Value) != "world" {
		t.Error("Cache writing failed")
	}

}

// func Test_FileCache(t *testing.T) {

// 	m := martini.Classic()

// 	engine := NewFilecacheEngine("__cache__")

// 	m.Use(Caches(engine))

// 	m.Get("/setfilecache", func(cacher Cache) string {
// 		cacher.Set("hello", []byte("world"))
// 		return "OK"
// 	})

// 	m.Get("/getfilecache", func(cacher Cache) string {
// 		item, _ := cacher.Get("hello")
// 		if string(item.Value) != "world" {
// 			t.Error("Cache writing failed")
// 		}
// 		return "OK"
// 	})

// 	res := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/setfilecache", nil)
// 	m.ServeHTTP(res, req)

// 	res2 := httptest.NewRecorder()
// 	req2, _ := http.NewRequest("GET", "/getfilecache", nil)
// 	m.ServeHTTP(res2, req2)
// }

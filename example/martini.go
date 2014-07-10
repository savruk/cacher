package main

import (
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/savruk/cacher"
	"github.com/savruk/cacher/middlewares"
)

func main() {
	m := martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer())
	m.Use(middlewares.MartiniCacheMiddelware(cacher.NewMemcache(cacher.Servers{"127.0.0.1", "11211"})))
	m.Get("/", func(cache cacher.Engine, r render.Render) {
		err := cache.Set("hello", []byte("world"))
		if err != nil {
			log.Println(err)
		}
		item, err := cache.Get("hello")
		if err != nil {
			log.Println(err)
		}
		r.XML(200, string(item.Value))
	})

	m.Run()
}

package middlewares

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/savruk/cacher"
)

func MartiniCacheMiddelware(cache cacher.Engine) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		c.MapTo(cache, (*cacher.Engine)(nil))
	}
}

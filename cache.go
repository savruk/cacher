package cacher

type Cacher struct {
	engine Engine
}

func (c *Cacher) Get(key string) (*Item, error) {
	return c.engine.Get(key)
}

func (c *Cacher) Set(key string, value []byte) (err error) {
	return c.engine.Set(key, value)
}

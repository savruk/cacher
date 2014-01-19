package cacher

type Cacher struct {
	Engine Engine
}

func (c *Cacher) Get(key string) (*Item, error) {
	return c.Engine.Get(key)
}

func (c *Cacher) Set(key string, value []byte) (err error) {
	return c.Engine.Set(key, value)
}

func (c *Cacher) Flush() (err error) {
	return c.Engine.Flush()
}

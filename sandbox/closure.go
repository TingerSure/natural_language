package sandbox

import (
	"errors"
)

type Closure struct {
	value  map[string]Variable
	local  map[string]bool
	cache  []Variable
	parent *Closure
}

func (c *Closure) InitLocal(key string) {
	c.local[key] = true
}

func (c *Closure) GetLocal(key string) (Variable, error) {
	if !c.local[key] {
		return nil, errors.New("Undefined variable.")
	}
	return c.value[key], nil
}

func (c *Closure) SetLocal(key string, value Variable) error {
	if !c.local[key] {
		return errors.New("Undefined variable.")
	}
	c.value[key] = value
	return nil
}

func (c *Closure) GetBubble(key string) (Variable, error) {
	if c.local[key] {
		return c.GetLocal(key)
	}
	if c.parent != nil {
		return c.parent.GetBubble(key)
	}
	return nil, errors.New("Undefined variable.")
}

func (c *Closure) SetBubble(key string, value Variable) error {
	if c.local[key] {
		return c.SetLocal(key, value)
	}
	if c.parent != nil {
		return c.parent.SetBubble(key, value)
	}
	return errors.New("Undefined variable.")
}

func (c *Closure) GetCache(index int) Variable {
	return c.cache[index]
}

func (c *Closure) SetCache(index int, value Variable) error {
	c.cache[index] = value
}

func (c *Closure) ClearCaches() {
	for index := range c.cache {
		c.cache[index] = nil
	}
}

func NewClosure(cacheSize int) *Closure {
	return &Closure{
		cache: make([]Variable, cacheSize),
		local: make(map[string]Variable),
	}
}

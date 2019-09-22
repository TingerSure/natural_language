package sandbox

import (
	"errors"
)

type Closure struct {
	returns map[string]Variable
	value   map[string]Variable
	local   map[string]bool
	cache   map[int]Variable
	parent  *Closure
}

func (c *Closure) SetParent(parent *Closure) {
	c.parent = parent
}

func (c *Closure) SetReturn(key string, value Variable) {
	c.returns[key] = value
}

func (c *Closure) Return() map[string]Variable {
	return c.returns
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

func (c *Closure) SetCache(index int, value Variable) {
	c.cache[index] = value
}

func (c *Closure) Clear() {
	c.cache = nil
	c.returns = nil
}

func NewClosure(parent *Closure) *Closure {
	return &Closure{
		parent:  parent,
		cache:   make(map[int]Variable),
		returns: make(map[string]Variable),
		value:   make(map[string]Variable),
		local:   make(map[string]bool),
	}
}

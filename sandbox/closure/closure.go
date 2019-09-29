package closure

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
)

type Closure struct {
	returns map[string]concept.Variable
	value   map[string]concept.Variable
	local   map[string]bool
	cache   map[int]concept.Variable
	parent  concept.Closure
}

func (c *Closure) SetParent(parent concept.Closure) {
	c.parent = parent
}

func (c *Closure) SetReturn(key string, value concept.Variable) {
	c.returns[key] = value
}

func (c *Closure) MergeReturn(other concept.Closure) {
	for key, value := range other.Return() {
		c.returns[key] = value
	}
}

func (c *Closure) Return() map[string]concept.Variable {
	return c.returns
}

func (c *Closure) InitLocal(key string) {
	c.local[key] = true
}

func (c *Closure) GetLocal(key string) (concept.Variable, concept.Interrupt) {
	if !c.local[key] {
		return nil, interrupt.NewException("none pionter", fmt.Sprintf("Undefined variable: \"%v\".", key))
	}
	return c.value[key], nil
}

func (c *Closure) SetLocal(key string, value concept.Variable) concept.Interrupt {
	if !c.local[key] {
		return interrupt.NewException("none pionter", fmt.Sprintf("Undefined variable: \"%v\".", key))
	}
	c.value[key] = value
	return nil
}

func (c *Closure) GetBubble(key string) (concept.Variable, concept.Interrupt) {
	if c.local[key] {
		return c.GetLocal(key)
	}
	if c.parent != nil {
		return c.parent.GetBubble(key)
	}
	return nil, interrupt.NewException("none pionter", fmt.Sprintf("Undefined variable: \"%v\".", key))
}

func (c *Closure) SetBubble(key string, value concept.Variable) concept.Interrupt {
	if c.local[key] {
		return c.SetLocal(key, value)
	}
	if c.parent != nil {
		return c.parent.SetBubble(key, value)
	}
	return interrupt.NewException("none pionter", fmt.Sprintf("Undefined variable: \"%v\".", key))
}

func (c *Closure) GetCache(index int) concept.Variable {
	return c.cache[index]
}

func (c *Closure) SetCache(index int, value concept.Variable) {
	c.cache[index] = value
}

func (c *Closure) Clear() {
	c.cache = nil
}

func NewClosure(parent concept.Closure) *Closure {
	return &Closure{
		parent:  parent,
		cache:   make(map[int]concept.Variable),
		returns: make(map[string]concept.Variable),
		value:   make(map[string]concept.Variable),
		local:   make(map[string]bool),
	}
}

package closure

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

const (
	historyTypeLocal  = 1
	historyTypeBubble = 2
)

type Closure struct {
	returns   map[string]concept.Variable
	value     map[string]concept.Variable
	local     map[string]bool
	parent    concept.Closure
	history   *History
	extempore *Extempore
}

func (c *Closure) IterateHistory(match func(string, concept.Variable) bool) bool {
	var selectedKey string
	var selectedTypes int
	ok := c.history.Iterate(func(key string, types int) bool {
		var value concept.Variable
		var suspend concept.Interrupt
		switch types {
		case historyTypeLocal:
			value, suspend = c.GetLocal(key)
			if nl_interface.IsNil(suspend) {
				return false
			}
		case historyTypeBubble:
			value, suspend = c.GetBubble(key)
			if nl_interface.IsNil(suspend) {
				return false
			}
		}
		selectedKey = key
		selectedTypes = types
		return match(key, value)
	})
	if ok {
		c.history.Set(selectedKey, selectedTypes)
	}
	return ok
}

func (c *Closure) IterateExtempore(match func(concept.Index, concept.Variable) bool) bool {
	return c.extempore.Iterate(match)
}

func (c *Closure) IterateLocal(match func(string, concept.Variable) bool) bool {
	for key, value := range c.value {
		if match(key, value) {
			c.history.Set(key, historyTypeLocal)
			return true
		}
	}
	return false
}

func (c *Closure) IterateBubble(match func(string, concept.Variable) bool) bool {
	if c.IterateLocal(match) {
		return true
	}
	if c.parent == nil {
		return false
	}
	var selectedKey string
	ok := c.parent.IterateBubble(func(key string, value concept.Variable) bool {
		selectedKey = key
		return match(key, value)
	})
	if ok {
		c.history.Set(selectedKey, historyTypeBubble)
	}
	return ok
}

func (c *Closure) SetParent(parent concept.Closure) {
	c.parent = parent
}

func (c *Closure) AddExtempore(index concept.Index, value concept.Variable) {
	c.extempore.Add(index, value)
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
	c.history.Set(key, historyTypeLocal)
	return c.value[key], nil
}

func (c *Closure) SetLocal(key string, value concept.Variable) concept.Interrupt {
	if !c.local[key] {
		return interrupt.NewException("none pionter", fmt.Sprintf("Undefined variable: \"%v\".", key))
	}
	c.history.Set(key, historyTypeLocal)
	c.value[key] = value
	return nil
}

func (c *Closure) GetBubble(key string) (concept.Variable, concept.Interrupt) {
	if c.local[key] {
		return c.GetLocal(key)
	}
	if c.parent != nil {
		value, suspend := c.parent.GetBubble(key)
		if nl_interface.IsNil(suspend) {
			c.history.Set(key, historyTypeBubble)
		}
		return value, suspend
	}
	return nil, interrupt.NewException("none pionter", fmt.Sprintf("Undefined variable: \"%v\".", key))
}

func (c *Closure) SetBubble(key string, value concept.Variable) concept.Interrupt {
	if c.local[key] {
		return c.SetLocal(key, value)
	}
	if c.parent != nil {
		suspend := c.parent.SetBubble(key, value)
		if nl_interface.IsNil(suspend) {
			c.history.Set(key, historyTypeBubble)
		}
		return suspend
	}
	return interrupt.NewException("none pionter", fmt.Sprintf("Undefined variable: \"%v\".", key))
}

func (c *Closure) Clear() {
}

func NewClosure(parent concept.Closure) *Closure {
	return &Closure{
		parent:    parent,
		returns:   make(map[string]concept.Variable),
		value:     make(map[string]concept.Variable),
		local:     make(map[string]bool),
		history:   NewHistory(),
		extempore: NewExtempore(),
	}
}

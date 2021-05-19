package closure

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	historyTypeLocal  = 1
	historyTypeBubble = 2
)

type ClosureSeed interface {
	NewException(string, string) concept.Exception
	NewNull() concept.Null
}

type Closure struct {
	local     *concept.Mapping
	parent    concept.Closure
	history   *History
	extempore *Extempore
	seed      ClosureSeed
}

func (c *Closure) IterateHistory(match func(concept.String, concept.Variable) bool) bool {
	return c.history.Iterate(func(key concept.String, types int) bool {
		var value concept.Variable
		var suspend concept.Exception
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
		if match(key, value) {
			c.history.Set(key, types)
			return true
		}
		return false
	})
}

func (c *Closure) IterateExtempore(match func(concept.Index, concept.Variable) bool) bool {
	return c.extempore.Iterate(match)
}

func (c *Closure) IterateLocal(match func(concept.String, concept.Variable) bool) bool {
	return c.local.Iterate(func(key concept.String, value interface{}) bool {
		if match(key, value.(concept.Variable)) {
			c.history.Set(key, historyTypeLocal)
			return true
		}
		return false
	})
}

func (c *Closure) IterateBubble(match func(concept.String, concept.Variable) bool) bool {
	if c.IterateLocal(match) {
		return true
	}
	if nl_interface.IsNil(c.parent) {
		return false
	}
	return c.parent.IterateBubble(func(key concept.String, value concept.Variable) bool {
		if match(key, value.(concept.Variable)) {
			c.history.Set(key, historyTypeBubble)
			return true
		}
		return false
	})
}

func (c *Closure) SetParent(parent concept.Closure) {
	c.parent = parent
}

func (c *Closure) AddExtempore(index concept.Index, value concept.Variable) {
	c.extempore.Add(index, value)
}

func (c *Closure) InitLocal(key concept.String, defaultValue concept.Variable) {
	c.local.Init(key, defaultValue)
}

func (c *Closure) PeekLocal(key concept.String) (concept.Variable, concept.Exception) {
	if !c.local.Has(key) {
		return nil, c.seed.NewException("none pionter", fmt.Sprintf("Undefined variable: %v.", key.ToString("")))
	}
	return c.local.Get(key).(concept.Variable), nil
}

func (c *Closure) HasLocal(key concept.String) bool {
	return c.local.Has(key)
}

func (c *Closure) GetLocal(key concept.String) (concept.Variable, concept.Exception) {
	if !c.local.Has(key) {
		return nil, c.seed.NewException("none pionter", fmt.Sprintf("Undefined variable: %v.", key.ToString("")))
	}
	c.history.Set(key, historyTypeLocal)
	return c.local.Get(key).(concept.Variable), nil
}

func (c *Closure) SetLocal(key concept.String, value concept.Variable) concept.Exception {
	if !c.local.Set(key, value) {
		return c.seed.NewException("none pionter", fmt.Sprintf("Undefined variable: %v.", key.ToString("")))
	}
	c.history.Set(key, historyTypeLocal)
	return nil
}

func (c *Closure) HasBubble(key concept.String) bool {
	yes := c.HasLocal(key)
	if yes {
		return true
	}
	if c.parent != nil {
		return c.parent.HasBubble(key)
	}
	return false
}

func (c *Closure) PeekBubble(key concept.String) (concept.Variable, concept.Exception) {
	value, suspend := c.PeekLocal(key)
	if nl_interface.IsNil(suspend) {
		return value, nil
	}
	if c.parent != nil {
		return c.parent.PeekBubble(key)
	}
	return c.seed.NewNull(), c.seed.NewException("none pionter", fmt.Sprintf("Undefined variable: %v.", key.ToString("")))
}

func (c *Closure) GetBubble(key concept.String) (concept.Variable, concept.Exception) {
	value, suspend := c.GetLocal(key)
	if nl_interface.IsNil(suspend) {
		c.history.Set(key, historyTypeBubble)
		return value, nil
	}
	if c.parent != nil {
		value, suspend = c.parent.GetBubble(key)
		if nl_interface.IsNil(suspend) {
			c.history.Set(key, historyTypeBubble)
		}
		return value, suspend
	}
	return c.seed.NewNull(), c.seed.NewException("none pionter", fmt.Sprintf("Undefined variable: %v.", key.ToString("")))
}

func (c *Closure) SetBubble(key concept.String, value concept.Variable) concept.Exception {
	suspend := c.SetLocal(key, value)
	if nl_interface.IsNil(suspend) {
		return nil
	}
	if c.parent != nil {
		suspend := c.parent.SetBubble(key, value)
		if nl_interface.IsNil(suspend) {
			c.history.Set(key, historyTypeBubble)
		}
		return suspend
	}
	return c.seed.NewException("none pionter", fmt.Sprintf("Undefined variable: %v.", key.ToString("")))
}

func (c *Closure) Clear() {
}

type ClosureCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	EmptyCreator     func() concept.Null
}

type ClosureCreator struct {
	param *ClosureCreatorParam
}

func (s *ClosureCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ClosureCreator) NewNull() concept.Null {
	return s.param.EmptyCreator()
}

func (s *ClosureCreator) New(parent concept.Closure) *Closure {
	return &Closure{
		parent:    parent,
		history:   NewHistory(),
		extempore: NewExtempore(),
		local: concept.NewMapping(&concept.MappingParam{
			AutoInit:   false,
			EmptyValue: s.NewNull(),
		}),
		seed: s,
	}
}

func NewClosureCreator(param *ClosureCreatorParam) *ClosureCreator {
	return &ClosureCreator{
		param: param,
	}
}

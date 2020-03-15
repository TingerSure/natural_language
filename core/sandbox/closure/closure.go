package closure

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/component"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

const (
	historyTypeLocal  = 1
	historyTypeBubble = 2
)

type Closure struct {
	param     *ClosureParam
	returns   *component.Mapping //map[string]concept.Variable
	local     *component.Mapping //map[string]concept.Variable
	parent    concept.Closure
	history   *History
	extempore *Extempore
}

func (c *Closure) IterateHistory(match func(concept.String, concept.Variable) bool) bool {
	return c.history.Iterate(func(key concept.String, types int) bool {
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

func (c *Closure) SetReturn(key concept.String, value concept.Variable) {
	c.returns.Set(key, value)
}

func (c *Closure) MergeReturn(other concept.Closure) {
	other.IterateReturn(func(key concept.String, value concept.Variable) bool {
		c.returns.Set(key, value)
		return false
	})
}

func (c *Closure) IterateReturn(on func(key concept.String, value concept.Variable) bool) bool {
	return c.returns.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Variable))
	})
}

func (c *Closure) InitLocal(key concept.String, defaultValue concept.Variable) {
	c.local.Init(key, defaultValue)
}

func (c *Closure) GetLocal(key concept.String) (concept.Variable, concept.Interrupt) {
	value := c.local.Get(key)
	if nl_interface.IsNil(value) {
		return nil, interrupt.NewException(c.param.StringCreator("none pionter"), c.param.StringCreator(fmt.Sprintf("Undefined variable: \"%v\".", key)))
	}
	c.history.Set(key, historyTypeLocal)
	return value.(concept.Variable), nil
}

func (c *Closure) SetLocal(key concept.String, value concept.Variable) concept.Interrupt {
	if !c.local.Set(key, value) {
		return interrupt.NewException(c.param.StringCreator("none pionter"), c.param.StringCreator(fmt.Sprintf("Undefined variable: \"%v\".", key)))
	}
	c.history.Set(key, historyTypeLocal)
	return nil
}

func (c *Closure) GetBubble(key concept.String) (concept.Variable, concept.Interrupt) {
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
	return nil, interrupt.NewException(c.param.StringCreator("none pionter"), c.param.StringCreator(fmt.Sprintf("Undefined variable: \"%v\".", key)))
}

func (c *Closure) SetBubble(key concept.String, value concept.Variable) concept.Interrupt {
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
	return interrupt.NewException(c.param.StringCreator("none pionter"), c.param.StringCreator(fmt.Sprintf("Undefined variable: \"%v\".", key)))
}

func (c *Closure) Clear() {
}

type ClosureParam struct {
	StringCreator func(string) concept.String
}

func NewClosure(parent concept.Closure, param *ClosureParam) *Closure {
	return &Closure{
		parent:    parent,
		param:     param,
		history:   NewHistory(),
		extempore: NewExtempore(),
		returns: component.NewMapping(&component.MappingParam{
			AutoInit: true,
		}),
		local: component.NewMapping(&component.MappingParam{
			AutoInit: false,
		}),
	}
}

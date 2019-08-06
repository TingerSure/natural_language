package grammar

import (
	"github.com/TingerSure/natural_language/tree"
)

type Collector struct {
	phrases []tree.Phrase
	length  int
}

func (c *Collector) Copy() *Collector {
	substitute := NewCollector()
	for _, treasure := range c.phrases {
		substitute.Push(treasure.Copy())
	}
	return substitute
}

func (c *Collector) Push(treasure tree.Phrase) {
	c.phrases = append(c.phrases, treasure)
	c.length++
}

func (c *Collector) Pop() tree.Phrase {
	if c.length <= 0 {
		return nil
	}
	c.length--
	peek := c.phrases[c.length]
	c.phrases = c.phrases[:c.length]
	return peek
}
func (c *Collector) PopMultiple(size int) []tree.Phrase {
	if c.length < size {
		return nil
	}
	c.length -= size
	peek := c.phrases[c.length:]
	c.phrases = c.phrases[:c.length]
	return peek
}
func (c *Collector) Peek() tree.Phrase {
	if c.length <= 0 {
		return nil
	}
	return c.phrases[c.length-1]
}
func (c *Collector) PeekMultiple(size int) []tree.Phrase {
	if c.length < size {
		return nil
	}
	return c.phrases[c.length-size:]
}

func (c *Collector) Len() int {
	return c.length
}
func (c *Collector) IsEmpty() bool {
	return c.length == 0
}
func (c *Collector) IsSingle() bool {
	return c.length == 1
}

func (c *Collector) init() *Collector {
	return c
}

func NewCollector() *Collector {
	return (&Collector{}).init()
}

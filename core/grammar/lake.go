package grammar

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Lake struct {
	phrases []tree.Phrase
	length  int
}

func (c *Lake) Copy() *Lake {
	substitute := NewLake()
	for _, treasure := range c.phrases {
		substitute.Push(treasure.Copy())
	}
	return substitute
}

func (c *Lake) Push(treasure tree.Phrase) {
	c.phrases = append(c.phrases, treasure)
	c.length++
}

func (c *Lake) Pop() tree.Phrase {
	if c.length <= 0 {
		return nil
	}
	c.length--
	peek := c.phrases[c.length]
	c.phrases = c.phrases[:c.length]
	return peek
}
func (c *Lake) PopMultiple(size int) []tree.Phrase {
	if c.length < size {
		return nil
	}
	c.length -= size
	peek := c.phrases[c.length:]
	c.phrases = c.phrases[:c.length]
	return peek
}
func (c *Lake) Peek() tree.Phrase {
	if c.length <= 0 {
		return nil
	}
	return c.phrases[c.length-1]
}

func (c *Lake) PeekAll() []tree.Phrase {
	return c.phrases
}

func (c *Lake) PeekMultiple(size int) []tree.Phrase {
	if c.length < size {
		return nil
	}
	return c.phrases[c.length-size:]
}

func (c *Lake) Len() int {
	return c.length
}
func (c *Lake) IsEmpty() bool {
	return c.length == 0
}
func (c *Lake) IsSingle() bool {
	return c.length == 1
}

func (c *Lake) init() *Lake {
	return c
}

func NewLake() *Lake {
	return (&Lake{}).init()
}

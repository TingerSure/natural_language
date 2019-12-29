package closure

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type extemporeNode struct {
	value concept.Variable
	line  concept.Index
	next  *extemporeNode
}

type Extempore struct {
	root *extemporeNode
}

func (c *Extempore) Clear() {
	c.root = nil
}

func (c *Extempore) Iterate(match func(concept.Index, concept.Variable) bool) bool {
	for cursor := c.root; cursor != nil; cursor = cursor.next {
		if match(cursor.line, cursor.value) {
			return true
		}
	}
	return false
}

func (c *Extempore) Add(line concept.Index, value concept.Variable) {
	c.root = &extemporeNode{
		value: value,
		line:  line,
		next:  c.root,
	}
}

func NewExtempore() *Extempore {
	return &Extempore{}
}

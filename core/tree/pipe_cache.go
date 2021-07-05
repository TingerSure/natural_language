package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type pipeCacheNode struct {
	value concept.Variable
	pipe  concept.Function
	next  *pipeCacheNode
}

type PipeCache struct {
	root *pipeCacheNode
}

func (c *PipeCache) Clear() {
	c.root = nil
}

func (c *PipeCache) Iterate(match func(concept.Function, concept.Variable) bool) bool {
	for cursor := c.root; cursor != nil; cursor = cursor.next {
		if match(cursor.pipe, cursor.value) {
			return true
		}
	}
	return false
}

func (c *PipeCache) Add(pipe concept.Function, value concept.Variable) {
	c.root = &pipeCacheNode{
		pipe:  pipe,
		value: value,
		next:  c.root,
	}
}

type PipeCacheParam struct {
	//TODO
	Size    int
	TimeOut int64
}

func NewPipeCache(param *PipeCacheParam) *PipeCache {
	return &PipeCache{}
}

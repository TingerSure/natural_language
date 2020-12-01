package parser

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type indexListNode struct {
	next  *indexListNode
	value tree.Phrase
}

type IndexList struct {
	size   int
	values []*indexListNode
}

func NewIndexList(size int) *IndexList {
	return &IndexList{
		size:   size,
		values: make([]*indexListNode, size),
	}
}

func (r *IndexList) Iterate(onPhrase func(tree.Phrase) bool) bool {
	for _, node := range r.values {
		for ; node != nil; node = node.next {
			if onPhrase(node.value) {
				return true
			}
		}
	}
	return false
}

func (r *IndexList) GetBySize(index int, size int) []tree.Phrase {
	back := []tree.Phrase{}
	for cursor := r.values[index]; cursor != nil; cursor = cursor.next {
		if cursor.value.ContentSize() == size {
			back = append(back, cursor.value)
		}
	}
	return back
}

func (r *IndexList) GetByTypesAndSize(index int, types string, size int) []tree.Phrase {
	values := []tree.Phrase{}
	for cursor := r.values[index]; cursor != nil; cursor = cursor.next {
		if cursor.value.ContentSize() > size {
			continue
		}
		if cursor.value.ContentSize() < size {
			break
		}
		if types == cursor.value.Types() {
			values = append(values, cursor.value)
		}
	}
	return values
}

func (r *IndexList) GetMaxBySize(index int) tree.Phrase {
	if r.values[index] == nil {
		return nil
	}
	return r.values[index].value
}

func (r *IndexList) GetAll(index int) []tree.Phrase {
	back := []tree.Phrase{}
	for cursor := r.values[index]; cursor != nil; cursor = cursor.next {
		back = append(back, cursor.value)
	}
	return back
}

func (r *IndexList) Has(index int) bool {
	return r.values[index] != nil
}

func (l *IndexList) Add(index int, value tree.Phrase) {
	l.values[index] = l.addNode(l.values[index], value)
}

func (r *IndexList) addNode(root *indexListNode, value tree.Phrase) *indexListNode {
	if root == nil || root.value.ContentSize() <= value.ContentSize() {
		return &indexListNode{
			value: value,
			next:  root,
		}
	}
	last := root
	cursor := root.next
	for cursor != nil && cursor.value.ContentSize() > value.ContentSize() {
		last = cursor
		cursor = cursor.next
	}
	last.next = &indexListNode{
		value: value,
		next:  cursor,
	}
	return root
}

func (r *IndexList) Remove(index int, value tree.Phrase) {
	r.values[index] = r.removeNode(r.values[index], value)
}

func (r *IndexList) removeNode(root *indexListNode, value tree.Phrase) *indexListNode {
	cursor := root
	var last *indexListNode = nil
	for cursor != nil {
		if value == cursor.value {
			if cursor == root {
				root = cursor.next
			} else {
				last.next = cursor.next
			}
		} else {
			last = cursor
		}
		cursor = cursor.next
	}
	return root
}

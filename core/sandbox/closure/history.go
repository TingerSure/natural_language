package closure

type historyNode struct {
	key   string
	types int
	next  *historyNode
}

type History struct {
	root *historyNode
}

func (c *History) Clear() {
	c.root = nil
}

func (c *History) Iterate(match func(string, int) bool) bool {
	for cursor := c.root; cursor != nil; cursor = cursor.next {
		if match(cursor.key, cursor.types) {
			return true
		}
	}
	return false
}

func (c *History) Set(key string, types int) {
	var last *historyNode
	var hit *historyNode
	for cursor := c.root; cursor != nil; cursor = cursor.next {
		if cursor.key == key && cursor.types == types {
			hit = cursor
			break
		}
		last = cursor
	}
	if last == nil || hit == nil {
		c.root = &historyNode{
			key:   key,
			types: types,
			next:  c.root,
		}
		return
	}
	last.next = hit.next
	hit.next = c.root
	c.root = hit
	hit.key = key
	hit.types = types
}

func NewHistory() *History {
	return &History{}
}

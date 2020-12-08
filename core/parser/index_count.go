package parser

type IndexCount struct {
	size   int
	values []int
}

func NewIndexCount(size int) *IndexCount {
	return &IndexCount{
		size:   size,
		values: make([]int, size),
	}
}

func (c *IndexCount) Get(index int) int {
	return c.values[index]
}

func (c *IndexCount) Add(index int) {
	c.values[index]++
}

func (c *IndexCount) Remove(index int) {
	c.values[index]--
}

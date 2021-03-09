package grammar

type Count struct {
	value int
}

func NewCount(start int) *Count {
	return &Count{
		value: start,
	}
}

func (c *Count) Next() (value int) {
	value = c.value
	c.value++
	return
}

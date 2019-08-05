package tree

type StructRule struct {
	size   int
	create func() Phrase
	from   string
	types  []string
}

func (r *StructRule) Size() int {
	return r.size
}

func (r *StructRule) GetFrom() string {
	return r.from
}

func (r *StructRule) Logic(treasures []Phrase) Phrase {
	if len(treasures) < r.size {
		return nil
	}
	active := treasures[len(treasures)-r.size:]

	match := true
	for index, treasure := range active {
		match = match && treasure.Types() == r.types[index]
	}
	if !match {
		return nil
	}
	new := r.create()
	for index, treasure := range active {
		new.SetChild(index, treasure)
	}
	return new
}

func NewStructRule(
	create func() Phrase,
	types []string,
	from string,
) *StructRule {
	size := len(types)
	if size == 0 {
		panic("There must be at least one type here.")
	}
	if create == nil {
		panic("no create function in this struct rule!")
	}
	return &StructRule{
		types:  types,
		create: create,
		size:   size,
		from:   from,
	}
}

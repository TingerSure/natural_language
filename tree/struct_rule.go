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

func (r *StructRule) Match(treasures []Phrase) bool {
	if len(treasures) < r.size {
		return false
	}
	match := true
	for index, treasure := range treasures[len(treasures)-r.size:] {
		match = match && treasure.Types() == r.types[index]
	}
	return match
}

func (r *StructRule) Create(treasures []Phrase) Phrase {
	new := r.create()
	for index, treasure := range treasures[len(treasures)-r.size:] {
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

package tree

type StructRule struct {
	size  int
	logic func(treasures []*Phrase) *Phrase
	from  string
}

func (r *StructRule) Size() int {
	return r.size
}

func (r *StructRule) GetFrom() string {
	return r.from
}

func (r *StructRule) Logic(treasures []*Phrase) *Phrase {
	if len(treasures) < r.size {
		return nil
	}
	mixture := r.logic(treasures[len(treasures)-r.size:])
	if mixture == nil {
		return nil
	}
	return mixture
}

func NewStructRule(
	logic func(treasures []*Phrase) *Phrase,
	size int,
	from string,
) *StructRule {
	if logic == nil {
		panic("no logic function in this struct rule!")
	}
	return &StructRule{
		logic: logic,
		size:  size,
		from:  from,
	}
}

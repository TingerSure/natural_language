package grammar

import ()

type Rule struct {
	result Symbol
	from   []Symbol
}

func NewRule(result Symbol, from ...Symbol) *Rule {
	return &Rule{
		result: result,
		from:   from,
	}
}

func (r *Rule) SetResult(result Symbol) {
	r.result = result
}

func (r *Rule) AppendFrom(from Symbol) {
	r.from = append(r.from, from)
}

func (r *Rule) GetResult() Symbol {
	return r.result
}

func (r *Rule) GetFrom(index int) Symbol {
	return r.from[index]
}

func (r *Rule) Size() int {
	return len(r.from)
}

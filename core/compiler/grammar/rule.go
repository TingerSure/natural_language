package grammar

import ()

type Rule struct {
	result   Symbol
	children []Symbol
}

func NewRule(result Symbol, children ...Symbol) *Rule {
	return &Rule{
		result:   result,
		children: children,
	}
}

func (r *Rule) SetResult(result Symbol) {
	r.result = result
}

func (r *Rule) AppendChild(child Symbol) {
	r.children = append(r.children, child)
}

func (r *Rule) GetResult() Symbol {
	return r.result
}

func (r *Rule) GetChild(index int) Symbol {
	return r.children[index]
}

func (r *Rule) Size() int {
	return len(r.children)
}

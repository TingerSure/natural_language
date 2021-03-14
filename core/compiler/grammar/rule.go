package grammar

import (
	"fmt"
	"strings"
)

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

func (r *Rule) ToString() string {
	names := []string{}
	for _, child := range r.children {
		names = append(names, child.Name())
	}
	return fmt.Sprintf("%v -> %v", r.result.Name(), strings.Join(names, " "))
}

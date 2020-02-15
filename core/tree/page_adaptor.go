package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type PageAdaptor struct {
	functions map[string]concept.Function
	classes   map[string]concept.Class
	consts    map[string]string
}

func (p *PageAdaptor) GetFunction(key string) concept.Function {
	return p.functions[key]
}

func (p *PageAdaptor) SetFunction(key string, value concept.Function) Page {
	p.functions[key] = value
	return p
}

func (p *PageAdaptor) GetClass(key string) concept.Class {
	return p.classes[key]
}

func (p *PageAdaptor) SetClass(key string, value concept.Class) Page {
	p.classes[key] = value
	return p
}

func (p *PageAdaptor) GetConst(key string) string {
	return p.consts[key]
}

func (p *PageAdaptor) SetConst(key string, value string) Page {
	p.consts[key] = value
	return p
}

func NewPageAdaptor() *PageAdaptor {
	return &PageAdaptor{
		functions: map[string]concept.Function{},
		classes:   map[string]concept.Class{},
		consts:    map[string]string{},
	}
}

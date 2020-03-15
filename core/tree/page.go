package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/component"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Page struct {
	functions *component.Mapping
	classes   *component.Mapping
	consts    *component.Mapping
}

func (p *Page) GetFunction(key concept.String) concept.Function {
	return p.functions.Get(key).(concept.Function)
}

func (p *Page) SetFunction(key concept.String, value concept.Function) *Page {
	p.functions.Set(key, value)
	return p
}

func (p *Page) GetClass(key concept.String) concept.Class {
	return p.classes.Get(key).(concept.Class)
}

func (p *Page) SetClass(key concept.String, value concept.Class) *Page {
	p.classes.Set(key, value)
	return p
}

func (p *Page) GetConst(key concept.String) concept.String {
	return p.consts.Get(key).(concept.String)
}

func (p *Page) SetConst(key concept.String, value concept.String) *Page {
	p.consts.Set(key, value)
	return p
}

func NewPage() *Page {
	return &Page{
		functions: component.NewMapping(&component.MappingParam{
			AutoInit: true,
		}),
		classes: component.NewMapping(&component.MappingParam{
			AutoInit: true,
		}),
		consts: component.NewMapping(&component.MappingParam{
			AutoInit: true,
		}),
	}
}

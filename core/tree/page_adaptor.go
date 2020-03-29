package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/component"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type PageAdaptor struct {
	functions *component.Mapping
	classes   *component.Mapping
	consts    *component.Mapping
}

func (p *PageAdaptor) GetFunction(key concept.String) concept.Function {
	return p.functions.Get(key).(concept.Function)
}

func (p *PageAdaptor) SetFunction(key concept.String, value concept.Function) Page {
	p.functions.Set(key, value)
	return p
}

func (p *PageAdaptor) GetClass(key concept.String) concept.Class {
	return p.classes.Get(key).(concept.Class)
}

func (p *PageAdaptor) SetClass(key concept.String, value concept.Class) Page {
	p.classes.Set(key, value)
	return p
}

func (p *PageAdaptor) GetConst(key concept.String) concept.String {
	return p.consts.Get(key).(concept.String)
}

func (p *PageAdaptor) SetConst(key concept.String, value concept.String) Page {
	p.consts.Set(key, value)
	return p
}

func NewPageAdaptor() *PageAdaptor {
	return &PageAdaptor{
		functions: component.NewMapping(&component.MappingParam{
			AutoInit:   true,
			EmptyValue: variable.NewNull(),
		}),
		classes: component.NewMapping(&component.MappingParam{
			AutoInit:   true,
			EmptyValue: variable.NewNull(),
		}),
		consts: component.NewMapping(&component.MappingParam{
			AutoInit:   true,
			EmptyValue: variable.NewNull(),
		}),
	}
}

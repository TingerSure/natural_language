package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/component"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type PageAdaptor struct {
	functions  *component.Mapping
	classes    *component.Mapping
	consts     *component.Mapping
	exceptions *component.Mapping
	sources    []Source
}

func (p *PageAdaptor) GetException(key concept.String) concept.Exception {
	return p.exceptions.Get(key).(concept.Exception)
}

func (p *PageAdaptor) SetException(key concept.String, value concept.Exception) Page {
	p.exceptions.Set(key, value)
	return p
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

func (p *PageAdaptor) GetSources() []Source {
	return p.sources
}

func (p *PageAdaptor) AddSource(source Source) {
	p.sources = append(p.sources, source)
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
		sources: []Source{},
	}
}

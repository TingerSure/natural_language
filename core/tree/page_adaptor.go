package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/creator"
)

type PageAdaptor struct {
	functions  *concept.Mapping
	classes    *concept.Mapping
	consts     *concept.Mapping
	exceptions *concept.Mapping
	indexes    *concept.Mapping
}

const (
	IndexPageType = "Page"
)

func (p *PageAdaptor) Type() string {
	return IndexPageType
}

func (p *PageAdaptor) Anticipate(concept.Closure) concept.Variable {
	return nil
}

func (p *PageAdaptor) Get(concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, nil
}

func (p *PageAdaptor) Set(concept.Closure, concept.Variable) concept.Interrupt {
	return nil
}

func (p *PageAdaptor) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (p *PageAdaptor) ToString(prefix string) string {
	// TODO
	return "class.ToString()"
}

func (p *PageAdaptor) ToLanguage(language string) string {
	return p.ToString("")
}

func (p *PageAdaptor) GetIndex(key concept.String) concept.Index {
	return p.indexes.Get(key).(concept.Index)
}

func (p *PageAdaptor) SetIndex(key concept.String, value concept.Index) Page {
	p.indexes.Set(key, value)
	return p
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

func NewPageAdaptor(sandboxSeed *creator.SandboxCreator) *PageAdaptor {
	return &PageAdaptor{
		indexes: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: sandboxSeed.Variable.Null.New(),
		}),
		exceptions: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: sandboxSeed.Variable.Null.New(),
		}),
		functions: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: sandboxSeed.Variable.Null.New(),
		}),
		classes: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: sandboxSeed.Variable.Null.New(),
		}),
		consts: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: sandboxSeed.Variable.Null.New(),
		}),
	}
}

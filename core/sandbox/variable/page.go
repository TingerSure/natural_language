package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"strings"
)

const (
	VariablePageType = "object"
)

type PageSeed interface {
	NewNull() concept.Null
	NewException(string, string) concept.Exception
	ToLanguage(string, *Page) string
	Type() string
}

type Page struct {
	seed    PageSeed
	imports *concept.Mapping
	exports *concept.Mapping
	vars    *concept.Mapping
	space   concept.Closure
}

func (o *Page) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	value, exception := o.GetField(specimen)
	if nl_interface.IsNil(exception) {
		return nil, exception
	}
	if !value.IsFunction() {
		return nil, o.seed.NewException("runtime error", fmt.Sprintf("There is no export function called %v to be called here.", specimen.ToString("")))
	}
	return value.(concept.Function).Exec(param, nil)
}

func (o *Page) SetImport(specimen concept.String, indexes concept.Index) concept.Exception {
	if o.space.HasLocal(specimen) {
		return o.seed.NewException("runtime error", fmt.Sprintf("Import %v already exists.", specimen.ToString("")))
	}
	value, exception := indexes.Get(nil)
	if !nl_interface.IsNil(exception) {
		return exception.(concept.Exception)
	}
	o.imports.Set(specimen, indexes)
	o.space.InitLocal(specimen, value)
	return nil
}

func (o *Page) SetExport(specimen concept.String, indexes concept.Index) concept.Exception {
	if o.space.HasLocal(specimen) {
		return o.seed.NewException("runtime error", fmt.Sprintf("Export %v already exists.", specimen.ToString("")))
	}
	value, exception := indexes.Get(nil)
	if !nl_interface.IsNil(exception) {
		return exception.(concept.Exception)
	}
	o.exports.Set(specimen, indexes)
	o.space.InitLocal(specimen, value)
	return nil
}

func (o *Page) SetVar(specimen concept.String, indexes concept.Index) concept.Exception {
	if o.space.HasLocal(specimen) {
		return o.seed.NewException("runtime error", fmt.Sprintf("Var %v already exists.", specimen.ToString("")))
	}
	value, exception := indexes.Get(o.space)
	if !nl_interface.IsNil(exception) {
		return exception.(concept.Exception)
	}
	o.vars.Set(specimen, indexes)
	o.space.InitLocal(specimen, value)
	return nil
}

func (o *Page) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	if !o.exports.Has(specimen) {
		return o.seed.NewException("runtime error", fmt.Sprintf("There is no export field called %v to be set here.", specimen.ToString("")))
	}
	return o.space.SetLocal(specimen, value)
}

func (o *Page) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	if !o.exports.Has(specimen) {
		return o.seed.NewNull(), nil
	}
	return o.space.GetLocal(specimen)
}

func (o *Page) HasField(specimen concept.String) bool {
	return o.exports.Has(specimen)
}

func (o *Page) ToString(prefix string) string {
	lines := []string{}
	o.imports.Iterate(func(key concept.String, value interface{}) bool {
		lines = append(lines, value.(*index.ImportIndex).ToString(prefix))
		return false
	})
	o.vars.Iterate(func(key concept.String, value interface{}) bool {
		lines = append(lines, value.(*index.VarIndex).ToString(prefix))
		return false
	})
	o.exports.Iterate(func(key concept.String, value interface{}) bool {
		lines = append(lines, value.(*index.ExportIndex).ToString(prefix))
		return false
	})

	return strings.Join(lines, "\n")
}

func (f *Page) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (o *Page) Type() string {
	return o.seed.Type()
}

func (o *Page) GetSource() concept.Variable {
	return nil
}

func (o *Page) GetClass() concept.Class {
	return nil
}

func (o *Page) GetMapping() *concept.Mapping {
	return nil
}

func (o *Page) IsFunction() bool {
	return false
}

func (o *Page) IsNull() bool {
	return false
}

type PageCreatorParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
	ClosureCreator   func(concept.Closure) concept.Closure
}

type PageCreator struct {
	Seeds map[string]func(string, *Page) string
	param *PageCreatorParam
}

func (s *PageCreator) New() *Page {
	return &Page{
		seed:  s,
		space: s.param.ClosureCreator(nil),
		imports: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.param.NullCreator(),
		}),
		exports: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.param.NullCreator(),
		}),
		vars: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.param.NullCreator(),
		}),
	}
}

func (s *PageCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *PageCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *PageCreator) ToLanguage(language string, instance *Page) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *PageCreator) Type() string {
	return VariablePageType
}

func NewPageCreator(param *PageCreatorParam) *PageCreator {
	return &PageCreator{
		Seeds: map[string]func(string, *Page) string{},
		param: param,
	}
}

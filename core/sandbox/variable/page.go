package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"strings"
)

const (
	VariablePageType = "page"
)

type PageSeed interface {
	NewNull() concept.Null
	NewException(string, string) concept.Exception
	ToLanguage(string, concept.Closure, *Page) string
	Type() string
}

type Page struct {
	seed     PageSeed
	publics  *concept.Mapping
	privates []concept.Index
	space    concept.Closure
}

func (o *Page) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	value, exception := o.GetField(specimen)
	if !nl_interface.IsNil(exception) {
		return nil, exception
	}
	if !value.IsFunction() {
		return nil, o.seed.NewException("runtime error", fmt.Sprintf("There is no public function called %v to be called here.", specimen.ToString("")))
	}
	return value.(concept.Function).Exec(param, nil)
}

func (o *Page) SetImport(specimen concept.String, indexes concept.Index) error {
	return o.SetPrivate(specimen, indexes)
}

func (o *Page) SetPublic(specimen concept.String, indexes concept.Index) error {
	err := o.SetPrivate(specimen, indexes)
	if err != nil {
		return err
	}
	o.publics.Set(specimen, indexes)
	return nil
}

func (o *Page) SetPrivate(specimen concept.String, indexes concept.Index) error {
	if o.space.HasLocal(specimen) {
		return o.seed.NewException("runtime error", fmt.Sprintf("Duplicate identifier: %v.", specimen.ToString("")))
	}
	value, suspend := indexes.Get(o.space)
	if !nl_interface.IsNil(suspend) {
		exception, yes := interrupt.InterruptFamilyInstance.IsException(suspend)
		if yes {
			return exception.(concept.Exception)
		}
		return fmt.Errorf("An illegal interrupt \"%v\" was thrown while declaring variable : %v.", suspend.InterruptType(), specimen.ToString(""))
	}
	o.privates = append(o.privates, indexes)
	o.space.InitLocal(specimen, value)
	return nil
}

func (o *Page) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	if !o.publics.Has(specimen) {
		return o.seed.NewException("runtime error", fmt.Sprintf("There is no public field called %v to be set here.", specimen.ToString("")))
	}
	return o.space.SetLocal(specimen, value)
}

func (o *Page) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	if !o.publics.Has(specimen) {
		return o.seed.NewNull(), nil
	}
	return o.space.GetLocal(specimen)
}

func (o *Page) HasField(specimen concept.String) bool {
	return o.publics.Has(specimen)
}

func (o *Page) SizeField() int {
	return o.publics.Size()
}

func (o *Page) Iterate(on func(concept.String, concept.Variable) bool) bool {
	return o.publics.Iterate(func(key concept.String, _ interface{}) bool {
		value, _ := o.space.GetLocal(key)
		return on(key, value)
	})
}

func (o *Page) ToString(prefix string) string {
	lines := []string{}
	for _, value := range o.privates {
		lines = append(lines, value.ToString(prefix))
	}
	return strings.Join(lines, "\n")
}

func (f *Page) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
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
	Seeds map[string]func(string, concept.Closure, *Page) string
	param *PageCreatorParam
}

func (s *PageCreator) New() *Page {
	return &Page{
		seed:  s,
		space: s.param.ClosureCreator(nil),
		publics: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.param.NullCreator(),
		}),
		privates: []concept.Index{},
	}
}

func (s *PageCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *PageCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *PageCreator) ToLanguage(language string, space concept.Closure, instance *Page) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *PageCreator) Type() string {
	return VariablePageType
}

func NewPageCreator(param *PageCreatorParam) *PageCreator {
	return &PageCreator{
		Seeds: map[string]func(string, concept.Closure, *Page) string{},
		param: param,
	}
}

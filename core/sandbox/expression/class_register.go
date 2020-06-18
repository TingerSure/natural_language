package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type ClassRegisterSeed interface {
	ToLanguage(string, *ClassRegister) string
	NewException(string, string) concept.Exception
}

type ClassRegister struct {
	*adaptor.ExpressionIndex
	object  concept.Index
	class   concept.Index
	mapping map[concept.String]concept.String
	alias   string
	seed    ClassRegisterSeed
}

func (f *ClassRegister) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *ClassRegister) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)
	items := []string{}
	for key, value := range a.mapping {
		items = append(items, fmt.Sprintf("%v%v : %v", subprefix, key.ToString(subprefix), value.ToString(subprefix)))
	}
	return fmt.Sprintf("%v <- %v <%v> {\n%v\n%v}",
		a.object.ToString(prefix),
		a.class.ToString(prefix),
		a.alias,
		strings.Join(items, ",\n"),
		prefix,
	)
}

func (e *ClassRegister) Anticipate(space concept.Closure) concept.Variable {
	return e.object.Anticipate(space)
}

func (a *ClassRegister) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preObject, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, yesObject := variable.VariableFamilyInstance.IsObject(preObject)
	if !yesObject {
		return nil, a.seed.NewException("type error", "Only Object can be use in ClassRegister")
	}

	preClass, suspend := a.class.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	class, yesClass := variable.VariableFamilyInstance.IsClass(preClass)
	if !yesClass {
		return nil, a.seed.NewException("type error", "Only Class can be use in ClassRegister")
	}

	if object.CheckMapping(class, a.mapping) {
		return nil, a.seed.NewException("type error", fmt.Sprintf("Class \"%v\" cannot register to Object \"%v\".", a.class.ToString(""), a.object.ToString("")))
	}

	return object, object.AddClass(class, a.alias, a.mapping)
}

type ClassRegisterCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
}

type ClassRegisterCreator struct {
	Seeds map[string]func(string, *ClassRegister) string
	param *ClassRegisterCreatorParam
}

func (s *ClassRegisterCreator) New(object concept.Index, class concept.Index, mapping map[concept.String]concept.String, alias string) *ClassRegister {
	back := &ClassRegister{
		object:  object,
		class:   class,
		mapping: mapping,
		alias:   alias,
		seed:    s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *ClassRegisterCreator) ToLanguage(language string, instance *ClassRegister) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ClassRegisterCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewClassRegisterCreator(param *ClassRegisterCreatorParam) *ClassRegisterCreator {
	return &ClassRegisterCreator{
		Seeds: map[string]func(string, *ClassRegister) string{},
		param: param,
	}
}

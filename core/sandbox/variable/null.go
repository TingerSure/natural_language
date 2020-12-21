package variable

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariableNullType = "null"
)

type NullSeed interface {
	ToLanguage(string, *Null) string
	Type() string
	GetNullPointerException() concept.Exception
}

type Null struct {
	seed NullSeed
}

func (o *Null) GetClasses() []string {
	return []string{}
}

func (o *Null) GetClass(className string) concept.Class {
	return nil
}

func (o *Null) GetAliases(class string) []string {
	return []string{}
}

func (o *Null) IsClassAlias(class string, alias string) bool {
	return false
}

func (o *Null) UpdateAlias(class string, old, new string) bool {
	return false
}

func (o *Null) CheckMapping(class concept.Class, mapping map[concept.String]concept.String) bool {
	return false
}

func (o *Null) GetMapping(class string, alias string) (map[concept.String]concept.String, concept.Exception) {
	return nil, o.seed.GetNullPointerException().Copy()
}

func (o *Null) RemoveClass(class string, alias string) concept.Exception {
	return o.seed.GetNullPointerException().Copy()
}

func (o *Null) AddClass(class concept.Class, alias string, mapping map[concept.String]concept.String) concept.Exception {
	return o.seed.GetNullPointerException().Copy()
}

func (o *Null) HasField(specimen concept.String) bool {
	return false
}

func (o *Null) InitField(specimen concept.String, defaultValue concept.Variable) concept.Exception {
	return o.seed.GetNullPointerException().Copy()
}

func (o *Null) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	return o.seed.GetNullPointerException().Copy()
}

func (o *Null) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	return nil, o.seed.GetNullPointerException().Copy()
}

func (o *Null) HasMethod(specimen concept.String) bool {
	return false
}

func (o *Null) SetMethod(specimen concept.String, value concept.Function) concept.Exception {
	return o.seed.GetNullPointerException().Copy()
}

func (o *Null) GetMethod(specimen concept.String) (concept.Function, concept.Exception) {
	return nil, o.seed.GetNullPointerException().Copy()
}

func (f *Null) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Null) ToString(prefix string) string {
	return "null"
}

func (n *Null) Type() string {
	return n.seed.Type()
}

type NullCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
}

type NullCreator struct {
	Seeds                map[string]func(string, *Null) string
	param                *NullCreatorParam
	onlyInstance         *Null
	nullPointerException concept.Exception
}

func (s *NullCreator) GetException() concept.Exception {
	if nl_interface.IsNil(s.nullPointerException) {
		s.nullPointerException = s.param.ExceptionCreator("runtime error", "null pointer exception.")
	}
	return s.nullPointerException
}

func (s *NullCreator) New() *Null {
	return s.onlyInstance
}

func (s *NullCreator) ToLanguage(language string, instance *Null) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *NullCreator) Type() string {
	return VariableNullType
}

func (s *NullCreator) GetNullPointerException() concept.Exception {
	return s.nullPointerException
}

func NewNullCreator(param *NullCreatorParam) *NullCreator {
	instance := &NullCreator{
		param: param,
		Seeds: map[string]func(string, *Null) string{},
	}

	instance.onlyInstance = &Null{
		seed: instance,
	}

	return instance
}

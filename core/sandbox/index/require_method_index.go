package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type RequireMethodIndexSeed interface {
	ToLanguage(string, *RequireMethodIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewPreObjectFunction(concept.Function, concept.Object) *variable.PreObjectFunction
	NewNull() concept.Null
}

type RequireMethodIndex struct {
	key    concept.String
	object concept.Index
	seed   RequireMethodIndexSeed
}

const (
	IndexRequireMethodType = "ObjectMethod"
)

func (f *RequireMethodIndex) Type() string {
	return f.seed.Type()
}

func (f *RequireMethodIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *RequireMethodIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *RequireMethodIndex) ToString(prefix string) string {
	return fmt.Sprintf("%s.%s", s.object.ToString(prefix), s.key.ToString(prefix))
}

func (s *RequireMethodIndex) Anticipate(space concept.Closure) concept.Variable {
	preObject := s.object.Anticipate(space)

	object, ok := variable.VariableFamilyInstance.IsObjectHome(preObject)
	if !ok {
		return s.seed.NewNull()
	}

	function, suspend := object.GetMethod(s.key)
	if !nl_interface.IsNil(suspend) {
		return s.seed.NewNull()
	}
	if nl_interface.IsNil(function) {
		return s.seed.NewNull()
	}

	return s.seed.NewPreObjectFunction(function, object)
}

func (s *RequireMethodIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	object, suspend := s.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return s.seed.NewNull(), suspend
	}
	function, suspend := object.GetRequireMethod(s.key)
	if !nl_interface.IsNil(suspend) {
		return s.seed.NewNull(), suspend
	}
	if nl_interface.IsNil(function) {
		return s.seed.NewNull(), s.seed.NewException("runtime error", fmt.Sprintf("Object don't have a require method named %s.", s.key))
	}

	return s.seed.NewPreObjectFunction(function, object.GetSource()), nil
}

func (s *RequireMethodIndex) Set(space concept.Closure, preFunction concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Require method cannot be changed.")
}

type RequireMethodIndexCreatorParam struct {
	ExceptionCreator         func(string, string) concept.Exception
	PreObjectFunctionCreator func(concept.Function, concept.Object) *variable.PreObjectFunction
	NullCreator              func() concept.Null
}

type RequireMethodIndexCreator struct {
	Seeds map[string]func(string, *RequireMethodIndex) string
	param *RequireMethodIndexCreatorParam
}

func (s *RequireMethodIndexCreator) New(object concept.Index, key concept.String) *RequireMethodIndex {
	return &RequireMethodIndex{
		key:    key,
		object: object,
		seed:   s,
	}
}
func (s *RequireMethodIndexCreator) ToLanguage(language string, instance *RequireMethodIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *RequireMethodIndexCreator) Type() string {
	return IndexRequireMethodType
}

func (s *RequireMethodIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *RequireMethodIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *RequireMethodIndexCreator) NewPreObjectFunction(function concept.Function, object concept.Object) *variable.PreObjectFunction {
	return s.param.PreObjectFunctionCreator(function, object)
}

func NewRequireMethodIndexCreator(param *RequireMethodIndexCreatorParam) *RequireMethodIndexCreator {
	return &RequireMethodIndexCreator{
		Seeds: map[string]func(string, *RequireMethodIndex) string{},
		param: param,
	}
}

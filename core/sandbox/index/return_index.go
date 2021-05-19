package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	returnIndexKey = "return"
)

type ReturnIndexSeed interface {
	ToLanguage(string, *ReturnIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewString(string) concept.String
	NewNull() concept.Null
}

type ReturnIndex struct {
	seed ReturnIndexSeed
}

const (
	IndexReturnType = "Return"
)

func (f *ReturnIndex) Type() string {
	return f.seed.Type()
}

func (f *ReturnIndex) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (s *ReturnIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ReturnIndex) ToString(prefix string) string {
	return returnIndexKey
}

func (s *ReturnIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *ReturnIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	funcs := s.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (s *ReturnIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := space.PeekBubble(s.seed.NewString(returnIndexKey))
	return value
}

func (s *ReturnIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(s.seed.NewString(returnIndexKey))
}

func (s *ReturnIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "Return cannot be changed.")
}

type ReturnIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	StringCreator    func(string) concept.String
	NullCreator      func() concept.Null
}

type ReturnIndexCreator struct {
	Seeds map[string]func(string, *ReturnIndex) string
	param *ReturnIndexCreatorParam
}

func (s *ReturnIndexCreator) New() *ReturnIndex {
	return &ReturnIndex{
		seed: s,
	}
}

func (s *ReturnIndexCreator) ToLanguage(language string, instance *ReturnIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ReturnIndexCreator) Type() string {
	return IndexReturnType
}

func (s *ReturnIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ReturnIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *ReturnIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ReturnIndexCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func NewReturnIndexCreator(param *ReturnIndexCreatorParam) *ReturnIndexCreator {
	return &ReturnIndexCreator{
		Seeds: map[string]func(string, *ReturnIndex) string{},
		param: param,
	}
}

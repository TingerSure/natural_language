package index

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	thisIndexKey = "this"
)

type ThisIndexSeed interface {
	ToLanguage(string, concept.Closure, *ThisIndex) string
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
	NewString(string) concept.String
	NewNull() concept.Null
}

type ThisIndex struct {
	seed ThisIndexSeed
}

const (
	IndexThisType = "This"
)

func (f *ThisIndex) Type() string {
	return f.seed.Type()
}

func (f *ThisIndex) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (s *ThisIndex) ToString(prefix string) string {
	return thisIndexKey
}

func (s *ThisIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", fmt.Sprintf("The \"%v\" is not a function.", s.ToString("")))
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *ThisIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	funcs := s.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (s *ThisIndex) Anticipate(space concept.Closure) concept.Variable {
	value, _ := space.PeekBubble(s.seed.NewString(thisIndexKey))
	return value
}

func (s *ThisIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return space.GetBubble(s.seed.NewString(thisIndexKey))
}

func (s *ThisIndex) Set(space concept.Closure, value concept.Variable) concept.Interrupt {
	return s.seed.NewException("read only", "This cannot be changed.")

}

type ThisIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	StringCreator    func(string) concept.String
	NullCreator      func() concept.Null
}

type ThisIndexCreator struct {
	Seeds map[string]func(string, concept.Closure, *ThisIndex) string
	param *ThisIndexCreatorParam
}

func (s *ThisIndexCreator) New() *ThisIndex {
	return &ThisIndex{
		seed: s,
	}
}

func (s *ThisIndexCreator) ToLanguage(language string, space concept.Closure, instance *ThisIndex) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *ThisIndexCreator) Type() string {
	return IndexThisType
}

func (s *ThisIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ThisIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *ThisIndexCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ThisIndexCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func NewThisIndexCreator(param *ThisIndexCreatorParam) *ThisIndexCreator {
	return &ThisIndexCreator{
		Seeds: map[string]func(string, concept.Closure, *ThisIndex) string{},
		param: param,
	}
}

package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type CallSeed interface {
	ToLanguage(string, *Call) string
	NewException(string, string) concept.Exception
	NewNull() concept.Null
	NewParam() concept.Param
}

type Call struct {
	*adaptor.ExpressionIndex
	funcs concept.Index
	param *NewParam
	seed  CallSeed
}

func (c *Call) Function() concept.Index {
	return c.funcs
}

func (c *Call) Param() *NewParam {
	return c.param
}

func (f *Call) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Call) ToString(prefix string) string {
	return fmt.Sprintf("%v(%v)", a.funcs.ToString(prefix), a.param.ToString(prefix))
}

func (a *Call) Anticipate(space concept.Closure) concept.Variable {
	preParam := a.param.Anticipate(space)
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return a.seed.NewParam()
	}
	return a.funcs.CallAnticipate(space, param)
}

func (a *Call) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, a.seed.NewException("type error", "Only Param can are passed to a Function")
	}

	return a.funcs.Call(space, param)
}

type CallCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	ParamCreator           func() concept.Param
	ConstIndexCreator      func(concept.Variable) *index.ConstIndex
	NullCreator            func() concept.Null
	NewParamCreator        func() *NewParam
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
}

type CallCreator struct {
	Seeds        map[string]func(string, *Call) string
	param        *CallCreatorParam
	defaultParam *NewParam
}

func (s *CallCreator) New(funcs concept.Index, param *NewParam) *Call {
	if nl_interface.IsNil(param) {
		param = s.defaultParam
	}
	back := &Call{
		funcs: funcs,
		param: param,
		seed:  s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *CallCreator) ToLanguage(language string, instance *Call) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *CallCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *CallCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *CallCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewCallCreator(param *CallCreatorParam) *CallCreator {
	return &CallCreator{
		Seeds:        map[string]func(string, *Call) string{},
		param:        param,
		defaultParam: param.NewParamCreator(),
	}
}

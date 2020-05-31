package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var ()

type CallSeed interface {
	ToLanguage(string, *Call) string
	NewException(string, string) concept.Exception
}

type Call struct {
	*adaptor.ExpressionIndex
	funcs concept.Index
	param concept.Index
	seed  CallSeed
}

func (c *Call) Function() concept.Index {
	return c.funcs
}

func (c *Call) Param() concept.Index {
	return c.param
}

func (f *Call) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *Call) ToString(prefix string) string {
	return fmt.Sprintf("%v(%v)", a.funcs.ToString(prefix), a.param.ToString(prefix))
}

func (a *Call) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preFuncs, suspend := a.funcs.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	funcs, yesFuncs := variable.VariableFamilyInstance.IsFunctionHome(preFuncs)
	if !yesFuncs {
		return nil, a.seed.NewException("type error", "Only Function can be Called.")
	}

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	yesParam := false
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, a.seed.NewException("type error", "Only Param can are passed to a Function")
	}

	return funcs.Exec(param, nil)
}

type CallCreatorParam struct {
	NewException func(string, string) concept.Exception
}

type CallCreator struct {
	Seeds        map[string]func(string, *Call) string
	param        *CallCreatorParam
	defaultParam concept.Param
}

func (s *CallCreator) New(funcs concept.Index, param concept.Index) *Call {
	if nl_interface.IsNil(param) {
		param = defaultParam
	}
	back := &Call{
		funcs: funcs,
		param: param,
		seed:  s,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

func (s *CallCreator) ToLanguage(language string, instance *Call) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *CallCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewCallCreator(param *CallCreatorParam) *CallCreator {
	return &CallCreator{
		Seeds:        map[string]func(string, *Call) string{},
		param:        param,
		defaultParam: index.NewConstIndex(param.ParamCreator()),
	}
}

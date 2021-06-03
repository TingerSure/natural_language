package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type NewFunctionSeed interface {
	ToLanguage(string, concept.Pool, *NewFunction) (string, concept.Exception)
	NewFunction(concept.Pool) *variable.Function
}

type NewFunction struct {
	*adaptor.ExpressionIndex
	steps   []concept.Pipe
	params  []concept.String
	returns []concept.String
	seed    NewFunctionSeed
}

func (f *NewFunction) SetReturns(returns []concept.Pipe) {
	for _, keyPre := range returns {
		key, yes := index.IndexFamilyInstance.IsKeyIndex(keyPre)
		if !yes {
			panic(fmt.Sprintf("Unsupported index type in NewFunction.SetReturn : %v", keyPre.Type()))
		}
		f.returns = append(f.returns, key.Key())
	}
}

func (f *NewFunction) SetParams(params []concept.Pipe) {
	for _, keyPre := range params {
		key, yes := index.IndexFamilyInstance.IsKeyIndex(keyPre)
		if !yes {
			panic(fmt.Sprintf("Unsupported index type in NewFunction.SetParam : %v", keyPre.Type()))
		}
		f.params = append(f.params, key.Key())
	}
}

func (f *NewFunction) SetSteps(steps []concept.Pipe) {
	f.steps = steps
}

func (f *NewFunction) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewFunction) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	params := []string{}
	for _, param := range a.params {
		params = append(params, param.Value())
	}
	returns := []string{}
	for _, back := range a.returns {
		returns = append(returns, back.Value())
	}
	steps := []string{}
	for _, step := range a.steps {
		steps = append(steps, fmt.Sprintf("%v%v;", subPrefix, step.ToString(subPrefix)))
	}
	return fmt.Sprintf("function(%v) %v {\n%v\n%v}",
		strings.Join(params, ", "),
		strings.Join(returns, ", "),
		strings.Join(steps, "\n"),
		prefix,
	)
}

func (a *NewFunction) Anticipate(space concept.Pool) concept.Variable {
	function := a.seed.NewFunction(space)
	function.AddParamName(a.params...)
	function.AddReturnName(a.returns...)
	function.AnticipateBody().AddStep(a.steps...)
	return function
}

func (a *NewFunction) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	function := a.seed.NewFunction(space)
	function.AddParamName(a.params...)
	function.AddReturnName(a.returns...)
	function.Body().AddStep(a.steps...)
	return function, nil
}

type NewFunctionCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	FunctionCreator        func(concept.Pool) *variable.Function
}

type NewFunctionCreator struct {
	Seeds map[string]func(concept.Pool, *NewFunction) (string, concept.Exception)
	param *NewFunctionCreatorParam
}

func (s *NewFunctionCreator) New() *NewFunction {
	back := &NewFunction{
		seed:    s,
		steps:   []concept.Pipe{},
		params:  []concept.String{},
		returns: []concept.String{},
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewFunctionCreator) NewFunction(parent concept.Pool) *variable.Function {
	return s.param.FunctionCreator(parent)
}

func (s *NewFunctionCreator) ToLanguage(language string, space concept.Pool, instance *NewFunction) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewFunctionCreator(param *NewFunctionCreatorParam) *NewFunctionCreator {
	return &NewFunctionCreator{
		Seeds: map[string]func(concept.Pool, *NewFunction) (string, concept.Exception){},
		param: param,
	}
}

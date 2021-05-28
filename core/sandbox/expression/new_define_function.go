package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type NewDefineFunctionSeed interface {
	ToLanguage(string, concept.Closure, *NewDefineFunction) string
	NewDefineFunction() *variable.DefineFunction
}

type NewDefineFunction struct {
	*adaptor.ExpressionIndex
	params  []concept.String
	returns []concept.String
	seed    NewDefineFunctionSeed
}

func (f *NewDefineFunction) SetReturns(returns []concept.Index) {
	for _, keyPre := range returns {
		key, yes := index.IndexFamilyInstance.IsKeyIndex(keyPre)
		if !yes {
			panic(fmt.Sprintf("Unsupported index type in NewDefineFunction.SetReturn : %v", keyPre.Type()))
		}
		f.returns = append(f.returns, key.Key())
	}
}

func (f *NewDefineFunction) SetParams(params []concept.Index) {
	for _, keyPre := range params {
		key, yes := index.IndexFamilyInstance.IsKeyIndex(keyPre)
		if !yes {
			panic(fmt.Sprintf("Unsupported index type in NewDefineFunction.SetParam : %v", keyPre.Type()))
		}
		f.params = append(f.params, key.Key())
	}
}

func (f *NewDefineFunction) ToLanguage(language string, space concept.Closure) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewDefineFunction) ToString(prefix string) string {
	params := []string{}
	for _, param := range a.params {
		params = append(params, param.Value())
	}
	returns := []string{}
	for _, back := range a.returns {
		returns = append(returns, back.Value())
	}
	return fmt.Sprintf("function(%v) %v", strings.Join(params, ", "), strings.Join(returns, ", "))
}

func (a *NewDefineFunction) Anticipate(space concept.Closure) concept.Variable {
	function := a.seed.NewDefineFunction()
	function.AddParamName(a.params...)
	function.AddReturnName(a.returns...)
	return function
}

func (a *NewDefineFunction) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	function := a.seed.NewDefineFunction()
	function.AddParamName(a.params...)
	function.AddReturnName(a.returns...)
	return function, nil
}

type NewDefineFunctionCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	DefineFunctionCreator  func() *variable.DefineFunction
}

type NewDefineFunctionCreator struct {
	Seeds map[string]func(string, concept.Closure, *NewDefineFunction) string
	param *NewDefineFunctionCreatorParam
}

func (s *NewDefineFunctionCreator) New() *NewDefineFunction {
	back := &NewDefineFunction{
		seed:    s,
		params:  []concept.String{},
		returns: []concept.String{},
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewDefineFunctionCreator) NewDefineFunction() *variable.DefineFunction {
	return s.param.DefineFunctionCreator()
}

func (s *NewDefineFunctionCreator) ToLanguage(language string, space concept.Closure, instance *NewDefineFunction) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewNewDefineFunctionCreator(param *NewDefineFunctionCreatorParam) *NewDefineFunctionCreator {
	return &NewDefineFunctionCreator{
		Seeds: map[string]func(string, concept.Closure, *NewDefineFunction) string{},
		param: param,
	}
}

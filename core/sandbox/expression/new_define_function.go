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
	ToLanguage(string, concept.Pool, *NewDefineFunction) (string, concept.Exception)
	NewDefineFunction([]concept.String, []concept.String) *variable.DefineFunction
}

type NewDefineFunction struct {
	*adaptor.ExpressionIndex
	params  []concept.String
	returns []concept.String
	seed    NewDefineFunctionSeed
}

func (f *NewDefineFunction) SetReturns(returns []concept.Pipe, lines []concept.Line) error {
	for cursor, keyPre := range returns {
		key, yes := index.IndexFamilyInstance.IsKeyIndex(keyPre)
		if !yes {
			return fmt.Errorf("Unsupported index type in NewDefineFunction.SetReturn : %v\n%v", keyPre.Type(), lines[cursor].ToString())
		}
		f.returns = append(f.returns, key.Key())
	}
	return nil
}

func (f *NewDefineFunction) SetParams(params []concept.Pipe, lines []concept.Line) error {
	for cursor, keyPre := range params {
		key, yes := index.IndexFamilyInstance.IsKeyIndex(keyPre)
		if !yes {
			return fmt.Errorf("Unsupported index type in NewDefineFunction.SetParam : %v\n%v", keyPre.Type(), lines[cursor].ToString())
		}
		f.params = append(f.params, key.Key())
	}
	return nil
}

func (f *NewDefineFunction) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
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

func (a *NewDefineFunction) Anticipate(space concept.Pool) concept.Variable {
	function := a.seed.NewDefineFunction(a.params, a.returns)
	return function
}

func (a *NewDefineFunction) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	function := a.seed.NewDefineFunction(a.params, a.returns)
	return function, nil
}

type NewDefineFunctionCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	DefineFunctionCreator  func([]concept.String, []concept.String) *variable.DefineFunction
}

type NewDefineFunctionCreator struct {
	Seeds map[string]func(concept.Pool, *NewDefineFunction) (string, concept.Exception)
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

func (s *NewDefineFunctionCreator) NewDefineFunction(paramNames []concept.String, returnNames []concept.String) *variable.DefineFunction {
	return s.param.DefineFunctionCreator(paramNames, returnNames)
}

func (s *NewDefineFunctionCreator) ToLanguage(language string, space concept.Pool, instance *NewDefineFunction) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewDefineFunctionCreator(param *NewDefineFunctionCreatorParam) *NewDefineFunctionCreator {
	return &NewDefineFunctionCreator{
		Seeds: map[string]func(concept.Pool, *NewDefineFunction) (string, concept.Exception){},
		param: param,
	}
}

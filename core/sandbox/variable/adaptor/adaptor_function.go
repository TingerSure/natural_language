package adaptor

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type AdaptorFunctionParam struct {
	NullCreator           func() concept.Null
	ExceptionCreator      func(string, string) concept.Exception
	ParamCreator          func() concept.Param
	DelayStringCreator    func(string) concept.String
	StringCreator         func(string) concept.String
	DelayFunctionCreator  func(func() concept.Function) concept.Function
	ArrayCreator          func() concept.Array
	SystemFunctionCreator func(
		funcs func(concept.Param, concept.Variable) (concept.Param, concept.Exception),
		anticipateFuncs func(concept.Param, concept.Variable) concept.Param,
		paramNames []concept.String,
		returnNames []concept.String,
	) concept.Function
}

type AdaptorFunction struct {
	*AdaptorVariable
	param               *AdaptorFunctionParam
	paramNames          []concept.String
	returnNames         []concept.String
	languageOnCallSeeds map[string]func(concept.Function, concept.Pool, string, concept.Param) string
}

func (s *AdaptorFunction) ParamNames() []concept.String {
	return s.paramNames
}

func (s *AdaptorFunction) ReturnNames() []concept.String {
	return s.returnNames
}

func (f *AdaptorFunction) AddParamName(paramNames ...concept.String) {
	f.paramNames = append(f.paramNames, paramNames...)
}

func (f *AdaptorFunction) AddReturnName(returnNames ...concept.String) {
	f.returnNames = append(f.returnNames, returnNames...)
}

func (o *AdaptorFunction) IsFunction() bool {
	return true
}

func (a *AdaptorFunction) GetLanguageOnCallSeed(language string) func(concept.Function, concept.Pool, string, concept.Param) string {
	return a.languageOnCallSeeds[language]
}

func (a *AdaptorFunction) SetLanguageOnCallSeed(language string, seed func(concept.Function, concept.Pool, string, concept.Param) string) {
	a.languageOnCallSeeds[language] = seed
}

func (a *AdaptorFunction) ToCallLanguageAdaptor(funcs concept.Function, language string, space concept.Pool, self string, param concept.Param) string {
	seed := funcs.GetLanguageOnCallSeed(language)
	if seed == nil {
		return funcs.ToLanguage(language, space)
	}
	return seed(funcs, space, self, param)
}

func (a *AdaptorFunction) ParamFormat(params concept.Param) concept.Param {
	keys := a.ParamNames()
	instance := a.param.ParamCreator()
	params.Iterate(func(target concept.String, value concept.Variable) bool {
		for _, src := range keys {
			if target.Equal(src) {
				instance.Set(src, value)
				return false
			}
		}
		instance.Set(target, value)
		return false
	})
	return instance
}

func (a *AdaptorFunction) ReturnFormat(back concept.String) concept.String {
	for _, name := range a.ReturnNames() {
		if name.Equal(back) {
			return name
		}
	}
	return back
}

func NewAdaptorFunction(param *AdaptorFunctionParam) *AdaptorFunction {
	instance := &AdaptorFunction{
		AdaptorVariable: NewAdaptorVariable(&AdaptorVariableParam{
			NullCreator:      param.NullCreator,
			ExceptionCreator: param.ExceptionCreator,
		}),
		param:               param,
		languageOnCallSeeds: map[string]func(concept.Function, concept.Pool, string, concept.Param) string{},
	}

	instance.SetField(param.DelayStringCreator("paramList"), param.DelayFunctionCreator(instance.fieldParamList))
	instance.SetField(param.DelayStringCreator("returnList"), param.DelayFunctionCreator(instance.fieldReturnList))

	return instance
}

func (s *AdaptorFunction) fieldParamList() concept.Function {
	backList := s.param.StringCreator("list")
	return s.param.SystemFunctionCreator(
		func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			paramNames := s.param.ArrayCreator()
			for _, paramName := range s.ParamNames() {
				paramNames.Append(paramName)
			}
			back := s.param.ParamCreator()
			back.Set(backList, paramNames)
			return back, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{
			backList,
		},
	)
}

func (s *AdaptorFunction) fieldReturnList() concept.Function {
	backList := s.param.StringCreator("list")
	return s.param.SystemFunctionCreator(
		func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			returnNames := s.param.ArrayCreator()
			for _, returnName := range s.ReturnNames() {
				returnNames.Append(returnName)
			}
			back := s.param.ParamCreator()
			back.Set(backList, returnNames)
			return back, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{
			backList,
		},
	)
}

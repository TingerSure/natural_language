package variable

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariableDelayFunctionType = "delay_function"
	FunctionDelayFunctionType = "delay"
)

type DelayFunctionSeed interface {
	ToLanguage(string, concept.Pool, *DelayFunction) (string, concept.Exception)
	Type() string
}

type DelayFunction struct {
	funcs  concept.Function
	create func() concept.Function
	seed   DelayFunctionSeed
}

func (o *DelayFunction) init() {
	if nl_interface.IsNil(o.funcs) {
		o.funcs = o.create()
	}
}

func (o *DelayFunction) IsFunction() bool {
	return true
}

func (o *DelayFunction) IsNull() bool {
	return false
}

func (o *DelayFunction) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	o.init()
	return o.funcs.SetField(specimen, value)
}

func (o *DelayFunction) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	o.init()
	return o.funcs.GetField(specimen)
}

func (o *DelayFunction) HasField(specimen concept.String) bool {
	o.init()
	return o.funcs.HasField(specimen)
}

func (o *DelayFunction) KeyField(specimen concept.String) concept.String {
	o.init()
	return o.funcs.KeyField(specimen)
}

func (o *DelayFunction) SizeField() int {
	o.init()
	return o.funcs.SizeField()
}

func (o *DelayFunction) Iterate(on func(concept.String, concept.Variable) bool) bool {
	o.init()
	return o.funcs.Iterate(on)
}

func (o *DelayFunction) ToString(prefix string) string {
	o.init()
	return o.funcs.ToString(prefix)
}

func (a *DelayFunction) GetLanguageOnCallSeed(language string) func(concept.Function, concept.Pool, string, concept.Param) (string, concept.Exception) {
	a.init()
	return a.funcs.GetLanguageOnCallSeed(language)
}

func (a *DelayFunction) SetLanguageOnCallSeed(language string, seed func(concept.Function, concept.Pool, string, concept.Param) (string, concept.Exception)) {
	a.init()
	a.funcs.SetLanguageOnCallSeed(language, seed)
}

func (f *DelayFunction) ParamFormat(params concept.Param) concept.Param {
	f.init()
	return f.funcs.ParamFormat(params)
}

func (f *DelayFunction) ReturnFormat(back concept.String) concept.String {
	f.init()
	return f.funcs.ReturnFormat(back)
}

func (f *DelayFunction) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	f.init()
	return f.funcs.Call(specimen, param)
}

func (f *DelayFunction) ToCallLanguage(language string, space concept.Pool, self string, param concept.Param) (string, concept.Exception) {
	f.init()
	return f.funcs.ToCallLanguage(language, space, self, param)
}

func (f *DelayFunction) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (f *DelayFunction) AddParamName(paramNames ...concept.String) {
	f.init()
	f.funcs.AddParamName(paramNames...)
}

func (f *DelayFunction) AddReturnName(returnNames ...concept.String) {
	f.init()
	f.funcs.AddReturnName(returnNames...)
}

func (s *DelayFunction) ParamNames() []concept.String {
	s.init()
	return s.funcs.ParamNames()
}

func (s *DelayFunction) ReturnNames() []concept.String {
	s.init()
	return s.funcs.ReturnNames()
}

func (f *DelayFunction) Exec(params concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	f.init()
	return f.funcs.Exec(params, object)
}

func (s *DelayFunction) Type() string {
	return s.seed.Type()
}

func (s *DelayFunction) FunctionType() string {
	return FunctionDelayFunctionType
}

type DelayFunctionCreatorParam struct {
}

type DelayFunctionCreator struct {
	Seeds map[string]func(concept.Pool, *DelayFunction) (string, concept.Exception)
	param *DelayFunctionCreatorParam
}

func (s *DelayFunctionCreator) New(create func() concept.Function) *DelayFunction {
	return &DelayFunction{
		create: create,
		seed:   s,
	}
}

func (s *DelayFunctionCreator) ToLanguage(language string, space concept.Pool, instance *DelayFunction) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *DelayFunctionCreator) Type() string {
	return VariableDelayFunctionType
}

func NewDelayFunctionCreator(param *DelayFunctionCreatorParam) *DelayFunctionCreator {
	return &DelayFunctionCreator{
		Seeds: map[string]func(concept.Pool, *DelayFunction) (string, concept.Exception){},
		param: param,
	}
}

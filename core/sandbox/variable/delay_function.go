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
	ToLanguage(string, *DelayFunction) string
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

func (o *DelayFunction) GetSource() concept.Variable {
	o.init()
	return o.funcs.GetSource()
}

func (o *DelayFunction) GetClass() concept.Class {
	o.init()
	return o.funcs.GetClass()
}

func (o *DelayFunction) GetMapping() *concept.Mapping {
	o.init()
	return o.funcs.GetMapping()
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

func (a *DelayFunction) GetLanguageOnCallSeed(language string) func(concept.Function, concept.Param) string {
	a.init()
	return a.funcs.GetLanguageOnCallSeed(language)
}

func (a *DelayFunction) SetLanguageOnCallSeed(language string, seed func(concept.Function, concept.Param) string) {
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

func (f *DelayFunction) ToLanguage(language string) string {
	f.init()
	return f.seed.ToLanguage(language, f)
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

func (f *DelayFunction) Anticipate(params concept.Param, object concept.Variable) concept.Param {
	f.init()
	return f.funcs.Anticipate(params, object)
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
	Seeds map[string]func(string, *DelayFunction) string
	param *DelayFunctionCreatorParam
}

func (s *DelayFunctionCreator) New(create func() concept.Function) *DelayFunction {
	return &DelayFunction{
		create: create,
		seed:   s,
	}
}

func (s *DelayFunctionCreator) ToLanguage(language string, instance *DelayFunction) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *DelayFunctionCreator) Type() string {
	return VariableDelayFunctionType
}

func NewDelayFunctionCreator(param *DelayFunctionCreatorParam) *DelayFunctionCreator {
	return &DelayFunctionCreator{
		Seeds: map[string]func(string, *DelayFunction) string{},
		param: param,
	}
}

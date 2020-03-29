package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/component"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableParamType = "param"
	ParamDefaultKey   = "default"
)

type Param struct {
	values *component.Mapping
}

func (a *Param) ToString(prefix string) string {
	if 0 == a.values.Size() {
		return "{}"
	}
	subPrefix := fmt.Sprintf("%v\t", prefix)
	paramsToString := make([]string, 0, a.values.Size())

	a.values.Iterate(func(key concept.String, value interface{}) bool {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
		return false
	})

	return fmt.Sprintf("{\n%v\n%v}", strings.Join(paramsToString, ",\n"), prefix)
}

func (o *Param) Type() string {
	return VariableParamType
}

func (o *Param) Set(key concept.String, value concept.Variable) concept.Param {
	o.values.Set(key, value)
	return o
}

func (o *Param) Get(key concept.String) concept.Variable {
	return o.values.Get(key).(concept.Variable)
}

func (o *Param) Copy() *Param {
	param := NewParam()
	o.values.Iterate(func(key concept.String, value interface{}) bool {
		param.Set(key, value.(concept.Variable))
		return false
	})
	return param
}

func (o *Param) Init(iterator func(func(concept.String, concept.Variable) bool) bool) {
	iterator(func(key concept.String, value concept.Variable) bool {
		o.Set(key, value)
		return false
	})
}

func NewParam() *Param {
	return &Param{
		values: component.NewMapping(&component.MappingParam{
			AutoInit:   true,
			EmptyValue: NewNull(),
		}),
	}
}

func NewParamWithIterate(iterator func(func(concept.String, concept.Variable) bool) bool) *Param {
	param := NewParam()
	param.Init(iterator)
	return param
}

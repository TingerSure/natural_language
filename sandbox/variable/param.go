package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"strings"
)

const (
	VariableParamType = "param"
	ParamDefaultKey   = "default"
)

type Param struct {
	values map[string]concept.Variable
}

func (a *Param) ToString(prefix string) string {
	if 0 == len(a.values) {
		return "{}"
	}
	subPrefix := fmt.Sprintf("%v\t", prefix)
	paramsToString := make([]string, 0, len(a.values))

	for key, value := range a.values {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key, value.ToString(subPrefix)))
	}

	return fmt.Sprintf("{\n%v\n%v}", strings.Join(paramsToString, ",\n"), prefix)
}

func (o *Param) Type() string {
	return VariableParamType
}

func (o *Param) Set(key string, value concept.Variable) {
	o.values[key] = value
}

func (o *Param) Get(key string) concept.Variable {
	return o.values[key]
}

func (o *Param) Copy() *Param {
	param := NewParam()
	for key, value := range o.values {
		param.Set(key, value)
	}
	return param
}

func (o *Param) Init(params map[string]concept.Variable) {
	o.values = params
}

func NewParam() *Param {
	return &Param{
		values: make(map[string]concept.Variable),
	}
}

func NewParamWithInit(params map[string]concept.Variable) *Param {
	return &Param{
		values: params,
	}
}

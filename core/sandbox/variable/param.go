package variable

import (
	"fmt"

	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

const (
	VariableParamType = "param"
	ParamDefaultKey   = "default"
)

type Param struct {
	values *concept.Mapping
}

var (
	ParamLanguageSeeds = map[string]func(string, *Param) string{}
)

func (f *Param) ToLanguage(language string) string {
	seed := ParamLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
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

func (o *Param) Iterate(on func(concept.String, concept.Variable) bool) bool {
	return o.values.Iterate(func(key concept.String, value interface{}) bool {
		return on(key, value.(concept.Variable))
	})
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
		values: concept.NewMapping(&concept.MappingParam{
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

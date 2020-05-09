package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type NewParam struct {
	*adaptor.ExpressionIndex
	values map[concept.String]concept.Index
}

var (
	NewParamLanguageSeeds = map[string]func(string, *NewParam) string{}
)

func (f *NewParam) Iterate(on func(concept.String, concept.Index) bool) bool {
	for key, value := range f.values {
		if on(key, value) {
			return true
		}
	}
	return false
}

func (f *NewParam) ToLanguage(language string) string {
	seed := NewParamLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (a *NewParam) ToString(prefix string) string {
	if 0 == len(a.values) {
		return "{}"
	}
	subPrefix := fmt.Sprintf("%v\t", prefix)
	paramsToString := make([]string, 0, len(a.values))

	for key, value := range a.values {
		paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.ToString(subPrefix)))
	}

	return fmt.Sprintf("{\n%v\n%v}", strings.Join(paramsToString, ",\n"), prefix)
}

func (a *NewParam) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	if len(a.values) == 0 {
		return variable.NewParam(), nil
	}
	var suspend concept.Interrupt = nil

	param := variable.NewParamWithIterate(func(on func(concept.String, concept.Variable) bool) bool {
		for key, index := range a.values {
			value, subSuspend := index.Get(space)
			if !nl_interface.IsNil(subSuspend) {
				suspend = subSuspend
				return true
			}
			if on(key, value) {
				return true
			}
		}
		return false

	})

	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return param, nil
}

func NewNewParamWithInit(values map[concept.String]concept.Index) *NewParam {
	back := NewNewParam()
	back.values = values
	return back
}

func NewNewParam() *NewParam {
	back := &NewParam{}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

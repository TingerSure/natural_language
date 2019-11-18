package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/variable"
	"strings"
)

type NewParam struct {
	*adaptor.ExpressionIndex
	values map[string]concept.Index
}

func (a *NewParam) ToString(prefix string) string {
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

func (a *NewParam) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	if len(a.values) == 0 {
		return variable.NewParam(), nil
	}
	params := map[string]concept.Variable{}
	for key, index := range a.values {
		value, suspend := index.Get(space)
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
		}
		params[key] = value
	}
	return variable.NewParamWithInit(params), nil
}

func NewNewParamWithInit(values map[string]concept.Index) *NewParam {
	back := NewNewParam()
	back.values = values
	return back
}

func NewNewParam() *NewParam {
	back := &NewParam{}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

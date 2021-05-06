package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"strings"
)

type NewParamSeed interface {
	ToLanguage(string, *NewParam) string
	NewParam() concept.Param
}

type NewParam struct {
	*adaptor.ExpressionIndex
	values *concept.Mapping
	list   []concept.Index
	types  int
	seed   NewParamSeed
}

func (f *NewParam) SetArray(list []concept.Index) {
	f.types = concept.ParamTypeList
	f.list = list
}

func (f *NewParam) SetKeyValue(keyValues []concept.Index) {
	f.types = concept.ParamTypeKeyValue
	for _, keyValuePre := range keyValues {
		keyValue, yes := index.IndexFamilyInstance.IsKeyValueIndex(keyValuePre)
		if !yes {
			panic(fmt.Sprintf("Unsupported index type in SetKeyValue : %v", keyValuePre.Type()))
		}
		f.values.Set(keyValue.Key(), keyValue.Value())
	}
}

func (f *NewParam) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewParam) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	if a.types == concept.ParamTypeList {
		if len(a.list) == 0 {
			return ""
		}
		paramsToString := make([]string, 0, len(a.list))
		for _, value := range a.list {
			paramsToString = append(paramsToString, value.ToString(subPrefix))
		}
		return strings.Join(paramsToString, ", ")
	}
	if a.types == concept.ParamTypeKeyValue {
		if a.values.Size() == 0 {
			return ""
		}
		paramsToString := make([]string, 0, a.values.Size())
		a.values.Iterate(func(key concept.String, value interface{}) bool {
			paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.Value(), value.(concept.ToString).ToString(subPrefix)))
			return false
		})
		return fmt.Sprintf("%v\n%v\n%v", prefix, strings.Join(paramsToString, ",\n"), prefix)
	}
	return ""
}

func (a *NewParam) Anticipate(space concept.Closure) concept.Variable {
	param := a.seed.NewParam()
	if a.types == concept.ParamTypeList {
		for _, item := range a.list {
			param.AppendIndex(item.Anticipate(space))
		}
		return param
	}
	if a.types == concept.ParamTypeKeyValue {
		a.values.Iterate(func(key concept.String, value interface{}) bool {
			param.Set(key, value.(concept.Index).Anticipate(space))
			return false
		})
		return param
	}
	return param
}

func (a *NewParam) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	param := a.seed.NewParam()
	var suspend concept.Interrupt = nil
	var value concept.Variable = nil

	if a.types == concept.ParamTypeList {
		for _, item := range a.list {
			value, suspend = item.Get(space)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend
			}
			param.AppendIndex(value)
		}
		return param, nil
	}

	if a.types == concept.ParamTypeKeyValue {
		if a.values.Iterate(func(key concept.String, item interface{}) bool {
			value, suspend = item.(concept.Index).Get(space)
			if !nl_interface.IsNil(suspend) {
				return true
			}
			param.Set(key, value)
			return false
		}) {
			return nil, suspend
		}
		return param, nil
	}
	return param, nil
}

type NewParamCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ParamCreator           func() concept.Param
	NullCreator            func() concept.Null
}

type NewParamCreator struct {
	Seeds map[string]func(string, *NewParam) string
	param *NewParamCreatorParam
}

func (s *NewParamCreator) New() *NewParam {
	back := &NewParam{
		seed: s,
		values: concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: s.param.NullCreator(),
		}),
		types: concept.ParamTypeKeyValue,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewParamCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *NewParamCreator) ToLanguage(language string, instance *NewParam) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewParamCreator(param *NewParamCreatorParam) *NewParamCreator {
	return &NewParamCreator{
		Seeds: map[string]func(string, *NewParam) string{},
		param: param,
	}
}

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
	ToLanguage(string, concept.Pool, *NewParam) string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
}

type NewParam struct {
	*adaptor.ExpressionIndex
	values *nl_interface.Mapping
	list   []concept.Pipe
	types  int
	seed   NewParamSeed
}

func (f *NewParam) SetArray(list []concept.Pipe) {
	f.types = concept.ParamTypeList
	f.list = list
}

func (f *NewParam) SetKeyValue(keyValues []concept.Pipe) {
	f.types = concept.ParamTypeKeyValue
	for _, keyValuePre := range keyValues {
		keyValue, yes := index.IndexFamilyInstance.IsKeyValueIndex(keyValuePre)
		if !yes {
			panic(fmt.Sprintf("Unsupported index type in NewParam.SetKeyValue : %v", keyValuePre.Type()))
		}
		f.values.Set(keyValue.Key(), keyValue.Value())
	}
}

func (f *NewParam) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewParam) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	if a.types == concept.ParamTypeList {
		if len(a.list) == 0 {
			return ""
		}
		paramsToString := make([]string, 0, len(a.list))
		for _, value := range a.list {
			paramsToString = append(paramsToString, value.ToString(prefix))
		}
		return strings.Join(paramsToString, ", ")
	}
	if a.types == concept.ParamTypeKeyValue {
		if a.values.Size() == 0 {
			return ""
		}
		paramsToString := make([]string, 0, a.values.Size())
		a.values.Iterate(func(key nl_interface.Key, value interface{}) bool {
			paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.(concept.String).Value(), value.(concept.ToString).ToString(subPrefix)))
			return false
		})
		return fmt.Sprintf("%v\n%v\n%v", prefix, strings.Join(paramsToString, ",\n"), prefix)
	}
	return ""
}

func (a *NewParam) Anticipate(space concept.Pool) concept.Variable {
	param := a.seed.NewParam()
	if a.types == concept.ParamTypeList {
		for _, item := range a.list {
			param.AppendIndex(item.Anticipate(space))
		}
		return param
	}
	if a.types == concept.ParamTypeKeyValue {
		a.values.Iterate(func(key nl_interface.Key, value interface{}) bool {
			param.Set(key.(concept.String), value.(concept.Pipe).Anticipate(space))
			return false
		})
		return param
	}
	return param
}

func (a *NewParam) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
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
		if a.values.Iterate(func(key nl_interface.Key, item interface{}) bool {
			value, suspend = item.(concept.Pipe).Get(space)
			if !nl_interface.IsNil(suspend) {
				return true
			}
			param.Set(key.(concept.String), value)
			return false
		}) {
			return nil, suspend
		}
		return param, nil
	}
	return nil, a.seed.NewException("system panic", fmt.Sprintf("Unknown param types in NewParam.Exec", a.types))
}

func (a *NewParam) Iterate(names []concept.String, on func(key concept.String, line concept.Pipe) bool) bool {
	if a.types == concept.ParamTypeList {
		for index, item := range a.list {
			if index >= len(names) {
				return false
			}
			if on(names[index], item) {
				return true
			}
		}
		return false
	}
	if a.types == concept.ParamTypeKeyValue {
		for _, name := range names {
			item := a.values.Get(name)
			if item == nil {
				continue
			}
			if on(name, item.(concept.Pipe)) {
				return true
			}
		}
		return false
	}
	return false
}

type NewParamCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ParamCreator           func() concept.Param
	ExceptionCreator       func(string, string) concept.Exception
	NullCreator            func() concept.Null
}

type NewParamCreator struct {
	Seeds map[string]func(string, concept.Pool, *NewParam) string
	param *NewParamCreatorParam
}

func (s *NewParamCreator) New() *NewParam {
	back := &NewParam{
		seed: s,
		values: nl_interface.NewMapping(&nl_interface.MappingParam{
			AutoInit:   true,
			EmptyValue: nil,
		}),
		types: concept.ParamTypeKeyValue,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewParamCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func (s *NewParamCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *NewParamCreator) ToLanguage(language string, space concept.Pool, instance *NewParam) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewNewParamCreator(param *NewParamCreatorParam) *NewParamCreator {
	return &NewParamCreator{
		Seeds: map[string]func(string, concept.Pool, *NewParam) string{},
		param: param,
	}
}

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
	ToLanguage(string, concept.Pool, *NewParam) (string, concept.Exception)
	NewException(string, string) concept.Exception
	NewParam() concept.Param
}

type NewParam struct {
	*adaptor.ExpressionIndex
	values []*index.KeyValueIndex
	list   []concept.Pipe
	types  int
	seed   NewParamSeed
}

func (f *NewParam) SetArray(list []concept.Pipe) {
	f.types = concept.ParamTypeList
	f.list = list
}

func (f *NewParam) SetKeyValue(keyValues []concept.Pipe, lines []concept.Line) error {
	f.types = concept.ParamTypeKeyValue
	valueMap := map[string]bool{}
	for cursor, keyValuePre := range keyValues {
		keyValue, yes := index.IndexFamilyInstance.IsKeyValueIndex(keyValuePre)
		if !yes {
			return fmt.Errorf("Unsupported index type in NewParam.SetKeyValue : %v\n%v", keyValuePre.Type(), lines[cursor].ToString())
		}
		if valueMap[keyValue.Key().Value()] {
			return fmt.Errorf("Duplicate key: '%v'\n%v", keyValue.Key().Value(), lines[cursor].ToString())
		}
		valueMap[keyValue.Key().Value()] = true
		f.values = append(f.values, keyValue)
	}
	return nil
}

func (f *NewParam) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
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
		if len(a.values) == 0 {
			return ""
		}
		paramsToString := make([]string, 0, len(a.values))
		for _, keyValue := range a.values {
			paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, keyValue.Key().Value(), keyValue.Value().ToString(subPrefix)))
		}
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
		for _, keyValue := range a.values {
			param.Set(keyValue.Key(), keyValue.Value().Anticipate(space))
		}
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
		for _, keyValue := range a.values {
			value, suspend = keyValue.Value().Get(space)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend
			}
			param.Set(keyValue.Key(), value)
		}
		return param, nil
	}
	panic(fmt.Sprintf("Unknown param types in NewParam.Exec", a.types))
}

func (a *NewParam) Iterate(names []concept.String, on func(concept.String, concept.Pipe) bool) bool {
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
			for _, keyValue := range a.values {
				if name.Equal(keyValue.Key()) {
					if on(name, keyValue.Value()) {
						return true
					}
					continue
				}
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
	Seeds map[string]func(concept.Pool, *NewParam) (string, concept.Exception)
	param *NewParamCreatorParam
}

func (s *NewParamCreator) New() *NewParam {
	back := &NewParam{
		seed:   s,
		values: []*index.KeyValueIndex{},
		types:  concept.ParamTypeKeyValue,
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

func (s *NewParamCreator) ToLanguage(language string, space concept.Pool, instance *NewParam) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewParamCreator(param *NewParamCreatorParam) *NewParamCreator {
	return &NewParamCreator{
		Seeds: map[string]func(concept.Pool, *NewParam) (string, concept.Exception){},
		param: param,
	}
}

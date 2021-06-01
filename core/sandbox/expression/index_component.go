package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type IndexComponentSeed interface {
	NewException(string, string) concept.Exception
	ToLanguage(string, concept.Pool, *IndexComponent) string
}

type IndexComponent struct {
	*adaptor.ExpressionIndex
	field  concept.Pipe
	object concept.Pipe
	seed   IndexComponentSeed
}

func (f *IndexComponent) ToLanguage(language string, space concept.Pool) string {
	return f.seed.ToLanguage(language, space, f)
}

func (a *IndexComponent) ToString(prefix string) string {
	return fmt.Sprintf("%v[%v]", a.object.ToString(prefix), a.field.ToString(prefix))
}

func (a *IndexComponent) Anticipate(space concept.Pool) concept.Variable {
	fieldPre := a.field.Anticipate(space)
	fieldNumber, yes := variable.VariableFamilyInstance.IsNumber(fieldPre)
	if yes {
		return a.indexAnticipate(space, fieldNumber)
	}
	fieldString, yes := variable.VariableFamilyInstance.IsStringHome(fieldPre)
	if yes {
		return a.stringAnticipate(space, fieldString)
	}
	return nil
}

func (a *IndexComponent) stringAnticipate(space concept.Pool, field concept.String) concept.Variable {
	object := a.object.Anticipate(space)
	value, _ := object.GetField(field)
	return value
}

func (a *IndexComponent) indexAnticipate(space concept.Pool, field concept.Number) concept.Variable {
	arrayPre := a.object.Anticipate(space)
	array, yes := variable.VariableFamilyInstance.IsArray(arrayPre)
	if !yes {
		return nil
	}
	value, _ := array.Get(int(field.Value()))
	return value
}

func (a *IndexComponent) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	fieldPre, suspend := a.field.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	fieldNumber, yes := variable.VariableFamilyInstance.IsNumber(fieldPre)
	if yes {
		return a.indexGet(space, fieldNumber)
	}
	fieldString, yes := variable.VariableFamilyInstance.IsStringHome(fieldPre)
	if yes {
		return a.stringGet(space, fieldString)
	}
	return nil, a.seed.NewException("runtime error", fmt.Sprintf("%v is not a string or number.", a.field.ToString("")))
}

func (a *IndexComponent) stringGet(space concept.Pool, field concept.String) (concept.Variable, concept.Interrupt) {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return object.GetField(field)
}

func (a *IndexComponent) indexGet(space concept.Pool, field concept.Number) (concept.Variable, concept.Interrupt) {
	arrayPre, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	array, yes := variable.VariableFamilyInstance.IsArray(arrayPre)
	if !yes {
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("%v is not an array.", a.object.ToString("")))
	}
	return array.Get(int(field.Value()))
}

func (a *IndexComponent) Set(space concept.Pool, value concept.Variable) concept.Interrupt {
	fieldPre, suspend := a.field.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	fieldNumber, yes := variable.VariableFamilyInstance.IsNumber(fieldPre)
	if yes {
		return a.indexSet(space, fieldNumber, value)
	}
	fieldString, yes := variable.VariableFamilyInstance.IsStringHome(fieldPre)
	if yes {
		return a.stringSet(space, fieldString, value)
	}
	return a.seed.NewException("runtime error", fmt.Sprintf("%v is not a string or number.", a.field.ToString("")))
}

func (a *IndexComponent) stringSet(space concept.Pool, field concept.String, value concept.Variable) concept.Interrupt {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	return object.SetField(field, value)
}

func (a *IndexComponent) indexSet(space concept.Pool, field concept.Number, value concept.Variable) concept.Interrupt {
	arrayPre, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}
	array, yes := variable.VariableFamilyInstance.IsArray(arrayPre)
	if !yes {
		return a.seed.NewException("runtime error", fmt.Sprintf("%v is not an array.", a.object.ToString("")))
	}
	return array.Set(int(field.Value()), value)
}

func (a *IndexComponent) Call(space concept.Pool, param concept.Param) (concept.Param, concept.Exception) {
	fieldPre, suspend := a.field.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend.(concept.Exception)
	}
	fieldNumber, yes := variable.VariableFamilyInstance.IsNumber(fieldPre)
	if yes {
		return a.indexCall(space, fieldNumber, param)
	}
	fieldString, yes := variable.VariableFamilyInstance.IsStringHome(fieldPre)
	if yes {
		return a.stringCall(space, fieldString, param)
	}
	return nil, a.seed.NewException("runtime error", fmt.Sprintf("%v is not a string or number.", a.field.ToString("")))
}

func (a *IndexComponent) stringCall(space concept.Pool, field concept.String, param concept.Param) (concept.Param, concept.Exception) {
	object, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend.(concept.Exception)
	}
	return object.Call(field, param)
}

func (a *IndexComponent) indexCall(space concept.Pool, field concept.Number, param concept.Param) (concept.Param, concept.Exception) {
	arrayPre, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend.(concept.Exception)
	}
	array, yes := variable.VariableFamilyInstance.IsArray(arrayPre)
	if !yes {
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("%v is not an array.", a.object.ToString("")))
	}
	funcsPre, suspend := array.Get(int(field.Value()))
	if !nl_interface.IsNil(suspend) {
		return nil, suspend.(concept.Exception)
	}
	funcs, yes := variable.VariableFamilyInstance.IsFunctionHome(funcsPre)
	if !yes {
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("%v is not a function.", a.ToString("")))
	}
	return funcs.Exec(param, nil)
}

type IndexComponentCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
}

type IndexComponentCreator struct {
	Seeds map[string]func(string, concept.Pool, *IndexComponent) string
	param *IndexComponentCreatorParam
}

func (s *IndexComponentCreator) New(object concept.Pipe, field concept.Pipe) *IndexComponent {
	back := &IndexComponent{
		field:  field,
		object: object,
		seed:   s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *IndexComponentCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *IndexComponentCreator) ToLanguage(language string, space concept.Pool, instance *IndexComponent) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func NewIndexComponentCreator(param *IndexComponentCreatorParam) *IndexComponentCreator {
	return &IndexComponentCreator{
		Seeds: map[string]func(string, concept.Pool, *IndexComponent) string{},
		param: param,
	}
}

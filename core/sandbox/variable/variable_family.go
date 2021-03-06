package variable

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

var (
	VariableFamilyInstance *VariableFamily = newVariableFamily()
)

type VariableFamily struct {
}

func (v *VariableFamily) IsPool(value concept.Variable) (concept.Pool, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariablePoolType {
		pool, yes := value.(concept.Pool)
		return pool, yes
	}
	return nil, false
}

func (v *VariableFamily) IsNull(value concept.Variable) (concept.Null, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableNullType {
		null, yes := value.(concept.Null)
		return null, yes
	}
	return nil, false
}

func (v *VariableFamily) IsException(value concept.Variable) (concept.Exception, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableExceptionType {
		exception, yes := value.(concept.Exception)
		return exception, yes
	}
	return nil, false
}

func (v *VariableFamily) IsString(value concept.Variable) (*String, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableStringType {
		str, yes := value.(*String)
		return str, yes
	}
	return nil, false
}

func (v *VariableFamily) IsDelayString(value concept.Variable) (*DelayString, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableDelayStringType {
		str, yes := value.(*DelayString)
		return str, yes
	}
	return nil, false
}

func (v *VariableFamily) IsNumber(value concept.Variable) (*Number, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableNumberType {
		number, yes := value.(*Number)
		return number, yes
	}
	return nil, false
}

func (v *VariableFamily) IsBool(value concept.Variable) (*Bool, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableBoolType {
		bool, yes := value.(*Bool)
		return bool, yes
	}
	return nil, false
}

func (v *VariableFamily) IsFunction(value concept.Variable) (*Function, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableFunctionType {
		function, yes := value.(*Function)
		return function, yes
	}
	return nil, false
}

func (v *VariableFamily) IsSystemFunction(value concept.Variable) (*SystemFunction, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableSystemFunctionType {
		function, yes := value.(*SystemFunction)
		return function, yes
	}
	return nil, false
}

func (v *VariableFamily) IsDefineFunction(value concept.Variable) (*DefineFunction, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableDefineFunctionType {
		function, yes := value.(*DefineFunction)
		return function, yes
	}
	return nil, false
}

func (v *VariableFamily) IsDelayFunction(value concept.Variable) (*DelayFunction, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableDelayFunctionType {
		function, yes := value.(*DelayFunction)
		return function, yes
	}
	return nil, false
}

func (v *VariableFamily) IsValueLanguageFunction(value concept.Variable) (*ValueLanguageFunction, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableValueLanguageFunctionType {
		function, yes := value.(*ValueLanguageFunction)
		return function, yes
	}
	return nil, false
}

func (v *VariableFamily) IsParam(value concept.Variable) (*Param, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableParamType {
		param, yes := value.(*Param)
		return param, yes
	}
	return nil, false
}

func (v *VariableFamily) IsArray(value concept.Variable) (*Array, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableArrayType {
		array, yes := value.(*Array)
		return array, yes
	}
	return nil, false
}

func (v *VariableFamily) IsClass(value concept.Variable) (*Class, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableClassType {
		class, yes := value.(*Class)
		return class, yes
	}
	return nil, false
}

func (v *VariableFamily) IsObject(value concept.Variable) (*Object, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableObjectType {
		object, yes := value.(*Object)
		return object, yes
	}
	return nil, false
}

func (v *VariableFamily) IsMappingObject(value concept.Variable) (*MappingObject, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == VariableMappingObjectType {
		object, yes := value.(*MappingObject)
		return object, yes
	}
	return nil, false
}

func (v *VariableFamily) IsStringHome(value concept.Variable) (concept.String, bool) {
	str, yes := v.IsString(value)
	if yes {
		return str, yes
	}

	delayStr, yes := v.IsDelayString(value)
	if yes {
		return delayStr, yes
	}

	return nil, false
}

func (v *VariableFamily) IsFunctionHome(value concept.Variable) (concept.Function, bool) {
	function, yes := v.IsFunction(value)
	if yes {
		return function, yes
	}

	systemFunction, yes := v.IsSystemFunction(value)
	if yes {
		return systemFunction, yes
	}

	delayFunction, yes := v.IsDelayFunction(value)
	if yes {
		return delayFunction, yes
	}

	defineFunction, yes := v.IsDefineFunction(value)
	if yes {
		return defineFunction, yes
	}

	valueLanguageFunction, yes := v.IsValueLanguageFunction(value)
	if yes {
		return valueLanguageFunction, yes
	}

	return nil, false
}

func (v *VariableFamily) IsObjectHome(value concept.Variable) (concept.Object, bool) {
	object, yes := v.IsObject(value)
	if yes {
		return object, yes
	}
	mappingObject, yes := v.IsMappingObject(value)
	if yes {
		return mappingObject, yes
	}
	return nil, false
}

func newVariableFamily() *VariableFamily {
	return &VariableFamily{}
}

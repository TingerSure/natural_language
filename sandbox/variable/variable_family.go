package variable

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type VariableFamily struct {
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
		funcs, yes := value.(*Function)
		return funcs, yes
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

func newVariableFamily() *VariableFamily {
	return &VariableFamily{}
}

var (
	VariableFamilyInstance *VariableFamily = newVariableFamily()
)

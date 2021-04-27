package adaptor

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type AdaptorVariableParam struct {
	NullCreator      func() concept.Null
	ExceptionCreator func(string, string) concept.Exception
}

type AdaptorVariable struct {
	fields *concept.Mapping
	param  *AdaptorVariableParam
}

func (o *AdaptorVariable) GetSource() concept.Variable {
	return nil
}

func (o *AdaptorVariable) GetClass() concept.Class {
	return nil
}

func (o *AdaptorVariable) GetMapping() *concept.Mapping {
	return nil
}

func (o *AdaptorVariable) IsFunction() bool {
	return false
}

func (o *AdaptorVariable) IsNull() bool {
	return false
}

func (o *AdaptorVariable) CallAdaptor(specimen concept.String, param concept.Param, this concept.Variable) (concept.Param, concept.Exception) {
	value, exception := o.GetField(specimen)
	if nl_interface.IsNil(exception) {
		return nil, exception
	}
	if !value.IsFunction() {
		return nil, o.param.ExceptionCreator("runtime error", fmt.Sprintf("There is no function called \"%v\" to be called here.", specimen.ToString("")))
	}
	return value.(concept.Function).Exec(param, this)
}

func (o *AdaptorVariable) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	o.initFields()
	o.fields.Set(specimen, value)
	return nil
}

func (o *AdaptorVariable) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	o.initFields()
	return o.fields.Get(specimen).(concept.Variable), nil
}

func (o *AdaptorVariable) HasField(specimen concept.String) bool {
	o.initFields()
	return o.fields.Has(specimen)
}

func (a *AdaptorVariable) ToString(prefix string) string {
	if a.fields == nil || 0 == a.fields.Size() {
		return "{}"
	}

	subPrefix := fmt.Sprintf("%v\t", prefix)

	paramsToString := make([]string, 0, a.fields.Size())
	if a.fields != nil {
		a.fields.Iterate(func(key concept.String, value interface{}) bool {
			paramsToString = append(paramsToString, fmt.Sprintf("%v%v : %v", subPrefix, key.ToString(subPrefix), value.(concept.ToString).ToString(subPrefix)))
			return false
		})
	}
	return fmt.Sprintf("{\n%v\n%v}", strings.Join(paramsToString, ",\n"), prefix)
}

func (o *AdaptorVariable) initFields() {
	if o.fields == nil {
		o.fields = concept.NewMapping(&concept.MappingParam{
			AutoInit:   true,
			EmptyValue: o.param.NullCreator(),
		})
	}
}

func NewAdaptorVariable(param *AdaptorVariableParam) *AdaptorVariable {
	return &AdaptorVariable{
		param: param,
	}
}

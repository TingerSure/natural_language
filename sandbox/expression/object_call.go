package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type ObjectCall struct {
	*adaptor.ExpressionIndex
	key    string
	object concept.Index
	param  concept.Index
}

var (
	objectCallDefaultParam = callDefaultParam
)

func (a *ObjectCall) ToString(prefix string) string {
	return fmt.Sprintf("%v.%v(%v)", a.object.ToString(prefix), a.key, a.param.ToString(prefix))
}

func (a *ObjectCall) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preObject, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	object, yesObject := variable.VariableFamilyInstance.IsObject(preObject)
	if !yesObject {
		return nil, interrupt.NewException("type error", "Only Object can be called in ObjectCall")
	}

	method, suspend := object.GetMethod(a.key)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	preParam, suspend := a.param.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	yesParam := false
	param, yesParam := variable.VariableFamilyInstance.IsParam(preParam)
	if !yesParam {
		return nil, interrupt.NewException("type error", "Only Param can are passed to a Function")
	}

	return method.Exec(param, object)
}

func NewObjectCall(object concept.Index, key string, param concept.Index) *ObjectCall {
	if nl_interface.IsNil(param) {
		param = objectCallDefaultParam
	}
	back := &ObjectCall{
		key:    key,
		object: object,
		param:  param,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

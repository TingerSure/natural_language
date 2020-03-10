package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ObjectGet struct {
	*adaptor.ExpressionIndex
	key    concept.KeySpecimen
	object concept.Index
}

func (a *ObjectGet) ToString(prefix string) string {
	return fmt.Sprintf("%v.%v", a.object.ToString(prefix), a.key.ToString(prefix))
}

func (a *ObjectGet) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preObject, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, yesObject := variable.VariableFamilyInstance.IsObject(preObject)
	if !yesObject {
		return nil, interrupt.NewException("type error", "Only Object can be get in ObjectGet")
	}

	value, suspend := object.GetField(a.key)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	if nl_interface.IsNil(value) {
		return nil, interrupt.NewException("type error", fmt.Sprintf("%v is empty.", a.ToString("")))
	}
	return value, nil
}

func NewObjectGet(object concept.Index, key concept.KeySpecimen) *ObjectGet {
	back := &ObjectGet{
		key:    key,
		object: object,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

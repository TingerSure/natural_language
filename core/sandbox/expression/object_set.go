package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ObjectSet struct {
	*adaptor.ExpressionIndex
	key    concept.String
	object concept.Index
	value  concept.Index
}

func (a *ObjectSet) ToString(prefix string) string {
	return fmt.Sprintf("%v.%v = %v", a.object.ToString(prefix), a.key.ToString(prefix), a.value.ToString(prefix))
}

func (a *ObjectSet) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preObject, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, yesObject := variable.VariableFamilyInstance.IsObject(preObject)
	if !yesObject {
		return nil, interrupt.NewException("type error", "Only Object can be get in ObjectSet")
	}

	preValue, suspend := a.value.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	suspend = object.SetField(a.key, preValue)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	return preObject, nil
}

func NewObjectSet(object concept.Index, key concept.String, value concept.Index) *ObjectSet {
	back := &ObjectSet{
		key:    key,
		object: object,
		value:  value,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

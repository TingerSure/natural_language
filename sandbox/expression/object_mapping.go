package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type ObjectMapping struct {
	*adaptor.ExpressionIndex
	object concept.Index
	class  string
	alias  string
}

func (a *ObjectMapping) ToString(prefix string) string {
	if a.alias == "" {
		return fmt.Sprintf("%v<%v>", a.object.ToString(prefix), a.class)
	}
	return fmt.Sprintf("%v<%v(%v)>", a.object.ToString(prefix), a.class, a.alias)
}

func (a *ObjectMapping) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {

	preObject, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, yesObject := variable.VariableFamilyInstance.IsObject(preObject)
	if !yesObject {
		return nil, interrupt.NewException("type error", "Only Object can be get in ObjectMapping")
	}

	mappingObject, exception := variable.NewMappingObject(object, a.class, a.alias)

	if nl_interface.IsNil(exception) {
		return nil, exception
	}
	return mappingObject, nil
}

func NewObjectMapping(object concept.Index, class string, alias string) *ObjectMapping {
	back := &ObjectMapping{
		object: object,
		class:  class,
		alias:  alias,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

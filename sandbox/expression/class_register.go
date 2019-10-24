package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
	"strings"
)

type ClassRegister struct {
	*adaptor.ExpressionIndex
	object  concept.Index
	class   concept.Index
	mapping map[string]string
	alias   string
}

func (a *ClassRegister) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)
	items := []string{}
	for key, value := range a.mapping {
		items = append(items, fmt.Sprintf("%v%v : %v", subprefix, key, value))
	}
	return fmt.Sprintf("%v <- %v <%v> {\n%v\n%v}",
		a.object.ToString(prefix),
		a.class.ToString(prefix),
		a.alias,
		strings.Join(items, ",\n"),
		prefix,
	)
}

func (a *ClassRegister) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	preObject, suspend := a.object.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	object, yesObject := variable.VariableFamilyInstance.IsObject(preObject)
	if !yesObject {
		return nil, interrupt.NewException("type error", "Only Object can be use in ClassRegister")
	}

	preClass, suspend := a.class.Get(space)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	class, yesClass := variable.VariableFamilyInstance.IsClass(preClass)
	if !yesClass {
		return nil, interrupt.NewException("type error", "Only Class can be use in ClassRegister")
	}

	for classKey, _ := range class.AllFields() {
		objectKey := a.mapping[classKey]
		if objectKey == "" {
			return nil, interrupt.NewException("type error", fmt.Sprintf("ClassKey \"%v\" is not found in mapping.", classKey))
		}
		if !object.HasField(objectKey) {
			return nil, interrupt.NewException("type error", fmt.Sprintf("ObjectKey \"%v\" is not found in object.", objectKey))
		}
	}

	return object, object.AddClass(class, a.alias, a.mapping)
}

func NewClassRegister(object concept.Index, class concept.Index, mapping map[string]string, alias string) *ClassRegister {
	back := &ClassRegister{
		object:  object,
		class:   class,
		mapping: mapping,
		alias:   alias,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

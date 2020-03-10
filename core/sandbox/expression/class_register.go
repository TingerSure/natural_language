package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type ClassRegister struct {
	*adaptor.ExpressionIndex
	object  concept.Index
	class   concept.Index
	mapping map[concept.KeySpecimen]concept.KeySpecimen
	alias   string
}

func (a *ClassRegister) ToString(prefix string) string {
	subprefix := fmt.Sprintf("%v\t", prefix)
	items := []string{}
	for key, value := range a.mapping {
		items = append(items, fmt.Sprintf("%v%v : %v", subprefix, key.ToString(subprefix), value.ToString(subprefix)))
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

	if object.CheckMapping(class, a.mapping) {
		return nil, interrupt.NewException("type error", fmt.Sprintf("Class \"%v\" cannot register to Object \"%v\".", a.class.ToString(""), a.object.ToString("")))
	}

	return object, object.AddClass(class, a.alias, a.mapping)
}

func NewClassRegister(object concept.Index, class concept.Index, mapping map[concept.KeySpecimen]concept.KeySpecimen, alias string) *ClassRegister {
	back := &ClassRegister{
		object:  object,
		class:   class,
		mapping: mapping,
		alias:   alias,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}

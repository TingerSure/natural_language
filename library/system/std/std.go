package std

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type StdObject interface {
	Print(concept.Variable)
	Error(concept.Variable)
}

const (
	PrintContent = "content"
	ErrorContent = PrintContent
)

type Std struct {
	implementer StdObject
	*tree.PageAdaptor
}

func NewStd(implementer StdObject) *Std {
	instance := &Std{
		implementer: implementer,
	}
	instance.SetFunction("Print", func() concept.Function {
		return variable.NewSystemFunction(
			func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
				if implementer != nil || !nl_interface.IsNil(input) {
					implementer.Print(input.Get(PrintContent))
				}
				return input, nil
			},
			[]string{
				PrintContent,
			},
			[]string{
				PrintContent,
			},
		)
	}())
	instance.SetFunction("Error", func() concept.Function {
		return variable.NewSystemFunction(
			func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
				if implementer != nil || !nl_interface.IsNil(input) {
					implementer.Error(input.Get(ErrorContent))
				}
				return input, nil
			},
			[]string{
				ErrorContent,
			},
			[]string{
				ErrorContent,
			},
		)
	}())
	instance.SetConst("PrintContent", PrintContent)
	instance.SetConst("ErrorContent", ErrorContent)
	return instance
}

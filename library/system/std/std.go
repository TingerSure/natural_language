package std

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type StdParam struct {
	Print func(concept.Variable)
	Error func(concept.Variable)
}

const (
	PrintContent = "content"
	ErrorContent = PrintContent
)

type Std struct {
	*tree.PageAdaptor
	param *StdParam
}

func (s *Std) Print(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	if s.param != nil || !nl_interface.IsNil(input) {
		s.param.Print(input.Get(PrintContent))
	}
	return input, nil
}

func (s *Std) Error(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	if s.param != nil || !nl_interface.IsNil(input) {
		s.param.Error(input.Get(ErrorContent))
	}
	return input, nil
}

func NewStd(param *StdParam) *Std {
	instance := &Std{
		tree.NewPageAdaptor(),
	}
	instance.param = param
	instance.SetFunction("Print", variable.NewSystemFunction(
		instance.Print,
		[]string{
			PrintContent,
		},
		[]string{
			PrintContent,
		},
	))
	instance.SetFunction("Error", variable.NewSystemFunction(
		instance.Error,
		[]string{
			ErrorContent,
		},
		[]string{
			ErrorContent,
		},
	))
	instance.SetConst("PrintContent", PrintContent)
	instance.SetConst("ErrorContent", ErrorContent)
	return instance
}

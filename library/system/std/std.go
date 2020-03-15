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
	*tree.Page
	param        *StdParam
	PrintContent concept.String
	ErrorContent concept.String
}

func (s *Std) Print(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	if s.param != nil || !nl_interface.IsNil(input) {
		s.param.Print(input.Get(s.PrintContent))
	}
	return input, nil
}

func (s *Std) Error(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	if s.param != nil || !nl_interface.IsNil(input) {
		s.param.Error(input.Get(s.ErrorContent))
	}
	return input, nil
}

func NewStd(param *StdParam) *Std {
	instance := &Std{
		param:        param,
		Page:         tree.NewPage(),
		PrintContent: variable.NewString(PrintContent),
		ErrorContent: variable.NewString(ErrorContent),
	}
	instance.SetFunction(variable.NewString("Print"), variable.NewSystemFunction(
		instance.Print,
		[]concept.String{
			instance.PrintContent,
		},
		[]concept.String{
			instance.PrintContent,
		},
	))
	instance.SetFunction(variable.NewString("Error"), variable.NewSystemFunction(
		instance.Error,
		[]concept.String{
			instance.ErrorContent,
		},
		[]concept.String{
			instance.ErrorContent,
		},
	))
	instance.SetConst(variable.NewString("PrintContent"), instance.PrintContent)
	instance.SetConst(variable.NewString("ErrorContent"), instance.ErrorContent)
	return instance
}

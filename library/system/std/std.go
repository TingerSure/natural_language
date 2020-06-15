package std

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
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
	tree.Page
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

func NewStd(libs *runtime.LibraryManager, param *StdParam) *Std {
	instance := &Std{
		param:        param,
		Page:         tree.NewPageAdaptor(libs.Sandbox),
		PrintContent: libs.Sandbox.Variable.String.New(PrintContent),
		ErrorContent: libs.Sandbox.Variable.String.New(ErrorContent),
	}
	instance.SetFunction(libs.Sandbox.Variable.String.New("Print"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("Print"),
		instance.Print,
		[]concept.String{
			instance.PrintContent,
		},
		[]concept.String{
			instance.PrintContent,
		},
	))
	instance.SetFunction(libs.Sandbox.Variable.String.New("Error"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("Error"),
		instance.Error,
		[]concept.String{
			instance.ErrorContent,
		},
		[]concept.String{
			instance.ErrorContent,
		},
	))
	instance.SetConst(libs.Sandbox.Variable.String.New("PrintContent"), instance.PrintContent)
	instance.SetConst(libs.Sandbox.Variable.String.New("ErrorContent"), instance.ErrorContent)
	return instance
}

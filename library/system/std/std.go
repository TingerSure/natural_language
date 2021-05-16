package std

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
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
	concept.Page
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

func NewStd(libs *tree.LibraryManager, param *StdParam) *Std {
	instance := &Std{
		param:        param,
		Page:         libs.Sandbox.Variable.Page.New(),
		PrintContent: libs.Sandbox.Variable.String.New(PrintContent),
		ErrorContent: libs.Sandbox.Variable.String.New(ErrorContent),
	}
	instance.SetPublic(libs.Sandbox.Variable.String.New("Print"), libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.SystemFunction.New(
		instance.Print,
		func(input concept.Param, _ concept.Object) concept.Param {
			return input
		},
		[]concept.String{
			instance.PrintContent,
		},
		[]concept.String{
			instance.PrintContent,
		},
	)))
	instance.SetPublic(libs.Sandbox.Variable.String.New("Error"), libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.SystemFunction.New(
		instance.Error,
		func(input concept.Param, _ concept.Object) concept.Param {
			return input
		},
		[]concept.String{
			instance.ErrorContent,
		},
		[]concept.String{
			instance.ErrorContent,
		},
	)))
	instance.SetPublic(libs.Sandbox.Variable.String.New("PrintContent"), libs.Sandbox.Index.ConstIndex.New(instance.PrintContent))
	instance.SetPublic(libs.Sandbox.Variable.String.New("ErrorContent"), libs.Sandbox.Index.ConstIndex.New(instance.ErrorContent))
	return instance
}

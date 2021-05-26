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
	libs         *tree.LibraryManager
	param        *StdParam
	PrintContent concept.String
	ErrorContent concept.String
}

func (s *Std) Print(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	if s.param != nil || !nl_interface.IsNil(input) {
		s.param.Print(input.Get(s.PrintContent))
	}
	return s.libs.Sandbox.Variable.Param.New(), nil
}

func (s *Std) Error(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	if s.param != nil || !nl_interface.IsNil(input) {
		s.param.Error(input.Get(s.ErrorContent))
	}
	return s.libs.Sandbox.Variable.Param.New(), nil
}

func NewStd(libs *tree.LibraryManager, param *StdParam) *Std {
	instance := &Std{
		libs:         libs,
		param:        param,
		Page:         libs.Sandbox.Variable.Page.New(),
		PrintContent: libs.Sandbox.Variable.String.New(PrintContent),
		ErrorContent: libs.Sandbox.Variable.String.New(ErrorContent),
	}
	instance.SetPublic(libs.Sandbox.Variable.String.New("Print"), libs.Sandbox.Index.PublicIndex.New("Print", libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.SystemFunction.New(
		instance.Print,
		func(input concept.Param, _ concept.Variable) concept.Param {
			return input
		},
		[]concept.String{
			instance.PrintContent,
		},
		[]concept.String{},
	))))
	instance.SetPublic(libs.Sandbox.Variable.String.New("Error"), libs.Sandbox.Index.PublicIndex.New("Error", libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.SystemFunction.New(
		instance.Error,
		func(input concept.Param, _ concept.Variable) concept.Param {
			return input
		},
		[]concept.String{
			instance.ErrorContent,
		},
		[]concept.String{},
	))))
	return instance
}

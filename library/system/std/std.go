package std

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type StdParam struct {
	Print func(concept.Variable)
	Error func(concept.Variable)
	Log   func(concept.Variable)
}

const (
	PrintContent = "content"
	ErrorContent = PrintContent
	LogContent   = PrintContent
)

type Std struct {
	concept.Page
	libs         *tree.LibraryManager
	param        *StdParam
	PrintContent concept.String
	ErrorContent concept.String
	LogContent   concept.String
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

func (s *Std) Log(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
	if s.param != nil || !nl_interface.IsNil(input) {
		s.param.Log(input.Get(s.LogContent))
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
		LogContent:   libs.Sandbox.Variable.String.New(LogContent),
	}
	instance.SetPublic(libs.Sandbox.Variable.String.New("print"), libs.Sandbox.Index.PublicIndex.New("print", libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.SystemFunction.New(
		instance.Print,
		func(input concept.Param, _ concept.Variable) concept.Param {
			return libs.Sandbox.Variable.Param.New()
		},
		[]concept.String{
			instance.PrintContent,
		},
		[]concept.String{},
	))))
	instance.SetPublic(libs.Sandbox.Variable.String.New("error"), libs.Sandbox.Index.PublicIndex.New("error", libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.SystemFunction.New(
		instance.Error,
		func(input concept.Param, _ concept.Variable) concept.Param {
			return libs.Sandbox.Variable.Param.New()
		},
		[]concept.String{
			instance.ErrorContent,
		},
		[]concept.String{},
	))))
	instance.SetPublic(libs.Sandbox.Variable.String.New("log"), libs.Sandbox.Index.PublicIndex.New("log", libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.SystemFunction.New(
		instance.Log,
		func(input concept.Param, _ concept.Variable) concept.Param {
			return libs.Sandbox.Variable.Param.New()
		},
		[]concept.String{
			instance.LogContent,
		},
		[]concept.String{},
	))))
	return instance
}

package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

/*
type Phrase interface {
--	Copy() Phrase
++	Size() int
++	ContentSize() int
++	Types() (string, concept.Exception)
--	SetTypes(types string)
++	GetContent() string
++	GetChild(index int) Phrase
--	SetChild(index int, child Phrase) Phrase
++	ToString() string
++	ToContent() string
--	ToStringOffset(index int) string
++	Index() (concept.Function, concept.Exception)
++	From() string
++	HasPriority() bool
--	DependencyCheckValue() Phrase
}
*/
func newPhrase(libs *tree.LibraryManager, source tree.Phrase) concept.Object {
	instance := libs.Sandbox.Variable.Object.New()
	instance.SetField(libs.Sandbox.Variable.String.New("size"), newPhraseSize(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("contentSize"), newPhraseContentSize(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("types"), newPhraseTypes(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("getContent"), newPhraseGetContent(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("getChild"), newPhraseGetChild(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("toString"), newPhraseToString(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("toContent"), newPhraseToContent(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("index"), newPhraseIndex(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("from"), newPhraseFrom(libs, source))
	instance.SetField(libs.Sandbox.Variable.String.New("hasPriority"), newPhraseHasPriority(libs, source))
	return instance
}

func newPhraseHasPriority(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	valueParam := libs.Sandbox.Variable.String.New("value")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			output := libs.Sandbox.Variable.Param.New()
			output.Set(valueParam, libs.Sandbox.Variable.Bool.New(source.HasPriority()))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{valueParam},
	)
}

func newPhraseFrom(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	fromParam := libs.Sandbox.Variable.String.New("from")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			output := libs.Sandbox.Variable.Param.New()
			output.Set(fromParam, libs.Sandbox.Variable.String.New(source.From()))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{fromParam},
	)
}

func newPhraseIndex(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	indexParam := libs.Sandbox.Variable.String.New("index")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			funcs, exception := source.Index()
			if !nl_interface.IsNil(exception) {
				return nil, exception
			}
			output := libs.Sandbox.Variable.Param.New()
			output.Set(indexParam, funcs)
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{indexParam},
	)
}

func newPhraseToContent(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	contentParam := libs.Sandbox.Variable.String.New("content")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			output := libs.Sandbox.Variable.Param.New()
			output.Set(contentParam, libs.Sandbox.Variable.String.New(source.ToContent()))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{contentParam},
	)
}

func newPhraseToString(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	valueParam := libs.Sandbox.Variable.String.New("value")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			output := libs.Sandbox.Variable.Param.New()
			output.Set(valueParam, libs.Sandbox.Variable.String.New(source.ToString()))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{valueParam},
	)
}

func newPhraseGetChild(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	indexParam := libs.Sandbox.Variable.String.New("index")
	childParam := libs.Sandbox.Variable.String.New("child")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			indexPre := input.Get(indexParam)
			index, yes := variable.VariableFamilyInstance.IsNumber(indexPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param index is not a number: %v", indexPre.ToString("")))
			}
			child := source.GetChild(int(index.Value()))
			output := libs.Sandbox.Variable.Param.New()
			output.Set(childParam, newPhrase(libs, child))
			return output, nil
		},
		nil,
		[]concept.String{indexParam},
		[]concept.String{childParam},
	)
}

func newPhraseGetContent(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	contentParam := libs.Sandbox.Variable.String.New("content")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			output := libs.Sandbox.Variable.Param.New()
			output.Set(contentParam, libs.Sandbox.Variable.String.New(source.GetContent()))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{contentParam},
	)
}

func newPhraseTypes(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	typesParam := libs.Sandbox.Variable.String.New("types")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			types, exception := source.Types()
			if !nl_interface.IsNil(exception) {
				return nil, exception
			}
			output := libs.Sandbox.Variable.Param.New()
			output.Set(typesParam, libs.Sandbox.Variable.String.New(types))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{typesParam},
	)
}

func newPhraseContentSize(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	sizeParam := libs.Sandbox.Variable.String.New("size")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			output := libs.Sandbox.Variable.Param.New()
			output.Set(sizeParam, libs.Sandbox.Variable.Number.New(float64(source.ContentSize())))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{sizeParam},
	)
}

func newPhraseSize(libs *tree.LibraryManager, source tree.Phrase) concept.Function {
	sizeParam := libs.Sandbox.Variable.String.New("size")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			output := libs.Sandbox.Variable.Param.New()
			output.Set(sizeParam, libs.Sandbox.Variable.Number.New(float64(source.Size())))
			return output, nil
		},
		nil,
		[]concept.String{},
		[]concept.String{sizeParam},
	)
}

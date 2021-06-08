package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newAddTypes(libs *tree.LibraryManager) concept.Function {
	nameParam := libs.Sandbox.Variable.String.New("name")
	instance := libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			namePre := input.Get(nameParam)
			name, yes := variable.VariableFamilyInstance.IsStringHome(namePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param name is not a string: %v", namePre.ToString("")))
			}
			libs.Types.AddTypes(tree.NewPhraseType(&tree.PhraseTypeParam{
				Name: name.Value(),
			}))
			return libs.Sandbox.Variable.Param.New(), nil
		},
		func(input concept.Param, _ concept.Variable) concept.Param {
			return libs.Sandbox.Variable.Param.New()
		},
		[]concept.String{
			nameParam,
		},
		[]concept.String{},
	)
	return instance
}

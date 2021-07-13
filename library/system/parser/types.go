package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newAddTypes(libs *tree.LibraryManager) concept.Function {
	nameParam := libs.Sandbox.Variable.String.New("name")
	parentsParam := libs.Sandbox.Variable.String.New("parents")
	instance := libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			namePre := input.Get(nameParam)
			name, yes := variable.VariableFamilyInstance.IsStringHome(namePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param name is not a string: %v", namePre.ToString("")))
			}
			typeParents := []*tree.PhraseTypeParent{}
			parentsPre := input.Get(parentsParam)
			if !parentsPre.IsNull() {
				parents, yes := variable.VariableFamilyInstance.IsArray(parentsPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param parents is not a string: %v", parentsPre.ToString("")))
				}
				for index := 0; index < parents.Length(); index++ {
					parentPre, exception := parents.Get(index)
					if !nl_interface.IsNil(exception) {
						return nil, exception
					}
					parent, yes := variable.VariableFamilyInstance.IsStringHome(parentPre)
					if !yes {
						return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param parents[%v] is not a string: %v", index, parentPre.ToString("")))
					}
					typeParents = append(typeParents, &tree.PhraseTypeParent{
						Types: parent.Value(),
					})
				}
			}
			libs.Types.AddTypes(tree.NewPhraseType(&tree.PhraseTypeParam{
				Name:    name.Value(),
				Parents: typeParents,
			}))
			return libs.Sandbox.Variable.Param.New(), nil
		},
		[]concept.String{
			nameParam,
			parentsParam,
		},
		[]concept.String{},
	)
	return instance
}

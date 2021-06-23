package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newAddStructRule(libs *tree.LibraryManager) concept.Function {
	nameParam := libs.Sandbox.Variable.String.New("name")
	contentTypesParam := libs.Sandbox.Variable.String.New("contentTypes")
	typesParam := libs.Sandbox.Variable.String.New("types")
	dynamicTypesParam := libs.Sandbox.Variable.String.New("dynamicTypes")
	createParam := libs.Sandbox.Variable.String.New("create")

	instance := libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			namePre := input.Get(nameParam)
			name, yes := variable.VariableFamilyInstance.IsStringHome(namePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param name is not a string: %v", namePre.ToString("")))
			}
			typesPre := input.Get(typesParam)
			var types concept.String = libs.Sandbox.Variable.String.New("")
			if !typesPre.IsNull() {
				types, yes = variable.VariableFamilyInstance.IsStringHome(typesPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param types is not a string: %v", typesPre.ToString("")))
				}
			}
			listTypes := []string{}
			contentTypesPre := input.Get(contentTypesParam)
			contentTypes, yes := variable.VariableFamilyInstance.IsArray(contentTypesPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param contentTypes is not a string: %v", contentTypesPre.ToString("")))
			}
			for index := 0; index < contentTypes.Length(); index++ {
				contentTypePre, exception := contentTypes.Get(index)
				if !nl_interface.IsNil(exception) {
					return nil, exception
				}
				contentType, yes := variable.VariableFamilyInstance.IsStringHome(contentTypePre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param contentTypes[%v] is not a string: %v", index, contentTypePre.ToString("")))
				}
				listTypes = append(listTypes, contentType.Value())
			}
			createPre := input.Get(createParam)
			create, yes := variable.VariableFamilyInstance.IsFunctionHome(createPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param create is not a function: %v", createPre.ToString("")))
			}
			dynamicTypesPre := input.Get(dynamicTypesParam)
			var dynamicTypesFuncs func(phrase []tree.Phrase) (string, concept.Exception) = nil
			if !dynamicTypesPre.IsNull() {
				dynamicTypes, yes := variable.VariableFamilyInstance.IsFunctionHome(dynamicTypesPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param dynamicTypes is not a function: %v", dynamicTypesPre.ToString("")))
				}
				dynamicTypesFuncs = func(phrases []tree.Phrase) (string, concept.Exception) {
					line := tree.NewLine("[dynamic_type]", "")
					contents := libs.Sandbox.Variable.Array.New()
					for _, phrase := range phrases {
						content, exception := phrase.Types()
						if !nl_interface.IsNil(exception) {
							return "", exception.AddExceptionLine(line)
						}
						contents.Append(libs.Sandbox.Variable.String.New(content))
					}
					param := libs.Sandbox.Variable.Param.New()
					param.SetOriginal("contents", contents)
					output, exception := dynamicTypes.Exec(param, nil)
					if !nl_interface.IsNil(exception) {
						return "", exception.AddExceptionLine(line)
					}
					outputTypesPre := output.GetOriginal("types")
					outputTypes, yes := variable.VariableFamilyInstance.IsStringHome(outputTypesPre)
					if !yes {
						return "", libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Return types is not a string: %v", outputTypesPre.ToString(""))).AddExceptionLine(line)
					}
					return outputTypes.Value(), nil
				}
			}
			libs.Structs.AddRule(tree.NewStructRule(&tree.StructRuleParam{
				From:  name.Value(),
				Types: listTypes,
				Create: func() tree.Phrase {
					return tree.NewPhraseStruct(&tree.PhraseStructParam{
						Types:        types.Value(),
						From:         name.Value(),
						Size:         len(listTypes),
						DynamicTypes: dynamicTypesFuncs,
						Index: func(phrases []tree.Phrase) (concept.Function, concept.Exception) {
							line := tree.NewLine(fmt.Sprintf("[struct_parse]:%v", name.Value()), "")
							contents := libs.Sandbox.Variable.Array.New()
							for _, phrase := range phrases {
								content, exception := phrase.Index()
								if !nl_interface.IsNil(exception) {
									return nil, exception.AddExceptionLine(line)
								}
								contents.Append(content)
							}
							param := libs.Sandbox.Variable.Param.New()
							param.SetOriginal("contents", contents)
							output, exception := create.Exec(param, nil)
							if !nl_interface.IsNil(exception) {
								return nil, exception.AddExceptionLine(line)
							}
							pipe, yes := variable.VariableFamilyInstance.IsFunctionHome(output.GetOriginal("pipe"))
							if !yes {
								return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", "struct parse pipe must be defined as 'function()value'.").AddExceptionLine(line)
							}
							return pipe, nil
						},
					})
				},
			}))
			return libs.Sandbox.Variable.Param.New(), nil
		},
		func(input concept.Param, _ concept.Variable) concept.Param {
			return libs.Sandbox.Variable.Param.New()
		},
		[]concept.String{
			nameParam,
			contentTypesParam,
			typesParam,
			dynamicTypesParam,
			createParam,
		},
		[]concept.String{},
	)
	return instance
}

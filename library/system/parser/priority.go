package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newAddPriorityRule(libs *tree.LibraryManager) concept.Function {
	nameParam := libs.Sandbox.Variable.String.New("name")
	matchParam := libs.Sandbox.Variable.String.New("match")
	chooserParam := libs.Sandbox.Variable.String.New("chooser")

	instance := libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			namePre := input.Get(nameParam)
			name, yes := variable.VariableFamilyInstance.IsStringHome(namePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param name is not a string: %v", namePre.ToString("")))
			}
			matchPre := input.Get(matchParam)
			match, yes := variable.VariableFamilyInstance.IsFunctionHome(matchPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param match is not a function: %v", matchPre.ToString("")))
			}

			chooserPre := input.Get(chooserParam)
			chooser, yes := variable.VariableFamilyInstance.IsFunctionHome(chooserPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param chooser is not a function: %v", chooserPre.ToString("")))
			}
			libs.Priorities.AddRule(tree.NewPriorityRule(&tree.PriorityRuleParam{
				From: name.Value(),
				Match: func(left tree.Phrase, right tree.Phrase) (bool, concept.Exception) {
					line := tree.NewLine(fmt.Sprintf("[priority_match]: %v", name.Value()), "")
					param := libs.Sandbox.Variable.Param.New()
					param.SetOriginal("left", newPhrase(libs, left))
					param.SetOriginal("right", newPhrase(libs, right))
					output, exception := match.Exec(param, nil)
					if !nl_interface.IsNil(exception) {
						return false, exception.AddExceptionLine(line)
					}
					yes, ok := variable.VariableFamilyInstance.IsBool(output.GetOriginal("yes"))
					if !ok {
						return false, libs.Sandbox.Variable.Exception.NewOriginal("type error", "Return yes is not a bool.").AddExceptionLine(line)
					}
					return yes.Value(), nil
				},
				Chooser: func(left tree.Phrase, right tree.Phrase) (*tree.PriorityResult, concept.Exception) {
					line := tree.NewLine(fmt.Sprintf("[priority_choose]: %v", name.Value()), "")
					param := libs.Sandbox.Variable.Param.New()
					param.SetOriginal("left", newPhrase(libs, left))
					param.SetOriginal("right", newPhrase(libs, right))
					output, exception := chooser.Exec(param, nil)
					if !nl_interface.IsNil(exception) {
						return nil, exception.AddExceptionLine(line)
					}
					result, yes := variable.VariableFamilyInstance.IsNumber(output.GetOriginal("result"))
					if !yes {
						return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", "Return result is not a number.").AddExceptionLine(line)
					}
					localResult := tree.NewPriorityResult(int(result.Value()))
					abandonsPre := output.GetOriginal("abandons")
					if !abandonsPre.IsNull() {
						abandons, yes := variable.VariableFamilyInstance.IsArray(abandonsPre)
						if !yes {
							return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", "Return abandons is not an array.").AddExceptionLine(line)
						}
						for index := 0; index < abandons.Length(); index++ {
							abandonPre, exception := abandons.Get(index)
							if !nl_interface.IsNil(exception) {
								return nil, exception.AddExceptionLine(line)
							}
							abandon, yes := variable.VariableFamilyInstance.IsNumber(abandonPre)
							if !yes {
								return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Return abandons[%v] is not a number.", index)).AddExceptionLine(line)
							}
							localResult.AddAbandon(int(abandon.Value()))
						}
					}
					return localResult, nil
				},
			}))
			return libs.Sandbox.Variable.Param.New(), nil
		},
		func(input concept.Param, _ concept.Variable) concept.Param {
			return libs.Sandbox.Variable.Param.New()
		},
		[]concept.String{
			nameParam,
			matchParam,
			chooserParam,
		},
		[]concept.String{},
	)
	return instance
}

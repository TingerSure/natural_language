package variable

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"strconv"
)

type Variable struct {
	concept.Page
	libs *tree.LibraryManager
}

func NewVariable(libs *tree.LibraryManager) *Variable {
	instance := &Variable{
		libs: libs,
		Page: libs.Sandbox.Variable.Page.New(),
	}
	instance.SetPublic(
		libs.Sandbox.Variable.String.New("stringToNumber"),
		libs.Sandbox.Index.PublicIndex.New(
			"stringToNumber",
			libs.Sandbox.Index.ConstIndex.New(newStringToNumber(libs)),
		),
	)
	instance.SetPublic(
		libs.Sandbox.Variable.String.New("numberToString"),
		libs.Sandbox.Index.PublicIndex.New(
			"numberToString",
			libs.Sandbox.Index.ConstIndex.New(newNumberToString(libs)),
		),
	)
	return instance
}

func newStringToNumber(libs *tree.LibraryManager) concept.Function {
	fromParam := libs.Sandbox.Variable.String.New("from")
	valueParam := libs.Sandbox.Variable.String.New("value")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
			fromPre := input.Get(fromParam)
			from, yes := variable.VariableFamilyInstance.IsStringHome(fromPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param from is not a string: %v", fromPre.ToString("")))
			}
			value, err := strconv.ParseFloat(from.Value(), 64)
			if err != nil {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("value error", fmt.Sprintf("Param from is not a number string: %v", fromPre.ToString("")))
			}
			output := libs.Sandbox.Variable.Param.New()
			output.Set(valueParam, libs.Sandbox.Variable.Number.New(value))
			return output, nil
		},
		nil,
		[]concept.String{
			fromParam,
		},
		[]concept.String{
			valueParam,
		},
	)
}

func newNumberToString(libs *tree.LibraryManager) concept.Function {
	fromParam := libs.Sandbox.Variable.String.New("from")
	valueParam := libs.Sandbox.Variable.String.New("value")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
			fromPre := input.Get(fromParam)
			from, yes := variable.VariableFamilyInstance.IsNumber(fromPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param from is not a number: %v", fromPre.ToString("")))
			}
			output := libs.Sandbox.Variable.Param.New()
			output.Set(valueParam, libs.Sandbox.Variable.String.New(strconv.FormatFloat(from.Value(), 'E', -1, 64)))
			return output, nil
		},
		nil,
		[]concept.String{
			fromParam,
		},
		[]concept.String{
			valueParam,
		},
	)
}

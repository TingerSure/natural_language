package runtime

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newThrow(libs *tree.LibraryManager) concept.Function {
	nameParam := libs.Sandbox.Variable.String.New("name")
	messageParam := libs.Sandbox.Variable.String.New("message")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			namePre := param.Get(nameParam)
			name, yes := variable.VariableFamilyInstance.IsStringHome(namePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param name is not a string: %v", namePre.ToString("")))
			}
			messagePre := param.Get(messageParam)
			message, yes := variable.VariableFamilyInstance.IsStringHome(messagePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param message is not a string: %v", messagePre.ToString("")))
			}
			return nil, libs.Sandbox.Variable.Exception.New(name, message)
		},
		nil,
		[]concept.String{
			nameParam,
			messageParam,
		},
		[]concept.String{},
	)
}

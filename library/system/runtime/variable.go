package runtime

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

func newVariableCreator(libs *tree.LibraryManager) concept.Object {
	variables := libs.Sandbox.Variable.Object.New()
	variables.SetField(
		libs.Sandbox.Variable.String.New("string"),
		newStringCreator(libs),
	)
	return variables
}

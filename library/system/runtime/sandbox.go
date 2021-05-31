package runtime

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

func newSandboxCreator(libs *tree.LibraryManager) concept.Object {
	sandbox := libs.Sandbox.Variable.Object.New()
	sandbox.SetField(
		libs.Sandbox.Variable.String.New("variable"),
		newVariableCreator(libs),
	)
	return sandbox
}

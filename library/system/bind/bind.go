package bind

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Bind struct {
	concept.Page
	libs *tree.LibraryManager
}

func NewBind(libs *tree.LibraryManager) *Bind {
	instance := &Bind{
		libs: libs,
		Page: libs.Sandbox.Variable.Page.New(),
	}
	instance.SetPublic(
		libs.Sandbox.Variable.String.New("variable"),
		libs.Sandbox.Index.PublicIndex.New(
			"variable",
			libs.Sandbox.Index.ConstIndex.New(newVariableCreator(libs)),
		),
	)

	return instance
}

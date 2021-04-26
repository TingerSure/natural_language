package set

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Set struct {
	concept.Page
}

func NewSet(libs *tree.LibraryManager) *Set {
	instance := &Set{
		Page: libs.Sandbox.Variable.Page.New(),
	}

	instance.SetExport(libs.Sandbox.Variable.String.New("Is"), libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.String.New("Is")))
	instance.SetExport(libs.Sandbox.Variable.String.New("Equal"), libs.Sandbox.Index.ConstIndex.New(libs.Sandbox.Variable.String.New("Equal")))
	return instance
}

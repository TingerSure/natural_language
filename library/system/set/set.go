package set

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Set struct {
	tree.Page
}

func NewSet(libs *tree.LibraryManager) *Set {
	instance := &Set{
		Page: libs.Sandbox.Variable.Object.New(),
	}

	instance.SetField(libs.Sandbox.Variable.String.New("Is"), libs.Sandbox.Variable.String.New("Is"))
	instance.SetField(libs.Sandbox.Variable.String.New("Equal"), libs.Sandbox.Variable.String.New("Equal"))
	return instance
}

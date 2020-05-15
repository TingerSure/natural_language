package set

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type Set struct {
	tree.Page
}

func NewSet(libs *runtime.LibraryManager) *Set {
	instance := &Set{
		Page: tree.NewPageAdaptor(),
	}

	instance.SetConst(variable.NewString("Is"), variable.NewString("Is"))
	instance.SetConst(variable.NewString("Equal"), variable.NewString("Equal"))
	return instance
}

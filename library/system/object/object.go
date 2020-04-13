package object

import (
	"github.com/TingerSure/natural_language/core/tree"
)

type Object struct {
	tree.Page
}

func NewObject(libs *tree.LibraryManager) *Object {
	instance := &Object{
		Page: tree.NewPageAdaptor(),
	}
	initCreate(instance)
	initGetField(instance)
	initGetMethod(instance)
	initHasField(instance)
	initHasMethod(instance)
	initInitField(instance)
	initSetField(instance)
	initSetMethod(instance)
	return instance
}

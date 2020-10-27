package object

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/tree"
)

type Object struct {
	tree.Page
}

func NewObject(libs *runtime.LibraryManager) *Object {
	instance := &Object{
		Page: tree.NewPageAdaptor(libs.Sandbox),
	}
	initCreate(libs, instance)
	initGetField(libs, instance)
	initGetMethod(libs, instance)
	initHasField(libs, instance)
	initHasMethod(libs, instance)
	initInitField(libs, instance)
	initSetField(libs, instance)
	initSetMethod(libs, instance)
	return instance
}

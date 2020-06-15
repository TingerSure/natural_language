package object

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/auto_number"
)

type Object struct {
	tree.Page
	*auto_number.AutoNumber
}

func NewObject(libs *runtime.LibraryManager, autoNumber *auto_number.AutoNumber) *Object {
	instance := &Object{
		Page:       tree.NewPageAdaptor(libs.Sandbox),
		AutoNumber: autoNumber,
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

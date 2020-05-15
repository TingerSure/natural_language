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
		Page:       tree.NewPageAdaptor(),
		AutoNumber: autoNumber,
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

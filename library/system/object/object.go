package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type Object struct {
	tree.Page
	Create    *variable.SystemFunction
	GetField  *variable.SystemFunction
	HasField  *variable.SystemFunction
	GetMethod *variable.SystemFunction
	HasMethod *variable.SystemFunction
	InitField *variable.SystemFunction
	SetField  *variable.SystemFunction
	SetMethod *variable.SystemFunction
}

func NewObject() *Object {
	instance := &Object{
		Page:      tree.NewPageAdaptor(),
		Create:    Create,
		GetField:  GetField,
		HasField:  HasField,
		GetMethod: GetMethod,
		HasMethod: HasMethod,
		InitField: InitField,
		SetField:  SetField,
		SetMethod: SetMethod,
	}

	instance.SetFunction(variable.NewString("Create"), instance.Create)
	instance.SetConst(variable.NewString("CreateContent"), CreateContent)

	instance.SetFunction(variable.NewString("GetField"), instance.GetField)
	instance.SetConst(variable.NewString("GetFieldContent"), GetFieldContent)
	instance.SetConst(variable.NewString("GetFieldKey"), GetFieldKey)
	instance.SetConst(variable.NewString("GetFieldValue"), GetFieldValue)

	instance.SetFunction(variable.NewString("GetMethod"), instance.GetMethod)
	instance.SetConst(variable.NewString("GetMethodContent"), GetMethodContent)
	instance.SetConst(variable.NewString("GetMethodKey"), GetMethodKey)
	instance.SetConst(variable.NewString("GetMethodFunction"), GetMethodFunction)

	instance.SetFunction(variable.NewString("HasField"), instance.HasField)
	instance.SetConst(variable.NewString("HasFieldContent"), HasFieldContent)
	instance.SetConst(variable.NewString("HasFieldKey"), HasFieldKey)
	instance.SetConst(variable.NewString("HasFieldExist"), HasFieldExist)

	instance.SetFunction(variable.NewString("HasMethod"), instance.HasMethod)
	instance.SetConst(variable.NewString("HasMethodContent"), HasMethodContent)
	instance.SetConst(variable.NewString("HasMethodKey"), HasMethodKey)
	instance.SetConst(variable.NewString("HasMethodExist"), HasMethodExist)

	instance.SetFunction(variable.NewString("InitField"), instance.InitField)
	instance.SetConst(variable.NewString("InitFieldContent"), InitFieldContent)
	instance.SetConst(variable.NewString("InitFieldKey"), InitFieldKey)
	instance.SetConst(variable.NewString("InitFieldDefaultValue"), InitFieldDefaultValue)

	instance.SetFunction(variable.NewString("SetField"), instance.SetField)
	instance.SetConst(variable.NewString("SetFieldContent"), SetFieldContent)
	instance.SetConst(variable.NewString("SetFieldKey"), SetFieldKey)
	instance.SetConst(variable.NewString("SetFieldValue"), SetFieldValue)

	instance.SetFunction(variable.NewString("SetMethod"), instance.SetMethod)
	instance.SetConst(variable.NewString("SetMethodContent"), SetMethodContent)
	instance.SetConst(variable.NewString("SetMethodKey"), SetMethodKey)
	instance.SetConst(variable.NewString("SetMethodFunction"), SetMethodFunction)

	return instance
}

package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type Object struct {
	*tree.PageAdaptor
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
		PageAdaptor: tree.NewPageAdaptor(),
		Create:      Create,
		GetField:    GetField,
		HasField:    HasField,
		GetMethod:   GetMethod,
		HasMethod:   HasMethod,
		InitField:   InitField,
		SetField:    SetField,
		SetMethod:   SetMethod,
	}

	instance.SetFunction("Create", instance.Create)
	instance.SetConst("CreateContent", CreateContent)

	instance.SetFunction("GetField", instance.GetField)
	instance.SetConst("GetFieldContent", GetFieldContent)
	instance.SetConst("GetFieldKey", GetFieldKey)
	instance.SetConst("GetFieldValue", GetFieldValue)

	instance.SetFunction("GetMethod", instance.GetMethod)
	instance.SetConst("GetMethodContent", GetMethodContent)
	instance.SetConst("GetMethodKey", GetMethodKey)
	instance.SetConst("GetMethodFunction", GetMethodFunction)

	instance.SetFunction("HasField", instance.HasField)
	instance.SetConst("HasFieldContent", HasFieldContent)
	instance.SetConst("HasFieldKey", HasFieldKey)
	instance.SetConst("HasFieldExist", HasFieldExist)

	instance.SetFunction("HasMethod", instance.HasMethod)
	instance.SetConst("HasMethodContent", HasMethodContent)
	instance.SetConst("HasMethodKey", HasMethodKey)
	instance.SetConst("HasMethodExist", HasMethodExist)

	instance.SetFunction("InitField", instance.InitField)
	instance.SetConst("InitFieldContent", InitFieldContent)
	instance.SetConst("InitFieldKey", InitFieldKey)
	instance.SetConst("InitFieldDefaultValue", InitFieldDefaultValue)

	instance.SetFunction("SetField", instance.SetField)
	instance.SetConst("SetFieldContent", SetFieldContent)
	instance.SetConst("SetFieldKey", SetFieldKey)
	instance.SetConst("SetFieldValue", SetFieldValue)

	instance.SetFunction("SetMethod", instance.SetMethod)
	instance.SetConst("SetMethodContent", SetMethodContent)
	instance.SetConst("SetMethodKey", SetMethodKey)
	instance.SetConst("SetMethodFunction", SetMethodFunction)

	return instance
}

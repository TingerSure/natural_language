package concept

type Class interface {
	Variable

	GetName() string

	SetMethod(key String, action Function)
	GetMethod(key String) Function
	HasMethod(key String) bool
	SizeMethod() int
	IterateMethods(func(key String, value Function) bool) bool

	SetField(key String, defaultField Variable)
	GetField(key String) Variable
	HasField(key String) bool
	SizeField() int
	IterateFields(func(key String, value Variable) bool) bool

	SetStaticMethod(key String, action Function)
	GetStaticMethod(key String) Function
	HasStaticMethod(key String) bool
	SizeStaticMethod() int
	IterateStaticMethods(func(key String, value Function) bool) bool

	SetStaticField(key String, action Variable)
	GetStaticField(key String) Variable
	HasStaticField(key String) bool
	SizeStaticField() int
	IterateStaticFields(func(key String, value Variable) bool) bool
}

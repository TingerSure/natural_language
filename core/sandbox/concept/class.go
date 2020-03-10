package concept

type Class interface {
	Variable

	GetName() string

	SetMethod(key KeySpecimen, action Function)
	GetMethod(key KeySpecimen) Function
	HasMethod(key KeySpecimen) bool
	SizeMethod() int
	IterateMethods(func(key Key, value Function) bool) bool

	SetField(key KeySpecimen, defaultField Variable)
	GetField(key KeySpecimen) Variable
	HasField(key KeySpecimen) bool
	SizeField() int
	IterateFields(func(key Key, value Variable) bool) bool

	SetStaticMethod(key KeySpecimen, action Function)
	GetStaticMethod(key KeySpecimen) Function
	HasStaticMethod(key KeySpecimen) bool
	SizeStaticMethod() int
	IterateStaticMethods(func(key Key, value Function) bool) bool

	SetStaticField(key KeySpecimen, action Variable)
	GetStaticField(key KeySpecimen) Variable
	HasStaticField(key KeySpecimen) bool
	SizeStaticField() int
	IterateStaticFields(func(key Key, value Variable) bool) bool
}

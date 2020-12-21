package concept

type Class interface {
	Variable

	GetName() string

	SetMethodMould(key String, action Function)
	GetMethodMould(key String) Function
	HasMethodMould(key String) bool
	SizeMethodMould() int
	IterateMethodMoulds(func(key String, value Function) bool) bool

	SetFieldMould(key String, defaultFieldMould Variable)
	GetFieldMould(key String) Variable
	HasFieldMould(key String) bool
	SizeFieldMould() int
	IterateFieldMoulds(func(key String, value Variable) bool) bool
}

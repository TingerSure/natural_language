package concept

type Class interface {
	Variable

	GetName() string

	SetProvide(key String, action Function)
	GetProvide(key String) Function
	HasProvide(key String) bool
	IterateProvide(func(key String, value Function) bool) bool

	SetRequire(key String)
	RemoveRequire(key String)
	HasRequire(key String) bool
	IterateRequire(func(key String) bool) bool
}

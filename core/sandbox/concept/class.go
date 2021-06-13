package concept

type Class interface {
	Variable
	SetProvide(key String, action Function)
	GetProvide(key String) Function
	HasProvide(key String) bool
	KeyProvide(key String) String
	IterateProvide(func(key String, value Function) bool) bool

	SetRequire(key String, define Function)
	GetRequire(key String) Function
	HasRequire(key String) bool
	KeyRequire(key String) String
	IterateRequire(func(key String, value Function) bool) bool
}

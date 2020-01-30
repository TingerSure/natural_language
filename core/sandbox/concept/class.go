package concept

type Class interface {
	Variable

	GetName() string

	SetMethod(key string, action Function)
	GetMethod(key string) Function
	HasMethod(key string) bool
	AllMethods() map[string]Function

	SetField(key string, defaultField Variable)
	GetField(key string) Variable
	HasField(key string) bool
	AllFields() map[string]Variable

	SetStaticMethod(key string, action Function)
	GetStaticMethod(key string) Function
	HasStaticMethod(key string) bool
	AllStaticMethods() map[string]Function

	SetStaticField(key string, action Variable)
	GetStaticField(key string) Variable
	HasStaticField(key string) bool
	AllStaticFields() map[string]Variable
}

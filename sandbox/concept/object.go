package concept

type Object interface {
	Variable

	InitField(string, Variable) Exception
	HasField(string) bool
	SetField(string, Variable) Exception
	GetField(string) (Variable, Exception)

	HasMethod(string) bool
	SetMethod(string, Function) Exception
	GetMethod(string) (Function, Exception)
}

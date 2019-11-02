package concept

type Object interface {
	Variable

	GetClasses() []string
	GetClass(string) Class
	GetAliases(string) []string
	IsClassAlias(string, string) bool
	GetMapping(string, string) (map[string]string, Exception)

	InitField(string, Variable) Exception
	HasField(string) bool
	SetField(string, Variable) Exception
	GetField(string) (Variable, Exception)

	HasMethod(string) bool
	SetMethod(string, Function) Exception
	GetMethod(string) (Function, Exception)
}

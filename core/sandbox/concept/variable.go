package concept

type Variable interface {
	GetClasses() []string
	GetClass(string) Class
	GetAliases(string) []string
	IsClassAlias(string, string) bool
	GetMapping(string, string) (map[String]String, Exception)
	CheckMapping(Class, map[String]String) bool
	GetSource() Variable

	InitField(String, Variable) Exception
	HasField(String) bool
	SetField(String, Variable) Exception
	GetField(String) (Variable, Exception)

	HasMethod(String) bool
	SetMethod(String, Function) Exception
	GetMethod(String) (Function, Exception)

	HasRequireMethod(String) bool
	GetRequireMethod(String) (Function, Exception)

	Type() string
	ToString
}

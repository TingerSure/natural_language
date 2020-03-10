package concept

type Object interface {
	Variable

	GetClasses() []string
	GetClass(string) Class
	GetAliases(string) []string
	IsClassAlias(string, string) bool
	GetMapping(string, string) (map[KeySpecimen]KeySpecimen, Exception)
	CheckMapping(Class, map[KeySpecimen]KeySpecimen) bool

	InitField(KeySpecimen, Variable) Exception
	HasField(KeySpecimen) bool
	SetField(KeySpecimen, Variable) Exception
	GetField(KeySpecimen) (Variable, Exception)

	HasMethod(KeySpecimen) bool
	SetMethod(KeySpecimen, Function) Exception
	GetMethod(KeySpecimen) (Function, Exception)
}

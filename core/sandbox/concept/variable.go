package concept

type Variable interface {
	SetField(String, Variable) Exception
	GetField(String) (Variable, Exception)
	Call(String, Param) (Param, Exception)
	GetClass() Class
	GetMapping() *Mapping
	GetSource() Variable
	IsFunction() bool
	Type() string
	ToString
}

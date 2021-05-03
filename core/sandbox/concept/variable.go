package concept

type Variable interface {
	SetField(String, Variable) Exception
	GetField(String) (Variable, Exception)
	HasField(String) bool
	SizeField() int
	Iterate(func(String, Variable) bool) bool
	Call(String, Param) (Param, Exception)
	GetClass() Class
	GetMapping() *Mapping
	GetSource() Variable
	IsFunction() bool
	IsNull() bool
	Type() string
	ToString
}

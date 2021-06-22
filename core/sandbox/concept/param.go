package concept

const (
	ParamTypeList = iota
	ParamTypeKeyValue
)

type Param interface {
	Variable
	SetOriginal(string, Variable)
	GetOriginal(string) Variable
	Set(String, Variable)
	Get(String) Variable
	SizeIndex() int
	AppendIndex(Variable)
	SetIndex(int, Variable)
	GetIndex(int) Variable
	ParamType() int
}

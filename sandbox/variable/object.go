package variable

const (
	VariableObjectType = "object"
)

type Object struct {
}

func (a *Object) ToString(prefix string) string {
	// Todo
	return "TODO Object"
}

func (o *Object) Type() string {
	return VariableObjectType
}

func NewObject() *Object {
	return &Object{}
}

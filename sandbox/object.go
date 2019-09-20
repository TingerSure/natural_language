package sandbox

const (
	VariableObjectType = "object"
)

type Object struct {
}

func (o *Object) Type() string {
	return VariableObjectType
}

func NewObject() *Object {
	return &Object{}
}

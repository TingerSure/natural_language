package concept

type Function interface {
	Variable
	Exec(Param, Object) (Param, Exception)
	FunctionType() string
	ParamNames() []string
	ReturnNames() []string
}

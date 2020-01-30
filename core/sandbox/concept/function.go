package concept

type Function interface {
	Variable
	ExceptionStack
	Exec(Param, Object) (Param, Exception)
	FunctionType() string
	ParamNames() []string
	ReturnNames() []string
}

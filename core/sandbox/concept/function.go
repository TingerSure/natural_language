package concept

type Function interface {
	Variable
	ExceptionStack
	Exec(Param, Object) (Param, Exception)
	FunctionType() string
	ParamNames() []String
	ReturnNames() []String
}

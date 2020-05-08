package concept

type Function interface {
	Variable
	ExceptionStack
	Name() String
	Exec(Param, Object) (Param, Exception)
	FunctionType() string
	ParamNames() []String
	ReturnNames() []String
}

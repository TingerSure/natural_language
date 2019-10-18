package concept

type Function interface {
	Variable
	Exec(Param) (Param, Exception)
	FunctionType() string
}

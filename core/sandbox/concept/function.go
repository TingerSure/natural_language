package concept

type Function interface {
	Variable
	ExceptionStack
	Exec(Param, Variable) (Param, Exception)
	Anticipate(Param, Variable) Param
	FunctionType() string
	AddParamName(...String)
	AddReturnName(...String)
	ParamNames() []String
	ReturnNames() []String
	GetLanguageOnCallSeed(string) func(Function, *Mapping) string
	SetLanguageOnCallSeed(string, func(Function, *Mapping) string)
	ParamFormat(*Mapping) *Mapping
	ReturnFormat(String) String
}

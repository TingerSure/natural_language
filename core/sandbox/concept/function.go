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
	GetLanguageOnCallSeed(string) func(Function, Pool, string, Param) string
	SetLanguageOnCallSeed(string, func(Function, Pool, string, Param) string)
	ToCallLanguage(string, Pool, string, Param) string
	ParamFormat(Param) Param
	ReturnFormat(String) String
}

package concept

type Function interface {
	Variable
	Exec(Param, Variable) (Param, Exception)
	FunctionType() string
	AddParamName(...String)
	AddReturnName(...String)
	ParamNames() []String
	ReturnNames() []String
	GetLanguageOnCallSeed(string) func(Function, Pool, string, Param) (string, Exception)
	SetLanguageOnCallSeed(string, func(Function, Pool, string, Param) (string, Exception))
	ToCallLanguage(string, Pool, string, Param) (string, Exception)
	ParamFormat(Param) Param
	ReturnFormat(String) String
}

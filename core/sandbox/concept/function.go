package concept

type Function interface {
	Variable
	ExceptionStack
	Name() String
	Exec(Param, Object) (Param, Exception)
	Anticipate(Param, Object) Param
	FunctionType() string
	ParamNames() []String
	ReturnNames() []String
	GetLanguageOnCallSeed(string) func(Function, *Mapping) string
	SetLanguageOnCallSeed(string, func(Function, *Mapping) string)
	ParamFormat(*Mapping) *Mapping
	ReturnFormat(String) String
}

package concept

type Exception interface {
	error
	Interrupt
	Variable
	Name() String
	Message() String
	Copy() Exception
	IterateStacks(func(ExceptionStack) bool) bool
	AddStack(ExceptionStack) Exception
}

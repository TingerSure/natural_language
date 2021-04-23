package concept

type Exception interface {
	Interrupt
	Variable
	Name() String
	Message() String
	Copy() Exception
	IterateStacks(func(ExceptionStack) bool) bool
	AddStack(ExceptionStack) Exception
}

package concept

type Exception interface {
	Interrupt
	ToString
	Name() string
	Message() string
	Copy() Exception
	IterateStacks(func(ExceptionStack) bool) bool
	AddStack(ExceptionStack) Exception
}

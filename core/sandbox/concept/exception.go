package concept

type Exception interface {
	error
	Interrupt
	Variable
	Name() String
	Message() String
	Copy() Exception
	AddExceptionLine(Line) Exception
}

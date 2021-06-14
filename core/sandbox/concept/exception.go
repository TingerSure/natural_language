package concept

type Exception interface {
	error
	Interrupt
	Variable
	Name() String
	Message() String
	Copy() Exception
	IterateLines(func(Line) bool) bool
	AddLine(Line) Exception
}

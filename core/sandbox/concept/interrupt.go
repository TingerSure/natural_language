package concept

type Interrupt interface {
	InterruptType() string
	IterateLines(func(Line) bool) bool
	AddLine(Line) Interrupt
}

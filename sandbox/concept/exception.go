package concept

type Exception interface {
	Interrupt
	ToString
	Name() string
	Message() string
}

package concept

type Exception interface {
	Interrupt
	Name() string
	Message() string
}

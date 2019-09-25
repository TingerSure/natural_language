package sandbox

const (
	ExceptionInterruptType = "exception"
)

type Exception struct {
	name    string
	message string
}

func (e *Exception) InterruptType() string {
	return ExceptionInterruptType
}

func (e *Exception) Name() string {
	return e.name
}

func (e *Exception) Message() string {
	return e.message
}
func NewException(name string, message string) *Exception {
	return &Exception{
		name:    name,
		message: message,
	}
}

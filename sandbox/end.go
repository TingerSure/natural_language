package sandbox

const (
	EndInterruptType = "end"
)

type End struct {
}

func (e *End) InterruptType() string {
	return EndInterruptType
}

func NewEnd() *End {
	return &End{}
}

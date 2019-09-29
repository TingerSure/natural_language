package interrupt

const (
	ContinueInterruptType = "continue"
)

type Continue struct {
	tag string
}

func (e *Continue) InterruptType() string {
	return ContinueInterruptType
}

func (e *Continue) Tag() string {
	return e.tag
}

func NewContinueWithTag(tag string) *Continue {
	return &Continue{
		tag: tag,
	}
}

func NewContinue() *Continue {
	return &Continue{
		tag: "",
	}
}

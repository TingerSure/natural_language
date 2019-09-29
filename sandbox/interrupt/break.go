package interrupt

const (
	BreakInterruptType = "break"
)

type Break struct {
	tag string
}

func (e *Break) InterruptType() string {
	return BreakInterruptType
}

func (e *Break) Tag() string {
	return e.tag
}

func NewBreak() *Break {
	return &Break{
		tag: "",
	}
}

func NewBreakWithTag(tag string) *Break {
	return &Break{
		tag: tag,
	}
}

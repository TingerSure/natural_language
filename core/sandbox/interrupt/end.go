package interrupt

const (
	EndInterruptType = "end"
)

type EndSeed interface {
	Type() string
}

type End struct {
	seed EndSeed
}

func (e *End) InterruptType() string {
	return e.seed.Type()
}

type EndCreatorParam struct {
}

type EndCreator struct {
	param *EndCreatorParam
}

func (s *EndCreator) New() *End {
	return &End{
		seed: s,
	}
}

func (s *EndCreator) Type() string {
	return EndInterruptType
}

func NewEndCreator(param *EndCreatorParam) *EndCreator {
	return &EndCreator{
		param: param,
	}
}

package concept

type Key interface {
	Variable
	Get(KeySpecimen)
	Set(KeySpecimen)
	Is(KeySpecimen) bool
	Equal(Key) bool
	Iterate(func(KeySpecimen) bool) bool
}

package concept

type Matcher interface {
	ToString
	Match(Variable) bool
}

package concept

type Variable interface {
	Type() string
	ToString(string) string
}

package sandbox

type Variable interface {
	Type() string
	ToString(prefix string) string
}

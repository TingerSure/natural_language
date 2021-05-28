package concept

type ToString interface {
	ToString(prefix string) string
	ToLanguage(language string, space Closure) string
}

package concept

type ToString interface {
	ToString(prefix string) string
	ToLanguage(language string) string
}

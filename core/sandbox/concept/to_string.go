package concept

type ToString interface {
	ToString(prefix string) string
	ToLanguage(language string, space Pool) string
}

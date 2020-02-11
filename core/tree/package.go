package tree

type Package interface {
	GetSources() []Source
	AddSource(Source)
}

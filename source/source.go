package source

type Source interface {
	GetName() string
	GetWords(firstCharacter string) []string
}

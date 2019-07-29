package tree

type Phrase interface {
	Copy() Phrase
	Size() int
	GetContent() *Vocabulary
	GetChild(index int) Phrase
	SetChild(index int, child Phrase)
	ToString() string
	ToStringOffset(index int) string
}

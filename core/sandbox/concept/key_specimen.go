package concept

type KeySpecimen interface {
	Variable
	GetLanguage() string
	GetName() string
	SetLanguage(string)
	SetName(string)
	Equal(KeySpecimen) bool
}

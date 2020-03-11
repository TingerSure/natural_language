package concept

type String interface {
	Variable
	Value() string
	GetSystem() string
	SetSystem(string)
	GetLanguage(string) string
	SetLanguage(string, string)
	IsLanguage(string, string) bool
	EqualLanguage(String) bool
	Equal(String) bool
	Clone() String
	IterateLanguages(func(string, string) bool) bool
}

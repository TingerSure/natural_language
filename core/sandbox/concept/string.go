package concept

type String interface {
	Variable
	Value() string
	GetLanguage(string) string
	SetLanguage(string, string)
	HasLanguage(string) bool
	IsLanguage(string, string) bool
	Equal(String) bool
	Clone() String
	CloneTo(String)
	IterateLanguages(func(string, string) bool) bool
}

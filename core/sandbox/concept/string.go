package concept

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type String interface {
	nl_interface.Key
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

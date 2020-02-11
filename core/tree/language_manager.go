package tree

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type LanguageManager struct {
	languages map[string]Language
}

func (l *LanguageManager) AddLanguage(name string, lib Language) {
	l.languages[name] = lib
}

func (l *LanguageManager) GetLanguage(name string) Language {
	return l.languages[name]
}

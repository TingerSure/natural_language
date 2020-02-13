package tree

import (
// "github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type LanguageManager struct {
	languages map[string]Language
}

func (l *LanguageManager) AddLanguage(name string, language Language) {
	l.languages[name] = language
}

func (l *LanguageManager) GetLanguage(name string) Language {
	return l.languages[name]
}

func (l *LanguageManager) AllLanguages() []string {
	names := []string{}
	for name, _ := range l.languages {
		names = append(names, name)
	}
	return names
}

func (l *LanguageManager) LanguagesIterate(onLanguage func(string, Language) bool) bool {
	for name, language := range l.languages {
		if onLanguage(name, language) {
			return true
		}
	}
	return false
}

func NewLanguageManager() *LanguageManager {
	return &LanguageManager{}
}

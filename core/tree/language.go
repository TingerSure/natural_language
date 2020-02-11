package tree

import ()

type Language struct {
	packages []Package
}

func (l *Language) SetPackage(name string, instance Package) {
	l.packages[name] = instance
}

func (l *Language) GetPackage(name string) Package {
	return l.packages[name]
}

func NewLanguage() *Language {
	return &Language{}
}

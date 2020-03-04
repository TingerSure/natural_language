package tree

import ()

type Language struct {
	packages map[string]Package
	name     string
}

func (l *Language) SetName(name string) {
	l.name = name
}

func (l *Language) GetName() string {
	return l.name
}

func (l *Language) SetPackage(name string, instance Package) {
	l.packages[name] = instance
}

func (l *Language) GetPackage(name string) Package {
	return l.packages[name]
}

func (l *Language) PackagesIterate(onPackage func(string, Package) bool) bool {
	for name, instance := range l.packages {
		if onPackage(name, instance) {
			return true
		}
	}
	return false
}

func (l *Language) AllPackages() []string {
	names := []string{}
	for name, _ := range l.packages {
		names = append(names, name)
	}
	return names
}

func NewLanguage() *Language {
	return &Language{
		packages: map[string]Package{},
	}
}

package tree

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type LibraryManager struct {
	libraries map[string]Library
}

func (l *LibraryManager) PageIterate(on func(page Page) bool) bool {
	for _, lib := range l.libraries {
		if lib.PageIterate(on) {
			return true
		}
	}
	return false
}

func (l *LibraryManager) GetLibraryPage(libraryName string, pageName string) Page {
	if libraryName == "" {
		return l.GetPage(pageName)
	}
	return l.GetLibrary(libraryName).GetPage(pageName)
}

func (l *LibraryManager) GetPage(pageName string) Page {
	for _, library := range l.libraries {
		page := library.GetPage(pageName)
		if !nl_interface.IsNil(page) {
			return page
		}
	}
	return nil
}

func (l *LibraryManager) AddLibrary(name string, lib Library) {
	l.libraries[name] = lib
}

func (l *LibraryManager) AddSystemLibrary(lib Library) {
	l.AddLibrary("system", lib)
}

func (l *LibraryManager) GetLibrary(name string) Library {
	return l.libraries[name]
}

func NewLibraryManager() *LibraryManager {
	return &LibraryManager{
		libraries: map[string]Library{},
	}
}

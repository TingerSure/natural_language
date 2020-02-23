package tree

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type LibraryManager struct {
	librarys map[string]Library
}

func (l *LibraryManager) GetLibraryPage(libraryName string, pageName string) Page {
	if libraryName == "" {
		return l.GetPage(pageName)
	}
	return l.GetLibrary(libraryName).GetPage(pageName)
}

func (l *LibraryManager) GetPage(pageName string) Page {
	for _, library := range l.librarys {
		page := library.GetPage(pageName)
		if !nl_interface.IsNil(page) {
			return page
		}
	}
	return nil
}

func (l *LibraryManager) AddLibrary(name string, lib Library) {
	l.librarys[name] = lib
}

func (l *LibraryManager) AddSystemLibrary(lib Library) {
	l.AddLibrary("system", lib)
}

func (l *LibraryManager) GetLibrary(name string) Library {
	return l.librarys[name]
}

func NewLibraryManager() *LibraryManager {
	return &LibraryManager{
		librarys: map[string]Library{},
	}
}
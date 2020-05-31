package runtime

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/creator"
	"github.com/TingerSure/natural_language/core/tree"
)

type LibraryManager struct {
	libraries map[string]tree.Library
	Sandbox   *creator.SandboxCreator
}

func (l *LibraryManager) PageIterate(on func(page tree.Page) bool) bool {
	for _, lib := range l.libraries {
		if lib.PageIterate(on) {
			return true
		}
	}
	return false
}

func (l *LibraryManager) GetLibraryPage(libraryName string, pageName string) tree.Page {
	if libraryName == "" {
		return l.GetPage(pageName)
	}
	return l.GetLibrary(libraryName).GetPage(pageName)
}

func (l *LibraryManager) GetPage(pageName string) tree.Page {
	for _, library := range l.libraries {
		page := library.GetPage(pageName)
		if !nl_interface.IsNil(page) {
			return page
		}
	}
	return nil
}

func (l *LibraryManager) AddLibrary(name string, lib tree.Library) {
	l.libraries[name] = lib
}

func (l *LibraryManager) AddSystemLibrary(lib tree.Library) {
	l.AddLibrary("system", lib)
}

func (l *LibraryManager) GetLibrary(name string) tree.Library {
	return l.libraries[name]
}

func NewLibraryManager() *LibraryManager {
	return &LibraryManager{
		libraries: map[string]tree.Library{},
		Sandbox:   creator.NewSandboxCreator(),
	}
}

package semantic

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/tree"
)

type FilePage struct {
	tree.Page
	name    string
	imports map[string]tree.Page
}

func NewFilePage(libs *tree.LibraryManager) *FilePage {
	return &FilePage{
		Page:    tree.NewPageAdaptor(libs.Sandbox),
		imports: map[string]tree.Page{},
	}
}

func (f *FilePage) GetName() string {
	return f.name
}

func (f *FilePage) SetName(name string) {
	f.name = name
}

func (f *FilePage) AddImport(key string, page tree.Page) error {
	if !nl_interface.IsNil(f.imports[key]) {
		return errors.New(fmt.Sprintf("semantic error : import repeated : %v", key))
	}
	f.imports[key] = page
	return nil
}

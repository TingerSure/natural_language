package tree

type LibraryAdaptor struct {
	pages map[string]*Page
}

func (p *LibraryAdaptor) GetPage(key string) *Page {
	return p.pages[key]
}

func (p *LibraryAdaptor) SetPage(key string, value *Page) Library {
	p.pages[key] = value
	return p
}

func NewLibraryAdaptor() *LibraryAdaptor {
	return &LibraryAdaptor{
		pages: map[string]*Page{},
	}
}

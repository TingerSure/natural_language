package tree

type LibraryAdaptor struct {
	pages map[string]Page
}

func (p *LibraryAdaptor) GetPage(key string) Page {
	return p.pages[key]
}

func (p *LibraryAdaptor) SetPage(key string, value Page) Library {
	p.pages[key] = value
	return p
}

func (p *LibraryAdaptor) PageIterate(on func(Page) bool) bool {
	for _, page := range p.pages {
		if on(page) {
			return true
		}
	}
	return false
}

func NewLibraryAdaptor() *LibraryAdaptor {
	return &LibraryAdaptor{
		pages: map[string]Page{},
	}
}

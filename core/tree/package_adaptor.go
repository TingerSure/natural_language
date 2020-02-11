package tree

import ()

type PackageAdaptor struct {
	sources []Source
}

func (p *PackageAdaptor) GetSources() []Source {
	return p.sources
}

func (p *PackageAdaptor) AddSource(source Source) {
	p.sources = append(p.sources, source)
}

func NewPackageAdaptor() *PackageAdaptor {
	return &PackageAdaptor{}
}

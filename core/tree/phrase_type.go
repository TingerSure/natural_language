package tree

type PhraseTypeParam struct {
	Parents []*PhraseTypeParent
	Name    string
	From    string
}

type PhraseType struct {
	param *PhraseTypeParam
}

func (wanted *PhraseType) Equal(given *PhraseType) bool {
	return wanted.param.Name == given.param.Name
}

func (p *PhraseType) Name() string {
	return p.param.Name
}

func (p *PhraseType) GetFrom() string {
	return p.param.From
}

func (p *PhraseType) Parents() []*PhraseTypeParent {
	return p.param.Parents
}

func NewPhraseType(param *PhraseTypeParam) *PhraseType {
	return &PhraseType{
		param: param,
	}
}

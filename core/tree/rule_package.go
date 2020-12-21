package tree

type PackageRuleParam struct {
	Pack func(Phrase) Phrase
	From string
}

type PackageRule struct {
	param *PackageRuleParam
}

func (r *PackageRule) GetFrom() string {
	return r.param.From
}

func (r *PackageRule) Pack(treasure Phrase) Phrase {
	return r.param.Pack(treasure)
}

func NewPackageRule(param *PackageRuleParam) *PackageRule {
	if param.Pack == nil {
		panic("No pack function in this package rule!")
	}
	return &PackageRule{
		param: param,
	}
}

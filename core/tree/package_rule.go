package tree

type PackageRuleParam struct {
	Create func(Phrase) Phrase
	From   string
}

type PackageRule struct {
	param *PackageRuleParam
}

func (r *PackageRule) GetFrom() string {
	return r.param.From
}

func (r *PackageRule) Create(treasure Phrase) Phrase {
	return r.param.Create(treasure)
}

func NewPackageRule(param *PackageRuleParam) *PackageRule {
	if param.Create == nil {
		panic("no create function in this package rule!")
	}
	return &PackageRule{
		param: param,
	}
}

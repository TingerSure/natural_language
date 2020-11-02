package phrase_type

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
)

const (
	NumberName           string = "types.number"
	EntityFromNumberName string = "package.entity.number"
)

type Number struct {
	*adaptor.SourceAdaptor
	NewAutoNumberContent concept.String
	NewAutoNumberObject  concept.String
	NewAutoNumber        concept.Function
}

func (p *Number) GetName() string {
	return NumberName
}

func (p *Number) GetPhraseTypes() []*tree.PhraseType {
	return []*tree.PhraseType{
		tree.NewPhraseType(&tree.PhraseTypeParam{
			Name: NumberName,
			From: NumberName,
			Parents: []*tree.PhraseTypeParent{
				&tree.PhraseTypeParent{
					Types: EntityName,
					Rule: tree.NewPackageRule(&tree.PackageRuleParam{
						Create: func(value tree.Phrase) tree.Phrase {
							return tree.NewPhrasePackage(&tree.PhrasePackageParam{
								Index: func(phrase tree.Phrase) concept.Index {
									return p.Libs.Sandbox.Expression.ParamGet.New(
										p.Libs.Sandbox.Expression.Call.New(
											p.Libs.Sandbox.Index.ConstIndex.New(p.NewAutoNumber),
											p.Libs.Sandbox.Expression.NewParam.New().Init(map[concept.String]concept.Index{
												p.NewAutoNumberContent: phrase.Index(),
											}),
										),
										p.NewAutoNumberObject,
									)
								},
								Types: EntityName,
								From:  EntityFromNumberName,
							}).SetValue(value)
						},
						From: NumberName,
					}),
				},
			},
		}),
	}
}

func NewNumber(param *adaptor.SourceAdaptorParam) *Number {
	instance := &Number{
		SourceAdaptor: adaptor.NewSourceAdaptor(param),
	}
	libObject := instance.Libs.GetLibraryPage("system", "auto_number")
	instance.NewAutoNumberContent = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("NewAutoNumberContent"))
	instance.NewAutoNumberObject = libObject.GetConst(instance.Libs.Sandbox.Variable.String.New("NewAutoNumberObject"))
	instance.NewAutoNumber = libObject.GetFunction(instance.Libs.Sandbox.Variable.String.New("NewAutoNumber"))
	return instance
}

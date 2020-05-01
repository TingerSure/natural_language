package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type TranslateRuleParam struct {
	Source concept.String
	Target concept.String
}

type TranslateRule struct {
	param *TranslateRuleParam
}

func (t *TranslateRule) Register(language string) {
	t.param.Source.SetLanguage(language, t.param.Target.GetLanguage(language))
}

func NewTranslateRule(param *TranslateRuleParam) *TranslateRule {
	return &TranslateRule{
		param: param,
	}
}

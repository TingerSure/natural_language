package bind

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newKeyBind(libs *tree.LibraryManager) *variable.SystemFunction {
	languageParam := libs.Sandbox.Variable.String.New("language")
	keysParam := libs.Sandbox.Variable.String.New("keys")
	valueParam := libs.Sandbox.Variable.String.New("value")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, object concept.Variable) (concept.Param, concept.Exception) {
			languagePre := input.Get(languageParam)
			language, yes := variable.VariableFamilyInstance.IsStringHome(languagePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param language is not a string: %v", languagePre.ToString("")))
			}
			keysPre := input.Get(keysParam)
			keys, yes := variable.VariableFamilyInstance.IsFunction(keysPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param keys is not a function: %v", keysPre.ToString("")))
			}
			valuePre := input.Get(valueParam)
			value, yes := variable.VariableFamilyInstance.IsStringHome(valuePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param value is not a string: %v", valuePre.ToString("")))
			}
			for cursor, keyPre := range keys.Body().Steps() {
				bubble, yes := index.IndexFamilyInstance.IsBubbleIndex(keyPre)
				if yes {
					exception := bubbleKeyBind(libs, keys.Parent(), bubble, language, value)
					if !nl_interface.IsNil(exception) {
						return nil, exception
					}
					continue
				}
				component, yes := expression.ExpressionFamilyInstance.IsComponent(keyPre)
				if yes {
					exception := componentKeyBind(libs, keys.Parent(), component, language, value)
					if !nl_interface.IsNil(exception) {
						return nil, exception
					}
					continue
				}
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Pipe in keys is unsupported. line : %v, pipe: %v", cursor, keys.ToString("")))
			}
			return libs.Sandbox.Variable.Param.New(), nil
		},
		[]concept.String{
			languageParam,
			keysParam,
			valueParam,
		},
		[]concept.String{},
	)
}

func componentKeyBind(libs *tree.LibraryManager, pool concept.Pool, component *expression.Component, language concept.String, value concept.String) concept.Exception {
	object, suspend := component.Object().Get(pool)
	if !nl_interface.IsNil(suspend) {
		return suspend.(concept.Exception)
	}
	if object.IsNull() {
		return libs.Sandbox.Variable.Exception.NewOriginal("none pionter", fmt.Sprintf("Null variable: %v.", component.Object().ToString("")))
	}
	object.KeyField(component.Field()).SetLanguage(language.Value(), value.Value())
	return nil
}

func bubbleKeyBind(libs *tree.LibraryManager, pool concept.Pool, bubble *index.BubbleIndex, language concept.String, value concept.String) concept.Exception {
	if !pool.HasBubble(bubble.Key()) {
		return libs.Sandbox.Variable.Exception.NewOriginal("none pionter", fmt.Sprintf("Undefined variable: %v.", bubble.ToString("")))
	}
	pool.KeyBubble(bubble.Key()).SetLanguage(language.Value(), value.Value())
	return nil
}

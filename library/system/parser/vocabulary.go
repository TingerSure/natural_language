package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newAddVocabularyRule(libs *tree.LibraryManager) concept.Function {
	nameParam := libs.Sandbox.Variable.String.New("name")
	typesParam := libs.Sandbox.Variable.String.New("types")
	wordsParam := libs.Sandbox.Variable.String.New("words")
	matchParam := libs.Sandbox.Variable.String.New("match")
	createParam := libs.Sandbox.Variable.String.New("create")

	instance := libs.Sandbox.Variable.SystemFunction.New(
		func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			namePre := input.Get(nameParam)
			name, yes := variable.VariableFamilyInstance.IsStringHome(namePre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param name is not a string: %v", namePre.ToString("")))
			}
			typesPre := input.Get(typesParam)
			types, yes := variable.VariableFamilyInstance.IsStringHome(typesPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param types is not a string: %v", typesPre.ToString("")))
			}
			matchPre := input.Get(matchParam)
			wordsPre := input.Get(wordsParam)
			if wordsPre.IsNull() && matchPre.IsNull() {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("param error", "Both words and match are empty.")
			}
			var match concept.String = libs.Sandbox.Variable.String.New("")
			if !matchPre.IsNull() {
				match, yes = variable.VariableFamilyInstance.IsStringHome(matchPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param match is not a string: %v", matchPre.ToString("")))
				}
			}
			vocabularyWords := []string{}
			if !wordsPre.IsNull() {
				words, yes := variable.VariableFamilyInstance.IsArray(wordsPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param words is not a string: %v", wordsPre.ToString("")))
				}
				for index := 0; index < words.Length(); index++ {
					wordPre, exception := words.Get(index)
					if !nl_interface.IsNil(exception) {
						return nil, exception
					}
					word, yes := variable.VariableFamilyInstance.IsStringHome(wordPre)
					if !yes {
						return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param words[%v] is not a string: %v", index, wordPre.ToString("")))
					}
					vocabularyWords = append(vocabularyWords, word.Value())
				}
			}
			createPre := input.Get(createParam)
			create, yes := variable.VariableFamilyInstance.IsFunctionHome(createPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param create is not a function: %v", createPre.ToString("")))
			}

			libs.Vocabularys.AddRule(tree.NewVocabularyRule(&tree.VocabularyRuleParam{
				From:  name.Value(),
				Words: vocabularyWords,
				Match: match.Value(),
				Create: func(treasure string) tree.Phrase {
					return tree.NewPhraseVocabulary(&tree.PhraseVocabularyParam{
						Content: treasure,
						Types:   types.Value(),
						From:    name.Value(),
						Index: func() (concept.Function, concept.Exception) {
							line := tree.NewLine(fmt.Sprintf("[vocabulary_parse]:%v ( %v )", name.Value(), treasure), "")
							param := libs.Sandbox.Variable.Param.New()
							param.SetOriginal("content", libs.Sandbox.Variable.String.New(treasure))
							output, exception := create.Exec(param, nil)
							if !nl_interface.IsNil(exception) {
								return nil, exception.AddExceptionLine(line)
							}
							pipe, yes := variable.VariableFamilyInstance.IsFunctionHome(output.GetOriginal("pipe"))
							if !yes {
								return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", "vocabulary parse pipe must be defined as 'function()value'.").AddExceptionLine(line)
							}
							return pipe, nil
						},
					})
				},
			}))
			return libs.Sandbox.Variable.Param.New(), nil
		},
		[]concept.String{
			nameParam,
			typesParam,
			wordsParam,
			matchParam,
			createParam,
		},
		[]concept.String{},
	)
	return instance
}

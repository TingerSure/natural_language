package parser

import (
	"github.com/TingerSure/natural_language/core/sandbox/pool"
	"github.com/TingerSure/natural_language/core/tree"
)

type Parser struct {
	section   *Section
	barricade *Barricade
	reach     *Reach
	lexer     *Lexer
	grammar   *Grammar
	types     *Types
	diversion *Diversion
	rootSpace *pool.Pool
}

func NewParser(rootSpace *pool.Pool) *Parser {
	types := NewTypes()
	section := NewSection()
	barricade := NewBarricade()
	diversion := NewDiversion(rootSpace)
	reach := NewReach(types)
	return &Parser{
		rootSpace: rootSpace,
		types:     types,
		section:   section,
		barricade: barricade,
		reach:     reach,
		diversion: diversion,
		lexer:     NewLexer(section),
		grammar:   NewGrammar(types, reach, barricade, diversion),
	}
}

func (p *Parser) Instance(sentence string) (*Road, error) {
	road := NewRoad(sentence, p.types)

	err := p.lexer.ParseVocabulary(road)
	if err != nil {
		return nil, err
	}

	err = p.grammar.ParseStruct(road)
	if err != nil {
		return nil, err
	}

	return road, nil
}

func (p *Parser) GetSection() *Section {
	return p.section
}

func (p *Parser) GetLexer() *Lexer {
	return p.lexer
}

func (p *Parser) AddSource(source tree.Source) {
	p.lexer.AddSource(source)
	p.section.AddRule(source.GetVocabularyRules())
	p.barricade.AddRule(source.GetPriorityRules())
	p.reach.AddRule(source.GetStructRules())
	p.diversion.AddRule(source.GetDutyRules())
	p.types.AddTypes(source.GetPhraseTypes())
}

func (p *Parser) RemoveSource(name string) {
	p.lexer.RemoveSource(name)
	p.section.RemoveRule(func(rule *tree.VocabularyRule) bool {
		return rule.GetFrom() == name
	})
	p.barricade.RemoveRule(func(rule *tree.PriorityRule) bool {
		return rule.GetFrom() == name
	})
	p.reach.RemoveRule(func(rule *tree.StructRule) bool {
		return rule.GetFrom() == name
	})
	p.diversion.RemoveRule(func(rule *tree.DutyRule) bool {
		return rule.GetFrom() == name
	})
	p.types.RemoveTypes(func(value *tree.PhraseType) bool {
		return value.GetFrom() == name
	})
}

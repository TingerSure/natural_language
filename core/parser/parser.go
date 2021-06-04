package parser

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Parser struct {
	barricade *Barricade
	reach     *Reach
	lexer     *Lexer
	grammar   *Grammar
	types     *Types
	diversion *Diversion
	rootSpace concept.Pool
}

func NewParser(rootSpace concept.Pool) *Parser {
	types := NewTypes()
	barricade := NewBarricade()
	diversion := NewDiversion(rootSpace)
	reach := NewReach(types)
	return &Parser{
		rootSpace: rootSpace,
		types:     types,
		barricade: barricade,
		reach:     reach,
		diversion: diversion,
		lexer:     NewLexer(),
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

func (p *Parser) GetBarricade() *Barricade {
	return p.barricade
}

func (p *Parser) GetDiversion() *Diversion {
	return p.diversion
}

func (p *Parser) GetReach() *Reach {
	return p.reach
}

func (p *Parser) GetTypes() *Types {
	return p.types
}

func (p *Parser) GetLexer() *Lexer {
	return p.lexer
}

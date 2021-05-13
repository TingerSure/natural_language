package grammar

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
)

type Automata struct {
	table *Table
}

func NewAutomata(table *Table) *Automata {
	return &Automata{
		table: table,
	}
}

func (a *Automata) Run(tokens *lexer.TokenList) (Phrase, error) {
	tokens.Reset()
	status := a.table.GetStart()
	phrase := NewPhraseToken(tokens.Next())
	statusList := []int{status}
	phraseList := []Phrase{}
	for {
		action := a.table.GetAction(status, phrase.GetToken().Type())
		if action == nil {
			return nil, errors.New(fmt.Sprintf("syntax error : unexpected : %v, status : %v, symbol : %v", phrase.GetToken().Value(), status, phrase.GetToken().Name()))
		}
		if action.Type() == ActionMoveType {
			status = action.Status()
			statusList = append(statusList, status)
			phraseList = append(phraseList, phrase)
			next := tokens.Next()
			if next == nil {
				if phrase.GetToken() != tokens.End() {
					return nil, errors.New("automata error : illegal token list")
				}
				return nil, errors.New("automata error : illegal status table")
			}
			phrase = NewPhraseToken(next)
			continue
		}
		if action.Type() == ActionPolymerizeType {
			rule := action.Rule()
			size := rule.Size()
			result := NewPhraseStruct(rule)
			result.AddChild(phraseList[len(phraseList)-size:]...)
			phraseList = phraseList[:len(phraseList)-size]
			statusList = statusList[:len(statusList)-size]
			status = statusList[len(statusList)-1]
			gotos := a.table.GetGoto(status, result.Type())
			if gotos == nil {
				return nil, errors.New("syntax error : unexpected block")
			}
			status = gotos.Status()
			phraseList = append(phraseList, result)
			statusList = append(statusList, status)
			continue
		}
		if action.Type() == ActionAcceptType {
			if len(phraseList) != 1 || phrase.GetToken() != tokens.End() {
				return nil, errors.New("automata error : illegal status table")
			}
			if !tokens.IsEnd() {
				return nil, errors.New("automata error : illegal token list")
			}
			break
		}
		return nil, errors.New(fmt.Sprintf("automata error : illegal action type : %v", action.Type()))
	}

	return phraseList[0], nil
}

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
	statusList := []int{}
	phraseList := []Phrase{}
	status := a.table.GetStart()
	phrase := NewPhraseToken(tokens.Next())
	for {
		action := a.table.GetAction(status, phrase.GetToken().Type())
		if action == nil {
			return nil, errors.New(fmt.Sprint("syntax error : unexpected : %v", phrase.GetToken().Value()))
		}
		if action.Type() == ActionMoveType {
			statusList = append(statusList, status)
			phraseList = append(phraseList, phrase)
			status = action.Status()
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
			phraseList = append(phraseList, result)
			statusList = statusList[:len(statusList)-size+1]
			lastStatus := statusList[len(statusList)-1]
			gotos := a.table.GetGoto(lastStatus, result.Type())
			if gotos == nil {
				return nil, errors.New("syntax error : unexpected block")
			}
			status = gotos.Status()
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
		return nil, errors.New(fmt.Sprint("automata error : illegal action type : %v", action.Type()))
	}

	return phraseList[0], nil
}

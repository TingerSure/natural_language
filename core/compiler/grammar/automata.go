package grammar

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler/lexer"
	"strings"
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
	line := tokens.StartLine()
	next := tokens.Next()
	status := a.table.GetStart()
	phrase := NewPhraseToken(next)

	statusList := []int{status}
	phraseList := []Phrase{}
	for {
		action := a.table.GetAction(status, phrase.GetToken().Type())
		if action == nil {
			expectations := a.table.GetExpect(status)
			names := []string{}
			for _, expectation := range expectations {
				names = append(names, expectation.Name())
			}
			return nil, fmt.Errorf("syntax error : unexpected : '%v' (%v), expecting : (%v).\n%v", phrase.GetToken().Value(), phrase.GetToken().Name(), strings.Join(names, ", "), line.End().ToString())
		}
		if action.Type() == ActionMoveType {
			status = action.Status()
			statusList = append(statusList, status)
			phraseList = append(phraseList, phrase)
			line = next.Line()
			next = tokens.Next()
			if next == nil {
				if phrase.GetToken() != tokens.Eof() {
					return nil, fmt.Errorf("automata error : illegal token list.")
				}
				return nil, fmt.Errorf("automata error : illegal status table.")
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
			if result.Size() == 0 {
				result.SetStartLine(line.End())
			}
			gotos := a.table.GetGoto(status, result.Type())
			if gotos == nil {
				return nil, fmt.Errorf("automata error : illegal status table.")
			}
			status = gotos.Status()
			phraseList = append(phraseList, result)
			statusList = append(statusList, status)
			continue
		}
		if action.Type() == ActionAcceptType {
			if len(phraseList) != 1 || phrase.GetToken() != tokens.Eof() {
				return nil, fmt.Errorf("automata error : illegal status table.")
			}
			if !tokens.IsEof() {
				return nil, fmt.Errorf("automata error : illegal token list.")
			}
			break
		}
		return nil, fmt.Errorf("automata error : illegal action type : %v", action.Type())
	}

	return phraseList[0], nil
}

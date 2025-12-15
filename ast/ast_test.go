package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	progerm := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if progerm.String() != "let myVar = anotherVar;" {
		t.Errorf("progerm.String() wrong. got=%q", progerm.String())
	}
}
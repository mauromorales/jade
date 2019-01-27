package ast

import (
	"testing"

	"github.com/mauromorales/jade/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.DEFINE, Literal: "definir"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "mi_variable"},
					Value: "mi_variable",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "otra_variable"},
					Value: "otra_variable",
				},
			},
		},
	}

	if program.String() != "definir mi_variable = otra_variable" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

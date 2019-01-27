package lexer

import (
	"testing"

	"github.com/mauromorales/jade/token"
)

func TestNextToken(t *testing.T) {
	input := `cinco = 5
diez = 10

suma = funcion(x, y) { x + y }

resultado = suma(cinco, diez)

!-/*5
5 < 10 > 5

si 5 < 10 entonces
	devolver verdadero
sino
	devolver falso
fin

10 == 10
10 != 9
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "cinco"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.IDENT, "diez"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.IDENT, "suma"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "funcion"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.IDENT, "resultado"},
		{token.ASSIGN, "="},
		{token.IDENT, "suma"},
		{token.LPAREN, "("},
		{token.IDENT, "cinco"},
		{token.COMMA, ","},
		{token.IDENT, "diez"},
		{token.RPAREN, ")"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.STAR, "*"},
		{token.INT, "5"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.IF, "si"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.THEN, "entonces"},
		{token.RETURN, "devolver"},
		{token.TRUE, "verdadero"},
		{token.ELSE, "sino"},
		{token.RETURN, "devolver"},
		{token.FALSE, "falso"},
		{token.END, "fin"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType || tok.Literal != tt.expectedLiteral {
			t.Logf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)

			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

package lexer

import (
	"testing"

	"github.com/mauromorales/jade/token"
)

func TestNextToken(t *testing.T) {
	input := `definir cinco = 5
definir diez = 10

definir suma = funcion(x, y) { x + y }

definir resultado = suma(cinco, diez)

!-/*5
5 < 10 > 5

si 5 < 10 entonces
	devolver verdadero
sino
	devolver falso
fin

10 == 10
10 != 9
"foobar"
"foo bar"
[1, 2]
{"foo": "bar"}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEFINE, "definir"},
		{token.IDENT, "cinco"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.NEW_LINE, "\n"},
		{token.DEFINE, "definir"},
		{token.IDENT, "diez"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.DEFINE, "definir"},
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
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.DEFINE, "definir"},
		{token.IDENT, "resultado"},
		{token.ASSIGN, "="},
		{token.IDENT, "suma"},
		{token.LPAREN, "("},
		{token.IDENT, "cinco"},
		{token.COMMA, ","},
		{token.IDENT, "diez"},
		{token.RPAREN, ")"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.STAR, "*"},
		{token.INT, "5"},
		{token.NEW_LINE, "\n"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.IF, "si"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.THEN, "entonces"},
		{token.NEW_LINE, "\n"},
		{token.RETURN, "devolver"},
		{token.TRUE, "verdadero"},
		{token.NEW_LINE, "\n"},
		{token.ELSE, "sino"},
		{token.NEW_LINE, "\n"},
		{token.RETURN, "devolver"},
		{token.FALSE, "falso"},
		{token.NEW_LINE, "\n"},
		{token.END, "fin"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.NEW_LINE, "\n"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.NEW_LINE, "\n"},
		{token.STRING, "foobar"},
		{token.NEW_LINE, "\n"},
		{token.STRING, "foo bar"},
		{token.NEW_LINE, "\n"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.NEW_LINE, "\n"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.NEW_LINE, "\n"},
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

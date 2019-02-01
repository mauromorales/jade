package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	STAR   = "*"
	SLASH  = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// DELIMITERS
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	NEW_LINE  = "\n"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	DEFINE   = "DEFINE"
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	THEN     = "THEN"
	ELSE     = "ELSE"
	END      = "END"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"definir":   DEFINE,
	"funcion":   FUNCTION,
	"verdadero": TRUE,
	"falso":     FALSE,
	"si":        IF,
	"entonces":  THEN,
	"sino":      ELSE,
	"fin":       END,
	"devolver":  RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

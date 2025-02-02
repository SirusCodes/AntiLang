package lexer

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ       = "=="
	NOT_EQ   = "!="
	LT_EQ    = "<="
	GT_EQ    = ">="
	PLUS_EQ  = "+="
	MINUS_EQ = "-="

	LOG_AND = "&&"
	LOG_OR  = "||"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	LSQBRAC = "["
	RSQBRAC = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	FOR      = "FOR"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"for":    FOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

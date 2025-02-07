package lexer_test

import (
	"testing"

	. "github.com/SirusCodes/anti-lang/src/lexer"
)

func TestNextToken(t *testing.T) {
	input := `=+(){}[],;
	<= >= == !=
	5 = five let,
	10 = ten let,

	2 * 5 == 10,
	2 * 5 != 11,
	10 / 5 == 2,

	func{} main (
		5 = num let,	
		{num == 5} if (
			0 return,
		) else (
			1 return, 
		)
	)

	{i += 1, i < num, 0 = i let} for (
		
	)

	&& ||
`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{ASSIGN, "="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{LSQBRAC, "["},
		{RSQBRAC, "]"},
		{COMMA, ","},
		{SEMICOLON, ";"},

		{LT_EQ, "<="},
		{GT_EQ, ">="},
		{EQ, "=="},
		{NOT_EQ, "!="},

		{INT, "5"},
		{ASSIGN, "="},
		{IDENT, "five"},
		{LET, "let"},
		{COMMA, ","},

		{INT, "10"},
		{ASSIGN, "="},
		{IDENT, "ten"},
		{LET, "let"},
		{COMMA, ","},

		{INT, "2"},
		{ASTERISK, "*"},
		{INT, "5"},
		{EQ, "=="},
		{INT, "10"},
		{COMMA, ","},

		{INT, "2"},
		{ASTERISK, "*"},
		{INT, "5"},
		{NOT_EQ, "!="},
		{INT, "11"},
		{COMMA, ","},

		{INT, "10"},
		{SLASH, "/"},
		{INT, "5"},
		{EQ, "=="},
		{INT, "2"},
		{COMMA, ","},

		{FUNCTION, "func"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{IDENT, "main"},
		{LPAREN, "("},
		{INT, "5"},
		{ASSIGN, "="},
		{IDENT, "num"},
		{LET, "let"},
		{COMMA, ","},
		{LBRACE, "{"},
		{IDENT, "num"},
		{EQ, "=="},
		{INT, "5"},
		{RBRACE, "}"},
		{IF, "if"},
		{LPAREN, "("},
		{INT, "0"},
		{RETURN, "return"},
		{COMMA, ","},
		{RPAREN, ")"},
		{ELSE, "else"},
		{LPAREN, "("},
		{INT, "1"},
		{RETURN, "return"},
		{COMMA, ","},
		{RPAREN, ")"},
		{RPAREN, ")"},

		{LBRACE, "{"},
		{IDENT, "i"},
		{PLUS_EQ, "+="},
		{INT, "1"},
		{COMMA, ","},
		{IDENT, "i"},
		{LT, "<"},
		{IDENT, "num"},
		{COMMA, ","},
		{INT, "0"},
		{ASSIGN, "="},
		{IDENT, "i"},
		{LET, "let"},
		{RBRACE, "}"},
		{FOR, "for"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LOG_AND, "&&"},
		{LOG_OR, "||"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSaveRestoreLexer(t *testing.T) {
	input := "=+(){}[],;"

	l := New(input)

	tok := l.NextToken()
	if tok.Type != ASSIGN {
		t.Fatalf("expected =, got %q", tok.Literal)
	}

	l.MoveReaderForTemp(func() {
		tok = l.NextToken()
		if tok.Type != PLUS {
			t.Fatalf("expected +, got %q", tok.Literal)
		}

		tok = l.NextToken()
		if tok.Type != LPAREN {
			t.Fatalf("expected (, got %q", tok.Literal)
		}
	})

	tok = l.NextToken()
	if tok.Type != PLUS {
		t.Fatalf("expected +, got %q", tok.Literal)
	}
}

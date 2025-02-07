package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

var (
	tempPosition     int
	tempReadPosition int
	tempCh           byte
)

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = l.makeTwoCharToken(EQ, ASSIGN)
	case '+':
		tok = l.makeTwoCharToken(PLUS_EQ, PLUS)
	case '-':
		tok = l.makeTwoCharToken(MINUS_EQ, MINUS)
	case '!':
		tok = l.makeTwoCharToken(NOT_EQ, BANG)
	case '/':
		tok = l.makeTwoCharToken(SLASH_EQ, SLASH)
	case '*':
		tok = l.makeTwoCharToken(ASTER_EQ, ASTERISK)
	case '<':
		tok = l.makeTwoCharToken(LT_EQ, LT)
	case '>':
		tok = l.makeTwoCharToken(GT_EQ, GT)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '[':
		tok = newToken(LSQBRAC, l.ch)
	case ']':
		tok = newToken(RSQBRAC, l.ch)
	case '%':
		tok = newToken(MOD, l.ch)
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: LOG_AND, Literal: literal}
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: LOG_OR, Literal: literal}
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) MoveReaderForTemp(fn func()) {
	l.saveTokenState()
	defer l.restoreTokenState()

	fn()
}

func (l *Lexer) saveTokenState() {
	tempPosition = l.position
	tempReadPosition = l.readPosition
	tempCh = l.ch
}

func (l *Lexer) restoreTokenState() {
	l.position = tempPosition
	l.readPosition = tempReadPosition
	l.ch = tempCh

	tempCh = 0
	tempPosition = 0
	tempReadPosition = 0
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) makeTwoCharToken(twoCharTokenType TokenType, oneCharTokenType TokenType) Token {
	if l.peekChar() == '=' {
		ch := l.ch
		l.readChar()
		literal := string(ch) + string(l.ch)
		return Token{Type: twoCharTokenType, Literal: literal}
	} else {
		return newToken(oneCharTokenType, l.ch)
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

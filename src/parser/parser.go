package parser

import (
	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/lexer"
)

var (
	tempCurToken  lexer.Token
	tempPeekToken lexer.Token
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // + or -
	PRODUCT     // * or /
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

// Defination of parser functions
type (
	infixParseFns  map[lexer.TokenType]func(ast.Expression) ast.Expression
	prefixParseFns map[lexer.TokenType]func() ast.Expression
)

type Parser struct {
	lexer *lexer.Lexer

	curToken  lexer.Token
	peekToken lexer.Token

	infixParseFns  infixParseFns
	prefixParseFns prefixParseFns

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{lexer: l, errors: []string{}}

	// To set both curToken and peekToken
	parser.nextToken()
	parser.nextToken()

	// All prefix parse functions
	parser.prefixParseFns = make(prefixParseFns)
	parser.registerPrefix(lexer.IDENT, parser.parseIdentifier)
	parser.registerPrefix(lexer.INT, parser.parseIntegerLiteral)

	// All infix parse functions
	parser.infixParseFns = make(infixParseFns)

	return parser
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) registerPrefix(tokenType lexer.TokenType, fn func() ast.Expression) {
	parser.prefixParseFns[tokenType] = fn
}

func (parser *Parser) registerInfix(tokenType lexer.TokenType, fn func(ast.Expression) ast.Expression) {
	parser.infixParseFns[tokenType] = fn
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !parser.curTokenIs(lexer.EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}

	return program
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) addGenericError(message string) {
	parser.errors = append(parser.errors, message)
}

func (parser *Parser) addError(t lexer.TokenType) {
	msg := "parser: expected next token to be %s, got %s instead"
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) curTokenIs(t lexer.TokenType) bool {
	return parser.curToken.Type == t
}

func (parser *Parser) peekTokenIs(t lexer.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) peekTokenTemp(fn func()) {
	tempCurToken = parser.curToken
	tempPeekToken = parser.peekToken

	parser.lexer.MoveReaderForTemp(fn)

	parser.curToken = tempCurToken
	parser.peekToken = tempPeekToken
}

func (parser *Parser) peekTokenAndNext(t lexer.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	} else {
		parser.addError(t)
		return false
	}
}

func (parser *Parser) curTokenIsAndNext(t lexer.TokenType) bool {
	if parser.curTokenIs(t) {
		parser.nextToken()
		return true
	} else {
		parser.addError(t)
		return false
	}
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
	case lexer.COMMA:
		return parser.parseStatementByComma()

	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
}

package parser

import (
	"fmt"

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

var precedences = map[lexer.TokenType]int{
	lexer.EQ:       EQUALS,
	lexer.NOT_EQ:   EQUALS,
	lexer.LT:       LESSGREATER,
	lexer.GT:       LESSGREATER,
	lexer.PLUS:     SUM,
	lexer.MINUS:    SUM,
	lexer.SLASH:    PRODUCT,
	lexer.ASTERISK: PRODUCT,
	lexer.LBRACE:   CALL,
}

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
	parser.registerPrefix(lexer.BANG, parser.parsePrefixExpression)
	parser.registerPrefix(lexer.MINUS, parser.parsePrefixExpression)
	parser.registerPrefix(lexer.TRUE, parser.parseBoolean)
	parser.registerPrefix(lexer.FALSE, parser.parseBoolean)
	parser.registerPrefix(lexer.LBRACE, parser.parseLBraceExpression)

	// All infix parse functions
	parser.infixParseFns = make(infixParseFns)
	parser.registerInfix(lexer.LOG_AND, parser.parseInfixExpression)
	parser.registerInfix(lexer.ASTERISK, parser.parseInfixExpression)
	parser.registerInfix(lexer.EQ, parser.parseInfixExpression)
	parser.registerInfix(lexer.GT, parser.parseInfixExpression)
	parser.registerInfix(lexer.LT, parser.parseInfixExpression)
	parser.registerInfix(lexer.MINUS, parser.parseInfixExpression)
	parser.registerInfix(lexer.MOD, parser.parseInfixExpression)
	parser.registerInfix(lexer.NOT_EQ, parser.parseInfixExpression)
	parser.registerInfix(lexer.LOG_OR, parser.parseInfixExpression)
	parser.registerInfix(lexer.PLUS, parser.parseInfixExpression)
	parser.registerInfix(lexer.SLASH, parser.parseInfixExpression)

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
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, parser.peekToken.Type)
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
	case lexer.LSQBRAC:
		return parser.parseBlockStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
}

func (parser *Parser) peekPrecedence() int {
	if p, ok := precedences[parser.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (parser *Parser) curPrecedence() int {
	if p, ok := precedences[parser.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

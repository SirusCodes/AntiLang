package parser

import (
	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/lexer"
)

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: parser.curToken}

	stmt.Expression = parser.parseExpression(LOWEST, lexer.COMMA)

	return stmt
}

func (parser *Parser) parseStatementByComma() ast.Statement {
	parser.nextToken()

	var token lexer.Token
	isAssign := false
	parser.peekTokenTemp(func() {
		for parser.curToken.Type != lexer.LET && parser.curToken.Type != lexer.RETURN && parser.curToken.Type != lexer.EOF && !parser.curTokenIs(lexer.COMMA) {
			isAssign = parser.isCurTokenAny(lexer.ASSIGN, lexer.ASTER_EQ, lexer.PLUS_EQ, lexer.MINUS_EQ, lexer.SLASH_EQ) || isAssign
			parser.nextToken()
		}

		token = parser.curToken
	})

	switch token.Type {
	case lexer.LET:
		return parser.parseLetStatement()
	case lexer.RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := &ast.LetStatement{}

	letStatement.Value = parser.parseExpression(LOWEST, lexer.ASSIGN)

	// move from '=' to var name
	parser.nextToken()
	parser.nextToken()

	if !parser.curTokenIs(lexer.IDENT) {
		parser.addGenericError("parser: expected token to be IDENT, got " + parser.curToken.Literal + " instead")
		return nil
	}
	letStatement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}

	// move from var name to 'let'
	parser.nextToken()

	letStatement.Token = parser.curToken

	return letStatement
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	returnStatement := &ast.ReturnStatement{}

	returnStatement.ReturnValue = parser.parseExpression(LOWEST, lexer.RETURN)
	parser.nextToken()

	returnStatement.Token = parser.curToken

	return returnStatement
}

func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: parser.curToken}

	parser.nextToken()

	for !parser.curTokenIs(lexer.RSQBRAC) && !parser.curTokenIs(lexer.EOF) {
		stmt := parser.parseStatement()
		block.Statements = append(block.Statements, stmt)
		parser.nextToken()
	}

	return block
}

package parser

import (
	"fmt"
	"strconv"

	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/lexer"
)

func (parser *Parser) parseExpression(precedence int, endToken lexer.TokenType) ast.Expression {

	prefix := parser.prefixParseFns[parser.curToken.Type]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s", parser.curToken.Type)
		parser.errors = append(parser.errors, msg)
		return nil
	}
	leftExp := prefix()

	for !parser.peekTokenIs(endToken) && precedence < parser.peekPrecedence() && !parser.peekTokenIs(lexer.EOF) {
		infix := parser.infixParseFns[parser.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		parser.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (parser *Parser) parsePrefixExpression() ast.Expression {
	pe := &ast.PrefixExpression{
		Token:    parser.curToken,
		Operator: parser.curToken.Literal,
	}

	parser.nextToken()
	pe.Right = parser.parseExpression(PREFIX, lexer.SEMICOLON)

	return pe
}

func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	ie := &ast.InfixExpression{
		Token:    parser.curToken,
		Left:     left,
		Operator: parser.curToken.Literal,
	}

	precedence := parser.curPrecedence()
	parser.nextToken()
	ie.Right = parser.parseExpression(precedence, lexer.SEMICOLON)

	return ie
}

func (parser *Parser) parseIntegerLiteral() ast.Expression {
	itl := &ast.IntegerLiteral{Token: parser.curToken}

	value, err := strconv.ParseInt(parser.curToken.Literal, 0, 64)

	if err != nil {
		msg := "could not parse " + parser.curToken.Literal + " as integer"
		parser.errors = append(parser.errors, msg)
		return nil
	}

	itl.Value = value
	return itl
}

func (parser *Parser) parseBoolean() ast.Expression {
	return &ast.BooleanLiteral{Token: parser.curToken, Value: parser.curTokenIs(lexer.TRUE)}
}

func (parser *Parser) parseLBraceExpression() ast.Expression {
	var token lexer.Token
	var isFuncDef bool
	parser.peekTokenTemp(func() {
		for !parser.curTokenIs(lexer.RBRACE) && !parser.curTokenIs(lexer.EOF) {
			parser.nextToken()
		}

		parser.nextToken()

		token = parser.curToken
		isFuncDef = parser.peekTokenIs(lexer.FUNCTION)
	})

	switch token.Type {
	case lexer.IDENT:
		if isFuncDef {
			return parser.parseFunctionExpression()
		}
		return parser.parseCallExpression()
	case lexer.IF:
		return parser.parseIfExpression()
	case lexer.WHILE:
		return parser.parseWhileExpression()
	default:
		return parser.parseGroupedExpression()
	}
}

func (parser *Parser) parseFunctionExpression() ast.Expression {
	fe := &ast.FunctionExpression{}

	fe.Parameters = parser.parseFunctionParameters()

	parser.nextToken()

	fe.Token = parser.curToken

	parser.nextToken()

	if !parser.peekTokenAndNext(lexer.LSQBRAC) {
		parser.addError(lexer.LSQBRAC)
		return nil
	}

	fe.Body = parser.parseBlockStatement()

	return fe
}

func (parser *Parser) parseFunctionParameters() []*ast.Identifier {
	var identifiers []*ast.Identifier

	if parser.peekTokenIs(lexer.RBRACE) {
		parser.nextToken()
		return identifiers
	}

	parser.nextToken()

	ident := &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
	identifiers = append(identifiers, ident)

	for parser.peekTokenIs(lexer.SEMICOLON) && !parser.peekTokenIs(lexer.EOF) {
		parser.nextToken()
		parser.nextToken()
		ident := &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !parser.peekTokenAndNext(lexer.RBRACE) {
		return nil
	}

	return identifiers
}

func (parser *Parser) parseIfExpression() ast.Expression {
	ifExp := &ast.IfExpression{}

	if !parser.curTokenIsAndNext(lexer.LBRACE) {
		parser.addError(lexer.LBRACE)
		return nil
	}

	ifExp.Condition = parser.parseExpression(LOWEST, lexer.RBRACE)

	parser.nextToken()

	if !parser.peekTokenAndNext(lexer.IF) {
		parser.addError(lexer.IF)
		return nil
	}

	ifExp.Token = parser.curToken

	if !parser.peekTokenIs(lexer.LSQBRAC) {
		parser.addError(lexer.LSQBRAC)
		return nil
	}

	parser.nextToken()

	ifExp.Consequence = parser.parseBlockStatement()

	if parser.peekTokenIs(lexer.ELSE) {
		parser.nextToken()

		if !parser.peekTokenAndNext(lexer.LSQBRAC) {
			parser.addError(lexer.LSQBRAC)
			return nil
		}

		ifExp.Alternative = parser.parseBlockStatement()
	}

	return ifExp
}

func (parser *Parser) parseCallExpression() ast.Expression {
	ce := &ast.CallExpression{}
	ce.Arguments = parser.parseExpressionList(lexer.RBRACE)
	parser.nextToken()
	ce.Function = parser.parseIdentifier().(*ast.Identifier)
	ce.Token = parser.curToken
	return ce
}

func (parser *Parser) parseExpressionList(end lexer.TokenType) []ast.Expression {
	var list []ast.Expression

	if parser.peekTokenIs(end) {
		parser.nextToken()
		return list
	}

	parser.nextToken()
	list = append(list, parser.parseExpression(LOWEST, lexer.SEMICOLON))

	for parser.peekTokenIs(lexer.SEMICOLON) && !parser.peekTokenIs(lexer.EOF) {
		parser.nextToken()
		parser.nextToken()
		list = append(list, parser.parseExpression(LOWEST, lexer.SEMICOLON))
	}

	if !parser.peekTokenAndNext(end) {
		return nil
	}

	return list
}

func (parser *Parser) parseGroupedExpression() ast.Expression {
	parser.nextToken()

	exp := parser.parseExpression(LOWEST, lexer.SEMICOLON)

	if !parser.peekTokenAndNext(lexer.RBRACE) {
		return nil
	}

	return exp
}

func (parser *Parser) parseWhileExpression() ast.Expression {
	we := &ast.WhileExpression{}

	if !parser.curTokenIsAndNext(lexer.LBRACE) {
		return nil
	}

	we.Condition = parser.parseExpression(LOWEST, lexer.RBRACE)

	parser.nextToken()
	we.Token = parser.curToken
	parser.nextToken()

	if !parser.peekTokenAndNext(lexer.LSQBRAC) {
		return nil
	}

	we.Body = parser.parseBlockStatement()

	return we
}

func (parser *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: parser.curToken, Value: parser.curToken.Literal}
}

func (parser *Parser) parseLParenExpression() ast.Expression {
	isIndexExp := false

	parser.peekTokenTemp(func() {
		for !parser.curTokenIs(lexer.RPAREN) && !parser.curTokenIs(lexer.EOF) {
			parser.nextToken()
		}

		isIndexExp = parser.peekTokenIs(lexer.IDENT) || parser.peekTokenIs(lexer.LPAREN) || parser.peekTokenIs(lexer.LSQBRAC)
	})

	if isIndexExp {
		return parser.parseIndexExpression()
	}

	return parser.parseArrayLiteral()
}

func (parser *Parser) parseArrayLiteral() ast.Expression {
	al := &ast.ArrayLiteral{Token: parser.curToken}
	al.Elements = parser.parseExpressionList(lexer.RPAREN)
	return al
}

func (parser *Parser) parseIndexExpression() ast.Expression {
	ie := &ast.IndexExpression{Token: parser.curToken}

	parser.nextToken()
	ie.Index = parser.parseExpression(LOWEST, lexer.RPAREN)

	if !parser.peekTokenAndNext(lexer.RPAREN) {
		return nil
	}

	parser.nextToken()

	if parser.curTokenIs(lexer.IDENT) {
		ie.Array = parser.parseIdentifier()
	} else if parser.curTokenIs(lexer.LPAREN) {
		ie.Array = parser.parseArrayLiteral()
	} else if parser.curTokenIs(lexer.LSQBRAC) {
		ie.Array = parser.parseHashLiteral()
	}

	return ie
}

func (parser *Parser) parseHashLiteral() ast.Expression {
	hl := &ast.HashLiteral{Token: parser.curToken}
	hl.Pairs = make(map[ast.Expression]ast.Expression)

	for !parser.peekTokenIs(lexer.RSQBRAC) && !parser.peekTokenIs(lexer.EOF) {
		parser.nextToken()
		key := parser.parseExpression(LOWEST, lexer.ASSIGN)

		if !parser.peekTokenAndNext(lexer.ASSIGN) {
			return nil
		}

		parser.nextToken()
		value := parser.parseExpression(LOWEST, lexer.SEMICOLON)

		hl.Pairs[key] = value

		if !parser.peekTokenIs(lexer.RSQBRAC) && !parser.peekTokenAndNext(lexer.SEMICOLON) {
			return nil
		}
	}

	if !parser.peekTokenAndNext(lexer.RSQBRAC) {
		return nil
	}

	return hl
}

func (parser *Parser) parseAssignExpression(value ast.Expression) ast.Expression {
	ae := &ast.AssignExpression{}

	ae.Value = value

	ae.Operator = parser.curToken.Literal

	parser.nextToken()
	ae.Name = parser.parseIdentifier().(*ast.Identifier)
	ae.Token = parser.curToken

	return ae
}

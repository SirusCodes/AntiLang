package parser

import (
	"fmt"
	"strconv"

	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/lexer"
)

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.curToken.Type]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s", parser.curToken.Type)
		parser.errors = append(parser.errors, msg)
		return nil
	}
	leftExp := prefix()

	for !parser.peekTokenIs(lexer.SEMICOLON) && precedence < parser.peekPrecedence() {
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
	pe.Right = parser.parseExpression(PREFIX)

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
	ie.Right = parser.parseExpression(precedence)

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
	return &ast.Boolean{Token: parser.curToken, Value: parser.curTokenIs(lexer.TRUE)}
}

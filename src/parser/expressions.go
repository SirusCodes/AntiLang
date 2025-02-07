package parser

import (
	"strconv"

	"github.com/SirusCodes/anti-lang/src/ast"
)

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	return nil
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

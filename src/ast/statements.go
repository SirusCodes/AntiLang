package ast

import "github.com/SirusCodes/anti-lang/src/lexer"

// EXPRESSION statement

type ExpressionStatement struct {
	Token      lexer.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

//  LET statement

type LetStatement struct {
	Token lexer.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out string

	out += ","
	if ls.Value != nil {
		out += ls.Value.String()
	}
	out += " = "
	out += ls.TokenLiteral() + " "
	out += ls.Name.String()

	return out
}

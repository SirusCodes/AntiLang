package ast

import "github.com/SirusCodes/anti-lang/src/lexer"

// EXPRESSION statement
type ExpressionStatement struct {
	Statement
	Token      lexer.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

//  LET statement

type LetStatement struct {
	Statement
	Token lexer.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out string

	out += ","
	if ls.Value != nil {
		out += ls.Value.String()
	}
	out += " = "
	out += ls.Name.String() + " "
	out += ls.TokenLiteral()

	return out
}

// RETURN statement
type ReturnStatement struct {
	Statement
	Token       lexer.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out string

	out += ","
	if rs.ReturnValue != nil {
		out += rs.ReturnValue.String() + " "
	}
	out += rs.TokenLiteral()

	return out
}

// BLOCK statement
type BlockStatement struct {
	Statement
	Token      lexer.Token // the '[' token
	Statements []Statement
}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out string

	out += "["
	for _, s := range bs.Statements {
		out += s.String()
	}
	out += "]"

	return out
}

package ast

import (
	"bytes"
	"strings"

	"github.com/SirusCodes/anti-lang/src/lexer"
)

// CallExpression represents a function call expression
type CallExpression struct {
	Token     lexer.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var args []string

	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}

	return "(" + "{" + strings.Join(args, ";") + "}" + ce.Function.String() + ")"
}

// InfixExpression represents an infix expression
type InfixExpression struct {
	Token    lexer.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Operator + " " + ie.Right.String() + ")"
}

// PrefixExpression represents a prefix expression
type PrefixExpression struct {
	Token    lexer.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	return "(" + pe.Operator + pe.Right.String() + ")"
}

// IfExpression represents an if expression
type IfExpression struct {
	Token       lexer.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) expressionNode() {}

func (i *IfExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString(i.Condition.String())
	out.WriteString("if")
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())

	if i.Alternative != nil {
		out.WriteString("else")
		out.WriteString(i.Alternative.String())
	}

	return out.String()
}

// FunctionExpression represents a function expression
type FunctionExpression struct {
	Token      lexer.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fe *FunctionExpression) expressionNode() {}

func (fe *FunctionExpression) TokenLiteral() string {
	return fe.Token.Literal
}

func (fe *FunctionExpression) String() string {
	var out bytes.Buffer
	var params []string

	for _, p := range fe.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(params, "; "))
	out.WriteString("}")
	out.WriteString(fe.TokenLiteral())
	out.WriteString("func")
	out.WriteString(fe.Body.String())

	return out.String()
}

// WhileExpression represents a while expression
type WhileExpression struct {
	Token     lexer.Token
	Condition Expression
	Body      *BlockStatement
}

func (we *WhileExpression) expressionNode() {}

func (we *WhileExpression) TokenLiteral() string {
	return we.Token.Literal
}

func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("{")
	out.WriteString(we.Condition.String())
	out.WriteString("}")
	out.WriteString("while")
	out.WriteString(we.Body.String())

	return out.String()
}

// IntegerLiteral represents an integer literal
type IntegerLiteral struct {
	Token lexer.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// Boolean represents a boolean literal
type Boolean struct {
	Token lexer.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

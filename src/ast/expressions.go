package ast

import (
	"bytes"
	"strings"

	"github.com/SirusCodes/anti-lang/src/lexer"
)

// Identifier is the AST node that represents an identifier
type Identifier struct {
	Expression
	Token lexer.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// CallExpression represents a function call expression
type CallExpression struct {
	Expression
	Token     lexer.Token
	Function  Expression
	Arguments []Expression
}

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
	Expression
	Token    lexer.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Operator + " " + ie.Right.String() + ")"
}

// PrefixExpression represents a prefix expression
type PrefixExpression struct {
	Expression
	Token    lexer.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	return "(" + pe.Operator + pe.Right.String() + ")"
}

// IfExpression represents an if expression
type IfExpression struct {
	Expression
	Token       lexer.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

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
	Expression
	Token      lexer.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

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
	Expression
	Token     lexer.Token
	Condition Expression
	Body      *BlockStatement
}

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
	Expression
	Token lexer.Token
	Value int64
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// Boolean represents a boolean literal
type Boolean struct {
	Expression
	Token lexer.Token
	Value bool
}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

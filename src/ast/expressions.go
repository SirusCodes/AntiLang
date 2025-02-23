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

// ConditionalExpression represents an if expression
type ConditionalExpression struct {
	Expression
	Token           lexer.Token
	Condition       Expression
	ExecutionBlock  *BlockStatement
	NextConditional *ConditionalExpression
}

func (i *ConditionalExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *ConditionalExpression) String() string {
	var out bytes.Buffer

	out.WriteString(i.Condition.String())
	out.WriteString("if")
	out.WriteString(" ")
	out.WriteString(i.ExecutionBlock.String())

	if i.NextConditional != nil {
		out.WriteString("else")
		out.WriteString(i.NextConditional.String())
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

// AssignExpression represents an assign expression
type AssignExpression struct {
	Expression
	Token    lexer.Token
	Name     *Identifier
	Operator string
	Value    Expression
}

func (ae *AssignExpression) TokenLiteral() string {
	return ae.Token.Literal
}

func (ae *AssignExpression) String() string {
	return ae.Value.String() + " " + ae.Operator + " " + ae.Name.String()
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

// FloatLiteral represents a float literal
type FloatLiteral struct {
	Expression
	Token lexer.Token
	Value float64
}

func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) String() string {
	return fl.Token.Literal
}

// BooleanLiteral represents a boolean literal
type BooleanLiteral struct {
	Expression
	Token lexer.Token
	Value bool
}

func (b *BooleanLiteral) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BooleanLiteral) String() string {
	return b.Token.Literal
}

// StringLiteral represents a string literal
type StringLiteral struct {
	Expression
	Token lexer.Token
	Value string
}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

// ArrayLiteral represents an array literal
type ArrayLiteral struct {
	Expression
	Token    lexer.Token
	Elements []Expression
}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) String() string {
	var elements []string

	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	return "(" + strings.Join(elements, "; ") + ")"
}

// IndexExpression represents an index expression
type IndexExpression struct {
	Expression
	Token lexer.Token
	Array Expression
	Index Expression
}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	return "{" + ie.Array.String() + "(" + ie.Index.String() + ")}"
}

// HashLiteral represents a hash literal
type HashLiteral struct {
	Expression
	Token lexer.Token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) String() string {
	var pairs []string

	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+"="+value.String())
	}

	return "[" + strings.Join(pairs, "; ") + "]"
}

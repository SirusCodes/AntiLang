package ast

import (
	"bytes"
)

// Node is the interface that all nodes in the AST implement
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is the interface that all statement nodes in the AST implement
type Statement interface {
	Node
	statementNode()
}

// Expression is the interface that all expression nodes in the AST implement
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST that the parser produces
type Program struct {
	Node
	Statements []Statement
}

// TokenLiteral returns the token literal of the first statement in the program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns the string representation of the program
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

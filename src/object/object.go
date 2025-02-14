package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/SirusCodes/anti-lang/src/ast"
)

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
)

type ObjectTypes string

type Object interface {
	Type() ObjectTypes
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectTypes { return INTEGER_OBJ }
func (i *Integer) Inspect() string   { return fmt.Sprint(i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectTypes { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string   { return fmt.Sprint(b.Value) }

type Null struct{}

func (n *Null) Type() ObjectTypes { return NULL_OBJ }
func (n *Null) Inspect() string   { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectTypes { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectTypes { return ERROR_OBJ }
func (e *Error) Inspect() string   { return "ERROR: " + e.Message }

type Function struct {
	Name       string
	Token      *ast.Identifier
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectTypes { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(params, "; "))
	out.WriteString("} ")
	out.WriteString(f.Name)
	out.WriteString(" ")
	out.WriteString(f.Token.TokenLiteral())
	out.WriteString(" ")
	out.WriteString("[\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n]")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectTypes { return STRING_OBJ }
func (s *String) Inspect() string   { return s.Value }

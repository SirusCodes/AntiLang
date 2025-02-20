package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
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
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

type ObjectTypes string

type Object interface {
	Type() ObjectTypes
	Inspect() string
}

type Integer struct {
	Hashable
	Value int64
}

func (i *Integer) Type() ObjectTypes { return INTEGER_OBJ }
func (i *Integer) Inspect() string   { return fmt.Sprint(i.Value) }
func (i *Integer) String() string    { return fmt.Sprint(i.Value) }

type Boolean struct {
	Hashable
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
	Hashable
	Value string
}

func (s *String) Type() ObjectTypes { return STRING_OBJ }
func (s *String) Inspect() string   { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectTypes { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string   { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectTypes { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}

	for _, el := range ao.Elements {
		elements = append(elements, el.Inspect())
	}

	out.WriteString("(")
	out.WriteString(strings.Join(elements, "; "))
	out.WriteString(")")

	return out.String()
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectTypes { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}

	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("[")
	out.WriteString(strings.Join(pairs, "; "))
	out.WriteString("]")

	return out.String()
}

type HashKey struct {
	Type  ObjectTypes
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}
func (s *String) HashKey() HashKey {
	h := fnv.New64a()

	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type HashPair struct {
	Key   Object
	Value Object
}

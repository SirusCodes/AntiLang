package object

import "fmt"

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
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

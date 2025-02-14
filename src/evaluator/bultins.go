package evaluator

import "github.com/SirusCodes/anti-lang/src/object"

var builtins = make(map[string]*object.Builtin)

func registerBuiltIns(name string, fn object.BuiltinFunction) {
	builtins[name] = &object.Builtin{Fn: fn}
}

// Built-in function to get the length
func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

// Registering built-in functions
func init() {
	registerBuiltIns("len", builtinLen)
}

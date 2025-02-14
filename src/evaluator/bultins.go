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
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func builtinFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func builtinLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func builtinRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func builtinPush(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}

func builtinPop(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `pop` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	if length > 0 {
		newElements := make([]object.Object, length-1)
		copy(newElements, arr.Elements)
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func builtinAddAt(args ...object.Object) object.Object {
	if len(args) != 3 {
		return newError("wrong number of arguments. got=%d, want=3", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("first argument to `addAt` must be ARRAY, got %s", args[0].Type())
	}

	if args[1].Type() != object.INTEGER_OBJ {
		return newError("second argument to `addAt` must be INTEGER, got %s", args[1].Type())
	}

	arr := args[0].(*object.Array)
	index := args[1].(*object.Integer).Value
	value := args[2]

	length := len(arr.Elements)
	if index < 0 || index > int64(length) {
		return newError("index out of bounds")
	}

	newElements := make([]object.Object, length+1)
	copy(newElements, arr.Elements[:index])
	newElements[index] = value
	copy(newElements[index+1:], arr.Elements[index:])

	return &object.Array{Elements: newElements}
}

func builtinRemoveAt(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("first argument to `removeAt` must be ARRAY, got %s", args[0].Type())
	}

	if args[1].Type() != object.INTEGER_OBJ {
		return newError("second argument to `removeAt` must be INTEGER, got %s", args[1].Type())
	}

	arr := args[0].(*object.Array)
	index := args[1].(*object.Integer).Value

	length := len(arr.Elements)
	if index < 0 || index >= int64(length) {
		return newError("index out of bounds")
	}

	newElements := make([]object.Object, length-1)
	copy(newElements, arr.Elements[:index])
	copy(newElements[index:], arr.Elements[index+1:])

	return &object.Array{Elements: newElements}
}

func builtinPrint(args ...object.Object) object.Object {
	for _, arg := range args {
		println(arg.Inspect())
	}
	return NULL
}

// Registering built-in functions
func init() {
	registerBuiltIns("len", builtinLen)
	registerBuiltIns("first", builtinFirst)
	registerBuiltIns("last", builtinLast)
	registerBuiltIns("rest", builtinRest)
	registerBuiltIns("push", builtinPush)
	registerBuiltIns("pop", builtinPop)
	registerBuiltIns("addAt", builtinAddAt)
	registerBuiltIns("removeAt", builtinRemoveAt)
	registerBuiltIns("print", builtinPrint)
}

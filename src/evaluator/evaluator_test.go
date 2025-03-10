package evaluator_test

import (
	"testing"

	"github.com/SirusCodes/anti-lang/src/evaluator"
	"github.com/SirusCodes/anti-lang/src/object"
	"github.com/SirusCodes/anti-lang/src/utils"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"--10", 10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * {5 + 10}", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * {3 * 3} + 10", 37},
		{"{5 + 10 * 2 + 15 / 3} * 2 + -10", 50},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.5", 5.5},
		{"10.5", 10.5},
		{"-5.5", -5.5},
		{"--10.5", 10.5},
		{"5.5 + 5.5 + 5.5 + 5.5 - 10.5", 11.5},
		{"2.5 * 2.5 * 2.5 * 2.5 * 2.5", 97.65625},
		{"-50.5 + 100.5 + -50.5", -0.5},
		{"5.5 * 2.5 + 10.5", 24.25},
		{"5.5 + 2.5 * 10.5", 31.75},
		{"20.5 + 2.5 * -10.5", -5.75},
		{"50.5 / 2.5 * 2.5 + 10.5", 61},
		{"2.5 * {5.5 + 10.5}", 40},
		{"3.5 * 3.5 * 3.5 + 10.5", 53.375},
		{"3.5 * {3.5 * 3.5} + 10.5", 53.375},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"{1 < 2} == true", true},
		{"{1 < 2} == false", false},
		{"{1 > 2} == true", false},
		{"{1 > 2} == false", true},
		{"{1 < 2} == {1 > 2}", false},
		{"{1 < 2} != {1 > 2}", true},
		{"{1 < 2} == {1 < 2}", true},
		{"{1 < 2} != {1 < 2}", false},
		{"{1 > 2} == {1 > 2}", true},
		{"{1 > 2} != {1 > 2}", false},
		{"2 % 2 == 0", true},
		{"2 % 2 != 0", false},
		{"1 <= 2", true},
		{"2 <= 2", true},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!false", false},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"{true} if [ 10 ]", 10},
		{"{false} if [ 10 ]", nil},
		{"{1} if [ 10 ]", 10},
		{"{1 < 2} if [ 10 ]", 10},
		{"{1 > 2} if [ 10 ]", nil},
		{"{1 > 2} if [ 10 ] else [ 20 ]", 20},
		{"{1 < 2} if [ 10 ] else [ 20 ]", 10},
		{"{1 > 2} if [ 10 ] {1 == 2} if else [ 20 ] else [ 30 ]", 30},
		{"{1 > 2} if [ 10 ] {1 < 2} if else [ 20 ] else [ 30 ]", 20},
		{"{1 < 2} if [ 10 ] {1 > 2} if else [ 20 ] else [ 30 ]", 10},
		{"{1 < 2} if [ 10 ] {1 < 2} if else [ 20 ] else [ 30 ]", 10},
		{"{1 > 2} if [ 10 ] {1 > 2} if else [ 20 ]", nil},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != evaluator.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{",10 return", 10},
		{",2 * 5 return", 10},
		{"9, 2 * 5 return", 10},
		{"{10 > 1} if [,10 return ]", 10},
		{"{10 > 1} if [,10 return ] else [,20 return ]", 10},
		{`{10 > 1} if [
			{10 > 1} if [
				,10 return
			] 
			,20 return
		]`,
			10},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + true", "type mismatch: INTEGER + BOOLEAN"},
		{"-true", "unknown operator: -BOOLEAN"},
		{"true + false", "unknown operator: BOOLEAN + BOOLEAN"},
		{"{10 > 1} if [ true + false ]", "unknown operator: BOOLEAN + BOOLEAN"},
		{`
		{10 > 1} if [
			{10 > 1} if [
				,true + false return
			]
			,10 return
		]
		`, "unknown operator: BOOLEAN + BOOLEAN"},
		{"foobar", "identifier not found: foobar"},
		{`$Hello$ - $World$`, "unknown operator: STRING - STRING"},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T (%+v)", evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expected {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expected, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{",5 = a let\n a", 5},
		{",5 * 5 = a let\na", 25},
		{",5 = a let\n,a = b let\n b", 5},
		{",5 = a let\n,a = b let\n,a + b + 5 = c let\n c", 15},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "{x} abc func [x + 2]"
	parsed := utils.ParseInput(t, input)
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(parsed, env)

	function, ok := env.Get("abc")
	if !ok {
		t.Fatalf("Function 'abc' not in environment")
	}
	fn, ok := function.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}
	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}
	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}
	expectedBody := "[(x + 2)]"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"{x} abc func [x + 2]\n{2}abc", 4},
		{"{x} abc func [x + 2]\n{5}abc", 7},
		{"{x} abc func [x + 2]\n{5 * 5}abc", 27},
		{"{a; b} add func [a + b]\n{2; 3}add", 5},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestStringLiteral(t *testing.T) {
	input := `$Hello World!$`
	evaluated := utils.EvalTest(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`$Hello$ + $ $ + $World!$`, "Hello World!"},
		{`$Hello$ + 1 + $World!$`, "Hello1World!"},
		{`$Hello$ + 1`, "Hello1"},
		{`1 + $Hello$`, "1Hello"},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
		}
		if str.Value != tt.expected {
			t.Errorf("String has wrong value. got=%q", str.Value)
		}
	}
}

func TestBuiltInFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"{$$}len", 0},
		{"{$hello$}len", 5},
		{"{$hello world$}len", 11},
		{"{2}len", "argument to `len` not supported, got INTEGER"},
		{"{true}len", "argument to `len` not supported, got BOOLEAN"},
		{"{$abc$; $def$}len", "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "(1; 2 * 2; 3 + 3)"
	evaluated := utils.EvalTest(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong number of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"(1)(1; 2; 3)", 1},
		{"(2)(1; 2; 3)", 2},
		{"(3)(1; 2; 3)", 3},
		{"(1 + 1 + 1)(1; 2; 3)", 3},
		{",(1; 2; 3) = myArray let\n(1)myArray", 1},
		{",(1; 2; 3) = myArray let\n(2)myArray", 2},
		{",(1; 2; 3) = myArray let\n(3)myArray", 3},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		default:
			testNullObject(t, evaluated)
		}
	}
}

func TestArrayIndexOutOfBounds(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"(4)(1; 2; 3)", "index out of bounds"},
		{"(0)(1; 2; 3)", "come on, you know arrays are 1-indexed"},
		{"(-1)(1; 2; 3)", "index out of bounds"},
	}

	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Fatalf("object is not Error. got=%T (%+v)", evaluated, evaluated)
		}
		if errObj.Message != tt.expected {
			t.Errorf("wrong error message. expected=%q, got=%q", "index out of bounds", errObj.Message)
		}
	}
}

func TestHashLiterals(t *testing.T) {
	input := "[1= 2; 2= 3]"
	evaluated := utils.EvalTest(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("object is not Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.Integer{Value: 1}).HashKey(): 2,
		(&object.Integer{Value: 2}).HashKey(): 3,
	}

	if len(result.Pairs) != 2 {
		t.Fatalf("hash has wrong number of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}
		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"(1)[1= 2; 2= 3]", 2},
		{"(2)[1= 2; 2= 3]", 3},
		{"(1 + 1)[1= 2; 2= 3]", 3},
		{",[1= 2; 2= 3] = myHash let\n(1)myHash", 2},
		{",[1= 2; 2= 3] = myHash let\n(2)myHash", 3},
		{",[1= 2; 2= 3] = myHash let\n(3)myHash", nil},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		default:
			testNullObject(t, evaluated)
		}
	}
}

func TestFuncCallAssignment(t *testing.T) {
	input := `{a; b} add func [
    ,a + b return
]

,{2; 4}add = res let
res`
	evaluated := utils.EvalTest(input)
	testIntegerObject(t, evaluated, 6)
}

func TestReassignmentOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{",5 = a let\n,5 += a\na", 10},
		{",5 = a let\n,5 -= a\na", 0},
		{",5 = a let\n,5 *= a\na", 25},
		{",5 = a let\n,5 /= a\na", 1},
	}
	for _, tt := range tests {
		evaluated := utils.EvalTest(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestWhileExpression(t *testing.T) {
	input := `,0 = x let

{x < 5} while [
	,1 += x
]

x
`
	evaluated := utils.EvalTest(input)
	testIntegerObject(t, evaluated, 5)
}

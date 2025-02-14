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

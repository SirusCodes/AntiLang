package parser_test

import (
	"fmt"
	"testing"

	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/utils"
)

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}
	for _, tt := range prefixTests {
		program := utils.ParseInput(t, tt.input)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestParsingInfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5+5", 5, "+", 5},
		{"5-5", 5, "-", 5},
		{"5*5", 5, "*", 5},
		{"5/5", 5, "/", 5},
		{"5%5", 5, "%", 5},
		{"5>5", 5, ">", 5},
		{"5<5", 5, "<", 5},
		{"5==5", 5, "==", 5},
		{"5!=5", 5, "!=", 5},
		{"true==true", true, "==", true},
		{"true!=false", true, "!=", false},
		{"1.3 + 1.3", 1.3, "+", 1.3},
	}
	for _, tt := range infixTests {
		program := utils.ParseInput(t, tt.input)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		if !testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a+b+c", "((a + b) + c)"},
		{"a+b-c", "((a + b) - c)"},
		{"a*b*c", "((a * b) * c)"},
		{"a*b/c", "((a * b) / c)"},
		{"a+b/c", "(a + (b / c))"},
		{"a+b*c+d/e-f", "(((a + (b * c)) + (d / e)) - f)"},
		{"5>4==3<4", "((5 > 4) == (3 < 4))"},
		{"5<4!=3>4", "((5 < 4) != (3 > 4))"},
		{"3+4*5==3*1+4*5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3>5==false", "((3 > 5) == false)"},
		{"3<5==true", "((3 < 5) == true)"},
		{"1+{2+3}+4", "((1 + (2 + 3)) + 4)"},
		{"{5+5}*2", "((5 + 5) * 2)"},
		{"2/{5+5}", "(2 / (5 + 5))"},
		{"-{5+5}", "(-(5 + 5))"},
		{"5.1 + 0.9", "(5.1 + 0.9)"},
		{"!{true==true}", "(!(true == true))"},
		{"a + {b; c}add + d", "((a + ({b;c}add)) + d)"},
		{"a % b == c % d", "((a % b) == (c % d))"},
		{"a <= b", "(a <= b)"},
	}
	for _, tt := range tests {
		program := utils.ParseInput(t, tt.input)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestIfExpression(t *testing.T) {
	input := "{x < y} if [ b ]"
	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.ConditionalExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IfExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.ExecutionBlock.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.ExecutionBlock.Statements))
	}

	consequence, ok := exp.ExecutionBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "b") {
		return
	}

	if exp.NextConditional != nil {
		t.Errorf("exp.NextConditional was not nil. got=%+v", exp.NextConditional)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := "{x < y} if [ b ] else [ c ]"
	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.ConditionalExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.ConditionalExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.ExecutionBlock.Statements) != 1 {
		t.Errorf("ExecutionBlock is not 1 statements. got=%d\n", len(exp.ExecutionBlock.Statements))
	}

	consequence, ok := exp.ExecutionBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "b") {
		return
	}

	if exp.NextConditional == nil {
		t.Fatal("exp.NextConditional was nil.")
	}

	if len(exp.NextConditional.ExecutionBlock.Statements) != 1 {
		t.Fatalf("ExecutionBlock.Statements is not 1 statements. got=%d\n", len(exp.NextConditional.ExecutionBlock.Statements))
	}

	alternative, ok := exp.NextConditional.ExecutionBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.NextConditional.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, alternative.Expression, "c") {
		return
	}
}

func TestIfElseLadderExpression(t *testing.T) {
	input := "{x < y} if [ b ] {x > y} if else [ c ] else [ d ]"
	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.ConditionalExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.ConditionalExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.ExecutionBlock.Statements) != 1 {
		t.Errorf("ExecutionBlock is not 1 statements. got=%d\n", len(exp.ExecutionBlock.Statements))
	}

	ifBlock, ok := exp.ExecutionBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, ifBlock.Expression, "b") {
		return
	}

	if exp.NextConditional == nil {
		t.Fatal("exp.NextConditional was nil.")
	}

	if !testInfixExpression(t, exp.NextConditional.Condition, "x", ">", "y") {
		return
	}

	if len(exp.NextConditional.ExecutionBlock.Statements) != 1 {
		t.Fatalf("ExecutionBlock.Statements is not 1 statements. got=%d\n", len(exp.NextConditional.ExecutionBlock.Statements))
	}

	elseIfBlock, ok := exp.NextConditional.ExecutionBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.NextConditional.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, elseIfBlock.Expression, "c") {
		return
	}

	if exp.NextConditional.NextConditional == nil {
		t.Fatal("exp.NextConditional.NextConditional was nil.")
	}

	if len(exp.NextConditional.NextConditional.ExecutionBlock.Statements) != 1 {
		t.Fatalf("ExecutionBlock.Statements is not 1 statements. got=%d\n", len(exp.NextConditional.NextConditional.ExecutionBlock.Statements))
	}

	elseBlock, ok := exp.NextConditional.NextConditional.ExecutionBlock.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.NextConditional.NextConditional.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, elseBlock.Expression, "d") {
		return
	}

	if exp.NextConditional.NextConditional.NextConditional != nil {
		t.Fatal("exp.NextConditional.NextConditional.NextConditional was not nil.")
	}
}

func TestFunctionExpressionParsing(t *testing.T) {
	input := `{x; y}add func [
		,x + y return
	]`

	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionExpression. got=%T", stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))
	}

	if function.TokenLiteral() != "add" {
		t.Fatalf("function.TokenLiteral is not 'add'. got=%s", function.TokenLiteral())
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n", len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ReturnStatement. got=%T", function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.ReturnValue, "x", "+", "y")
}

func TestWhileExpressionParsing(t *testing.T) {
	input := "{x < y} while [ ,b return ]"

	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	whileExp, ok := stmt.Expression.(*ast.WhileExpression)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.WhileExpression. got=%T", stmt.Expression)
	}

	if !testInfixExpression(t, whileExp.Condition, "x", "<", "y") {
		return
	}

	if len(whileExp.Body.Statements) != 1 {
		t.Fatalf("whileExp.Body.Statements has not 1 statements. got=%d\n", len(whileExp.Body.Statements))
	}

	bodyStmt, ok := whileExp.Body.Statements[0].(*ast.ReturnStatement)

	if !ok {
		t.Fatalf("while body stmt is not ast.ReturnStatement. got=%T", whileExp.Body.Statements[0])
	}

	testIdentifier(t, bodyStmt.ReturnValue, "b")
}

func TestStringLiteralExpression(t *testing.T) {
	input := `$hello world$`

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.StringLiteral)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.StringLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != "hello world" {
		t.Errorf("literal.Value not %q. got=%q", "hello world", literal.Value)
	}
}

func TestArrayLiteralExpression(t *testing.T) {
	input := "(1; 2 * 2; 3 + 3)"

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	array, ok := stmt.Expression.(*ast.ArrayLiteral)

	if !ok {
		t.Fatalf("stmt.Expression is not ast.ArrayLiteral. got=%T", stmt.Expression)
	}

	if len(array.Elements) != 3 {
		t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
	}

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestIndexExpression(t *testing.T) {
	input := "(1 + 1)myArray"

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IndexExpression. got=%T", stmt.Expression)
	}

	if !testIdentifier(t, indexExp.Array, "myArray") {
		return
	}

	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
		return
	}
}

func TestIndexOnArrayLiteral(t *testing.T) {
	input := "(0)(1; 1)"

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.IndexExpression. got=%T", stmt.Expression)
	}

	array, ok := indexExp.Array.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("indexExp.Array is not ast.ArrayLiteral. got=%T", indexExp.Array)
	}

	if len(array.Elements) != 2 {
		t.Fatalf("len(array.Elements) not 2. got=%d", len(array.Elements))
	}

	if !testIntegerLiteral(t, indexExp.Index, 0) {
		return
	}
}

func TestParsingHashLiteralsStringKeys(t *testing.T) {
	input := "[$one$ = 1; $two$ = 2; $three$ = 3]"

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := stmt.Expression.(*ast.HashLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.HashLiteral. got=%T", stmt.Expression)
	}

	expected := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if len(hash.Pairs) != len(expected) {
		t.Fatalf("hash.Pairs has wrong number of elements. got=%d", len(hash.Pairs))
	}

	for key, value := range hash.Pairs {
		strKey, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
		}

		expectedValue := expected[strKey.Value]
		testIntegerLiteral(t, value, expectedValue)
	}
}

func TestParsingEmptyHashLiteral(t *testing.T) {
	input := "[]"

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := stmt.Expression.(*ast.HashLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.HashLiteral. got=%T", stmt.Expression)
	}

	if len(hash.Pairs) != 0 {
		t.Errorf("hash.Pairs has wrong number of elements. got=%d", len(hash.Pairs))
	}
}

func TestParsingHashLiteralsWithExpressions(t *testing.T) {
	input := "[$one$ = 0 + 1; $two$ = 10 - 8; $three$ = 15 / 5]"

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	hash, ok := stmt.Expression.(*ast.HashLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.HashLiteral. got=%T", stmt.Expression)
	}

	tests := map[string]func(ast.Expression){
		"one": func(e ast.Expression) {
			testInfixExpression(t, e, 0, "+", 1)
		},
		"two": func(e ast.Expression) {
			testInfixExpression(t, e, 10, "-", 8)
		},
		"three": func(e ast.Expression) {
			testInfixExpression(t, e, 15, "/", 5)
		},
	}

	if len(hash.Pairs) != len(tests) {
		t.Fatalf("hash.Pairs has wrong number of elements. got=%d", len(hash.Pairs))
	}

	for key, value := range hash.Pairs {
		strKey, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
		}

		testFunc, ok := tests[strKey.Value]
		if !ok {
			t.Errorf("No test function for key %q found", strKey.Value)
			continue
		}

		testFunc(value)
	}
}

func TestParsingFunctionCallToLetStatement(t *testing.T) {
	input := ",{2; 4}add = res let"

	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.LetStatement. got=%T", program.Statements[0])
	}

	if stmt.Name.Value != "res" {
		t.Errorf("stmt.Name.Value not 'res'. got=%s", stmt.Name.Value)
	}

	callExpression, ok := stmt.Value.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Value is not ast.CallExpression. got=%T", stmt.Value)
	}

	if !testIdentifier(t, callExpression.Function, "add") {
		return
	}

	if len(callExpression.Arguments) != 2 {
		t.Fatalf("wrong length of arguments. got=%d", len(callExpression.Arguments))
	}

	testIntegerLiteral(t, callExpression.Arguments[0], 2)
	testIntegerLiteral(t, callExpression.Arguments[1], 4)
}

func TestParsingCommaCallExpression(t *testing.T) {
	input := ",{1; 2}add"

	program := utils.ParseInput(t, input)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	callExpression, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T", stmt.Expression)
	}

	if !testIdentifier(t, callExpression.Function, "add") {
		return
	}

	if len(callExpression.Arguments) != 2 {
		t.Fatalf("wrong length of arguments. got=%d", len(callExpression.Arguments))
	}

	testIntegerLiteral(t, callExpression.Arguments[0], 1)
	testIntegerLiteral(t, callExpression.Arguments[1], 2)
}

func TestAssignExpression(t *testing.T) {
	tests := []struct {
		input    string
		value    interface{}
		operator string
		name     string
	}{
		{",5 = a", 5, "=", "a"},
		{",5 += a", 5, "+=", "a"},
		{",5 -= a", 5, "-=", "a"},
		{",5 *= a", 5, "*=", "a"},
		{",5 /= a", 5, "/=", "a"},
	}

	for _, tt := range tests {
		program := utils.ParseInput(t, tt.input)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.AssignExpression)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.LetStatement. got=%T", program.Statements[0])
		}

		if !testLiteralExpression(t, stmt.Value, tt.value) {
			return
		}

		if stmt.Operator != tt.operator {
			t.Errorf("stmt.Operator is not '%s'. got=%s", tt.operator, stmt.Operator)
		}

		if stmt.Name.Value != tt.name {
			t.Errorf("stmt.Name.Value is not '%s'. got=%s", tt.name, stmt.Name.Value)
		}
	}
}

func TestIfBlockInFunction(t *testing.T) {
	input := `{x; y}add func [
		,{x < y} if [ ,x return ] else [ ,y return ]
	]`

	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionExpression. got=%T", stmt.Expression)
	}

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements does not contain 1 statements. got=%d", len(function.Body.Statements))
	}

	ifExp, ok := function.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.ConditionalExpression)
	if !ok {
		t.Fatalf("function.Body.Statements[0] is not ast.ExpressionStatement. got=%T", function.Body.Statements[0])
	}

	if !testInfixExpression(t, ifExp.Condition, "x", "<", "y") {
		return
	}

	if len(ifExp.ExecutionBlock.Statements) != 1 {
		t.Fatalf("consequence is not 1 statements. got=%d\n", len(ifExp.ExecutionBlock.Statements))
	}

	consequence, ok := ifExp.ExecutionBlock.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ReturnStatement. got=%T", ifExp.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, consequence.ReturnValue, "x") {
		return
	}

	if len(ifExp.NextConditional.ExecutionBlock.Statements) != 1 {
		t.Fatalf("consequence is not 1 statements. got=%d\n", len(ifExp.NextConditional.ExecutionBlock.Statements))
	}

	alternative, ok := ifExp.NextConditional.ExecutionBlock.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ReturnStatement. got=%T", ifExp.NextConditional.ExecutionBlock.Statements[0])
	}

	if !testIdentifier(t, alternative.ReturnValue, "y") {
		return
	}
}

// HELPER FUNCTIONS
func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false
	}
	return true
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", exp, exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case float64:
		return testFloatLiteral(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bool, ok := exp.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}
	if bool.Value != value {
		t.Errorf("bool.Value not %t. got=%t", value, bool.Value)
		return false
	}
	if bool.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bool.TokenLiteral not %t. got=%s", value, bool.TokenLiteral())
		return false
	}
	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func testFloatLiteral(t *testing.T, fl ast.Expression, value float64) bool {
	float, ok := fl.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("fl not *ast.FloatLiteral. got=%T", fl)
		return false
	}
	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}
	return true
}

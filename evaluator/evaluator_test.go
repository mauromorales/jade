package evaluator

import (
	"testing"

	"github.com/mauromorales/jade/lexer"
	"github.com/mauromorales/jade/object"
	"github.com/mauromorales/jade/parser"
)

func TestIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"verdadero", true},
		{"falso", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"verdadero == verdadero", true},
		{"falso == falso", true},
		{"verdadero == falso", false},
		{"verdadero != falso", true},
		{"falso != verdadero", true},
		{"(1 < 2) == verdadero", true},
		{"(1 < 2) == falso", false},
		{"(1 > 2) == verdadero", false},
		{"(1 > 2) == falso", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!verdadero", false},
		{"!falso", true},
		{"!!verdadero", true},
		{"!!falso", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"si (verdadero) { 10 }", 10},
		{"si (falso) { 10 }", nil},
		{"si (1) { 10 }", 10},
		{"si (1 < 2) { 10 }", 10},
		{"si (1 > 2) { 10 }", nil},
		//{"si (1 > 2) { 10 } sino { 20 }", 20},
		//{"si (1 < 2) { 10 } sino { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"devolver 10", 10},
		{"devolver 10\n9;", 10},
		{"devolver 2 * 5\n9;", 10},
		{"9\n devolver 2 * 5\n 9;", 10},
		{`
si (10 > 1) {
	si (10 > 1) {
		devolver 10
	}

	devolver 1
}
`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + verdadero\n",
			"tipos no coinciden: ENTERO + BOOLEANO",
		},
		{
			"5 + verdadero\n 5;",
			"tipos no coinciden: ENTERO + BOOLEANO",
		},
		{
			"-verdadero",
			"operador desconocido: -BOOLEANO",
		},
		{
			"verdadero + falso\n",
			"operador desconocido: BOOLEANO + BOOLEANO",
		},
		{
			"5\nverdadero + falso\n5",
			"operador desconocido: BOOLEANO + BOOLEANO",
		},
		{
			"si (10 > 1) { verdadero + falso }",
			"operador desconocido: BOOLEANO + BOOLEANO",
		},
		{`
si (10 > 1) {
	si (10 > 1) {
		devolver verdadero + falso
	}

	devolver 1
}
`,
			"operador desconocido: BOOLEANO + BOOLEANO",
		},
		{
			"foobar",
			"identificador no encontrado: foobar",
		},
		{
			`"Hello" - "World"`,
			"operador desconocido: CADENA_DE_CARACTERES - CADENA_DE_CARACTERES",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}

		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"definir a = 5\n a\n", 5},
		{"definir a = 5 * 5\n a\n", 25},
		{"definir a = 5\n definir b = a\n b\n", 5},
		{"definir a = 5\n definir b = a\n definir c = a + b + 5\n c\n", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "funcion(x) { x + 2 }"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"

	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"definir identity = funcion(x) { x\n }\n identity(5)\n", 5},
		{"definir identity = funcion(x) { devolver x\n }\n identity(5)\n", 5},
		{"definir double = funcion(x) { x * 2\n }\n double(5)\n", 10},
		{"definir add = funcion(x, y) { x + y\n }\n add(5, 5)\n", 10},
		{"definir add = funcion(x, y) { x + y\n }\n add(5 + 5, add(5, 5))\n", 20},
		{"funcion(x) { x\n }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestClosures(t *testing.T) {
	input := `
definir newAdder = funcion(x) {
	funcion(y) { x + y }
}

definir addTwo = newAdder(2)
addTwo(2)`

	testIntegerObject(t, testEval(input), 4)
}

func TestStringLiteral(t *testing.T) {
	input := `"hello world"`

	evaluated := testEval(input)

	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "hello world" {
		t.Errorf("str.Value not %q. got=%q", "hello world", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"hello" + " " + "world"`

	evaluated := testEval(input)

	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "hello world" {
		t.Errorf("str.Value not %q. got=%q", "hello world", str.Value)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
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

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}

	return true
}

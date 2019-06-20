package object

// Package object implements the object system (or value system) of Monkey
// used to both represent values as the evaluator encounters and constructs
// them as well as how the user interacts with values.

import (
	"byaigo/ast"
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	// INTEGER_OBJ is the Integer object type
	INTEGER_OBJ = "INTEGER"

	// STRING_OBJ is the String object type
	STRING_OBJ = "STRING"

	// BOOLEAN_OBJ is the Boolean object type
	BOOLEAN_OBJ = "BOOLEAN"

	// NULL_OBJ is the Null object type
	NULL_OBJ = "NULL"

	// RETURN_VALUE_OBJ is the Return Value object type
	RETURN_VALUE_OBJ = "RETURN_VALUE"

	// ERROR_OBJ is the error object
	ERROR_OBJ = "ERROR"

	// FUNCTION_OBJ is the function object type
	FUNCTION_OBJ = "FUNCTION"

	// BUILTIN_OBJ is the builtin function object type
	BUILTIN_OBJ = "BUILTIN"

	// ARRAY_OBJ is the list object type (keeping it named array for continuity with book)
	ARRAY_OBJ = "ARRAY"

	// HASH_OBJ is the hash object type
	HASH_OBJ = "HASH"

	// QUOTE_OBJ is the quote object type
	QUOTE_OBJ = "QUOTE"
)

type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() ObjectType { return QUOTE_OBJ }
func (q *Quote) Inspect() string  { return "QUOTE(" + q.Node.String() + ")" }

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

type Hashable interface {
	HashKey() HashKey
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}

// Error is the base struct of our Error type
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Object represents a value and implementations are expected to implement
// `Type()` and `Inspect()` functions
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer is the integer type used to represent integer literals and holds
// an internal int64 value
type Integer struct {
	Value int64
}

// Inspect returns a stringified version of the object for debugging
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type returns the type of the object
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// String is the string type used to represent string literals and holds
// an internal string value
type String struct {
	Value string
}

// Type returns the type of the object
func (s *String) Type() ObjectType { return STRING_OBJ }

// Inspect returns a stringified version of the object for debugging
func (s *String) Inspect() string { return s.Value }

// Boolean is the boolean type and used to represent boolean literals and
// holds an interval bool value
type Boolean struct {
	Value bool
}

// Inspect returns a stringified version of the object for debugging
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type returns the type of the object
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Null is the null type and used to represent the absence of a value
type Null struct{}

// Inspect returns a stringified version of the object for debugging
func (n *Null) Inspect() string { return "null" }

// Type returns the type of the object
func (n *Null) Type() ObjectType { return NULL_OBJ }

// ReturnValue is the return type and used to hold the value of another object.
// This is used for `return` statements and this object is tracked through
// the evalulator and when encountered stops evaluation of the program,
// or body of a function.
type ReturnValue struct {
	Value Object
}

// Type returns the type of the object
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect returns a stringified version of the object for debugging
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
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

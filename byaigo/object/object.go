package object

// Package object implements the object system (or value system) of Monkey
// used to both represent values as the evaluator encounters and constructs
// them as well as how the user interacts with values.

import "fmt"

// ObjectType represents the type of an object
type ObjectType string

const (
	// INTEGER_OBJ is the Integer object type
	INTEGER_OBJ = "INTEGER"

	// BOOLEAN_OBJ is the Boolean object type
	BOOLEAN_OBJ = "BOOLEAN"

	// NULL_OBJ is the Null object type
	NULL_OBJ = "NULL"

	// RETURN_VALUE_OBJ is the Return Value object type
	RETURN_VALUE_OBJ = "RETURN_VALUE"

	// ERROR_OBJ is the error object
	ERROR_OBJ = "ERROR"
)

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

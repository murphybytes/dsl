package ast

import "fmt"

type ErrType int

const (
	TypeMismatch ErrType = iota
	UnsupportedType
	UnsupportedOperator
	SyntaxError
	MissingKey
	IndexOutOfRange

)

type Error interface {
	error
	Type() ErrType
}

type ErrAst struct {
	typ ErrType
	msg string
}

func (ea ErrAst) Error() string {
	return ea.msg
}

func (ea ErrAst) Type() ErrType {
	return ea.typ
}

// UnsupportedTypeError is raised when data passed in the context doesn't map
// to a data type we support
func UnsupportedTypeError(t interface{}) error {
	return &ErrAst{
		msg: fmt.Sprintf("unsupported type %T", t),
		typ: UnsupportedType,
	}
}

// MissingKeyError is raised when the data passed in as context doesn't map to an expression variable.
func MissingKeyError(key string) error {
	return &ErrAst{
		msg: fmt.Sprintf("variable segment %q does not map to context data", key),
		typ: MissingKey,
	}
}

// IndexOutOfRangeError is raised when the data passed in as context doesn't map to an expression variable.
func IndexOutOfRangeError(index string) error {
	return &ErrAst{
		msg: fmt.Sprintf("variable segment %q does not map to context data", index),
		typ: IndexOutOfRange,
	}
}

// TypeMismatchError indicates types of the arguments don't match when they should.
func TypeMismatchError(l, r *Value) error {
	return &ErrAst{
		msg: fmt.Sprintf("type mismatch lval %#v rval %#v", l, r),
		typ: TypeMismatch,
	}
}

// NewUnsupportedOperatorError indicates that an unsupported operator is being used.
func NewUnsupportedOperatorError(s string) error {
	return &ErrAst{
		msg: fmt.Sprintf("unsupported operator %q", s),
		typ: UnsupportedOperator,
	}
}

func NewSyntaxError(formt string, v ...interface{}) error {
	return &ErrAst{
		msg: fmt.Sprintf("syntax error: %s", fmt.Sprintf(formt, v...)),
		typ: SyntaxError,
	}
}

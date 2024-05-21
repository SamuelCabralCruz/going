//go:build test

package roar

import "errors"

type someError struct {
	Roar[someError]
}

type anotherError struct {
	Roar[anotherError]
}

var standardError = errors.New("standard error")
var roarError = New[any]("roar error")
var someErrorBase = someError{New[someError]("base")}
var someErrorWithStandardErrorCause = someError{New[someError](
	"with standard error cause",
	WithCause(standardError))}
var someErrorWithMultipleCauses = someError{New[someError](
	"with multiple causes",
	WithCause(errors.New("first cause")),
	WithCause(errors.New("second cause")),
	WithCause(errors.New("last cause")))}
var someErrorWithField = someError{New[someError](
	"with field",
	WithField("name", "value"))}
var someErrorWithFields = someError{New[someError](
	"with fields",
	WithField("name1", "value1"),
	WithField("name2", "value2"))}
var someErrorWithCauseAndFields = someError{New[someError](
	"with cause and fields",
	WithCause(standardError),
	WithField("name1", "value1"),
	WithField("name2", "value2"))}
var someErrorWithNestedCause = someError{New[someError](
	"with nested cause",
	WithCause(anotherError{New[anotherError](
		"nested",
		WithCause(standardError))}))}
var someErrorWithEverything = someError{New[someError](
	"with everything",
	WithCause(anotherError{New[anotherError](
		"nested",
		WithCause(standardError),
		WithField("name1", "value1"),
		WithField("name2", "value2")),
	}),
	WithField("name3", "value3"),
	WithField("name4", "value4"),
)}

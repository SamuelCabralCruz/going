package xpctd

import (
	"fmt"
)

func newReporter[T any]() reporter[T] {
	return reporter[T]{
		polarity: positive,
		expected: func(actual T) string {
			return fmt.Sprintf(`"%+v"`, actual)
		},
		to: func(_ T) string {
			return "match"
		},
		but: func(_ T) string {
			return ""
		},
	}
}

type messagePartReporter[T any] func(actual T) string

type reporter[T any] struct {
	polarity polarity
	expected messagePartReporter[T]
	to       messagePartReporter[T]
	but      messagePartReporter[T]
}

var _ Reporter[any] = reporter[any]{}
var _ Expected[any] = reporter[any]{}
var _ To[any] = reporter[any]{}
var _ But[any] = reporter[any]{}

func (r reporter[T]) Positive() Reporter[T] {
	r.polarity = positive
	return r
}

func (r reporter[T]) Negative() Reporter[T] {
	r.polarity = negative
	return r
}

func (r reporter[T]) Report(actual T) string {
	message := fmt.Sprintf("Expected %s %s %s",
		r.expected(actual),
		r.polarity.String(),
		r.to(actual),
	)
	but := r.but(actual)
	if len(but) != 0 {
		message = fmt.Sprintf("%s, but %s", message, but)
	}
	return message
}

func (r reporter[T]) Error(actual T) error {
	return newExpectationError(r.Report(actual))
}

func (r reporter[T]) Validation(actual T) (bool, error) {
	return false, r.Error(actual)
}

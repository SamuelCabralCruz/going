package roar

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/samber/lo"
	"strings"
)

type Option func(*parameters)

func WithCause(cause error) Option {
	return func(p *parameters) {
		p.cause = cause
	}
}

type field lo.Entry[string, any]

func (f field) toString() string {
	return fmt.Sprintf("%s: %+v", f.Key, f.Value)
}

func WithField(key string, value any) Option {
	return func(p *parameters) {
		p.fields = append(p.fields, field{
			Key:   key,
			Value: value,
		})
	}
}

type parameters struct {
	cause  error
	fields []field
}

type Roar[T any] struct {
	message string
	cause   error
	fields  []field
}

var _ error = Roar[struct{}]{}

func New[T any](message string, options ...Option) Roar[T] {
	params := lo.Reduce(options, func(agg *parameters, opt Option, _ int) *parameters {
		opt(agg)
		return agg
	}, &parameters{})
	return Roar[T]{
		message: fmt.Sprintf("%s: %s", phi.Type[T](), message),
		cause:   params.cause,
		fields:  params.fields,
	}
}

func (r Roar[T]) Error() string {
	parts := []string{fmt.Sprintf("%s: %s", phi.Type[T](), r.message)}
	if len(r.fields) > 0 {
		fields := lo.Map(r.fields, func(f field, _ int) string { return f.toString() })
		parts = append(parts, fmt.Sprintf("[%s]", strings.Join(fields, ", ")))
	}
	if r.cause != nil {
		parts = append(parts, fmt.Sprintf("- %s", r.cause.Error()))
	}
	return strings.Join(parts, " ")
}

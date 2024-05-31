package reporter

import "fmt"

func (r reporter[T]) To(describe func(actual T) string) To[T] {
	r.to = describe
	return r
}

func (r reporter[T]) ToFormatted(format string, a ...any) To[T] {
	return r.To(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

func (r reporter[T]) ToBe(describe func(actual T) string) To[T] {
	return r.To(func(actual T) string {
		return fmt.Sprintf("be %s", describe(actual))
	})
}

func (r reporter[T]) ToBeFormatted(format string, a ...any) To[T] {
	return r.ToBe(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

func (r reporter[T]) ToHave(describe func(actual T) string) To[T] {
	return r.To(func(actual T) string {
		return fmt.Sprintf("have %s", describe(actual))
	})
}

func (r reporter[T]) ToHaveFormatted(format string, a ...any) To[T] {
	return r.ToHave(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

package xpctd

import "fmt"

func (r reporter[T]) But(describe func(actual T) string) But[T] {
	r.but = describe
	return r
}

func (r reporter[T]) ButFormatted(format string, a ...any) But[T] {
	return r.But(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

func (r reporter[T]) ButReceived(describe func(actual T) string) But[T] {
	return r.But(func(actual T) string {
		return fmt.Sprintf("received %s", describe(actual))
	})
}

func (r reporter[T]) ButReceivedFormatted(format string, a ...any) But[T] {
	return r.ButReceived(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

func (r reporter[T]) ButWas(describe func(actual T) string) But[T] {
	return r.But(func(actual T) string {
		return fmt.Sprintf("was %s", describe(actual))
	})
}

func (r reporter[T]) ButWasFormatted(format string, a ...any) But[T] {
	return r.ButWas(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

func (r reporter[T]) ButWasA(description string) But[T] {
	return r.ButWasFormatted("a %s", description)
}

func (r reporter[T]) ButWasOfType() But[T] {
	return r.ButWas(func(actual T) string {
		return fmt.Sprintf(`of type "%T"`, actual)
	})
}

func (r reporter[T]) ButHad(describe func(actual T) string) But[T] {
	return r.But(func(actual T) string {
		return fmt.Sprintf("had %s", describe(actual))
	})
}

func (r reporter[T]) ButHadFormatted(format string, a ...any) But[T] {
	return r.ButHad(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

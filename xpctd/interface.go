package xpctd

type Reporter[T any] interface {
	Positive() Reporter[T]
	Negative() Reporter[T]
	Report(T) string
	Error(T) error
	Validation(T) (bool, error)
}

type Expected[T any] interface {
	Reporter[T]
	To(describe func(actual T) string) To[T]
	ToFormatted(format string, a ...any) To[T]
	ToBe(describe func(actual T) string) To[T]
	ToBeFormatted(format string, a ...any) To[T]
	ToBeA(description string) To[T]
	ToBeOfType(value any) To[T]
	ToHave(describe func(actual T) string) To[T]
	ToHaveFormatted(format string, a ...any) To[T]
}

type To[T any] interface {
	Reporter[T]
	But(describe func(actual T) string) But[T]
	ButFormatted(format string, a ...any) But[T]
	ButReceived(describe func(actual T) string) But[T]
	ButReceivedFormatted(format string, a ...any) But[T]
	ButWas(describe func(actual T) string) But[T]
	ButWasFormatted(format string, a ...any) But[T]
	ButWasA(description string) But[T]
	ButWasOfType() But[T]
	ButHad(describe func(actual T) string) But[T]
	ButHadFormatted(format string, a ...any) But[T]
}

type But[T any] interface {
	Reporter[T]
}

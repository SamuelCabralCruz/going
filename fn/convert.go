package fn

func ToProducer[T any](supply Supplier[T]) Producer[T] {
	return func() (T, error) {
		return Safe(supply)
	}
}

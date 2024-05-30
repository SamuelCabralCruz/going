package tuple

func Swap[T any, U any](left T, right U) (U, T) {
	return right, left
}

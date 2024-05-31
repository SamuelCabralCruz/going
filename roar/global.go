package roar

func Aggregate(errs ...error) error {
	if len(errs) == 1 {
		return errs[0]
	}
	return NewAggregatedError(errs...)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

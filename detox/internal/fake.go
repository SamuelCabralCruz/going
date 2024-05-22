package internal

func NewFake[T any](mockName string, methodName string) *Fake[T] {
	return &Fake[T]{
		mockName:   mockName,
		methodName: methodName,
	}
}

type Fake[T any] struct {
	mockName       string
	methodName     string
	persistentImpl *T
	ephemeralImpl  []T
}

func (f *Fake[T]) RegisterImplementation(impl T) {
	f.persistentImpl = &impl
}

func (f *Fake[T]) RegisterImplementationOnce(impl T) {
	f.ephemeralImpl = append(f.ephemeralImpl, impl)
}

func (f *Fake[T]) InvokeFakeImplementation() T {
	if len(f.ephemeralImpl) == 0 && f.persistentImpl == nil {
		panic(newMissingFakeImplementationError(f.mockName, f.methodName))
	}
	if len(f.ephemeralImpl) > 0 {
		impl := f.ephemeralImpl[0]
		f.ephemeralImpl = f.ephemeralImpl[1:]
		return impl
	}
	return *f.persistentImpl
}

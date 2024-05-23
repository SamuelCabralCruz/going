package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal"
	"github.com/samber/lo"
)

func (m *mocked[T]) Assert() Asserter {
	return &asserter[T]{
		m,
	}
}

type Assertable interface {
	Name() string
	Assert() Asserter
}

type Asserter interface {
	HasBeenCalled() bool
	HasBeenCalledOnce() bool
	HasBeenCalledTimes(int) bool
	HasBeenCalledWith(...any) bool
	HasBeenCalledOnceWith(...any) bool
	HasBeenCalledTimesWith(int, ...any) bool
	HasCalls(...internal.Call) bool
	HasNthCall(int, internal.Call) bool
	HasOrderedCalls(...internal.Call) bool
}

type asserter[T any] struct {
	*mocked[T]
}

// HasBeenCalled Assert that mock has been called at least once
func (a *asserter[T]) HasBeenCalled() bool {
	return a.CallsCount() > 0
}

// HasBeenCalledTimes Assert that mock has been called exact number of times
func (a *asserter[T]) HasBeenCalledTimes(times int) bool {
	return a.CallsCount() == times
}

// HasBeenCalledOnce Assert that mock has been called exactly once
func (a *asserter[T]) HasBeenCalledOnce() bool {
	return a.HasBeenCalledTimes(1)
}

// HasBeenCalledWith Assert that mock has been called at least once with specified arguments
func (a *asserter[T]) HasBeenCalledWith(args ...any) bool {
	expected := internal.NewCall(args...)
	filtered := lo.Filter(a.Calls(), func(observed internal.Call, _ int) bool {
		return observed.EqualTo(expected)
	})
	return len(filtered) > 0
}

// HasBeenCalledOnceWith Assert that mock has been called exactly once with specified arguments
func (a *asserter[T]) HasBeenCalledOnceWith(args ...any) bool {
	expected := internal.NewCall(args...)
	filtered := lo.Filter(a.Calls(), func(observed internal.Call, _ int) bool {
		return observed.EqualTo(expected)
	})
	return len(filtered) == 1
}

// HasBeenCalledTimesWith Assert that mock has been called with specified arguments exactly number of times
func (a *asserter[T]) HasBeenCalledTimesWith(times int, args ...any) bool {
	expected := internal.NewCall(args...)
	filtered := lo.Filter(a.Calls(), func(observed internal.Call, _ int) bool {
		return observed.EqualTo(expected)
	})
	return len(filtered) == times
}

// HasCalls Assert that mock has specified calls
func (a *asserter[T]) HasCalls(calls ...internal.Call) bool {
	candidates := a.Calls()
	for _, expected := range calls {
		_, index, ok := lo.FindIndexOf(candidates, func(observed internal.Call) bool {
			return observed.EqualTo(expected)
		})
		if ok == false {
			return false
		}
		candidates = append(candidates[:index], candidates[index+1:]...)
	}
	return true
}

// HasNthCall Assert that mock has specified nth call
func (a *asserter[T]) HasNthCall(index int, call internal.Call) bool {
	return a.NthCall(index).EqualTo(call)
}

// HasOrderedCalls Assert that mock has specified calls
func (a *asserter[T]) HasOrderedCalls(calls ...internal.Call) bool {
	if a.CallsCount() != len(calls) {
		return false
	}
	for index, expected := range calls {
		if a.HasNthCall(index, expected) == false {
			return false
		}
	}
	return true
}

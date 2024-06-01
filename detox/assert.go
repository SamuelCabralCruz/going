package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/samber/lo"
)

func (m *mocked[T, U]) Assert() Asserter {
	return &asserter[T, U]{
		m,
	}
}

type Assertable interface {
	Describe() string
	Calls() []common.Call
	Assert() Asserter
}

type Call []any

type Asserter interface {
	HasBeenCalled() bool
	HasBeenCalledOnce() bool
	HasBeenCalledTimes(int) bool
	HasBeenCalledWith(...any) bool
	HasBeenCalledOnceWith(...any) bool
	HasBeenCalledTimesWith(int, ...any) bool
	HasCalls(...Call) bool
	HasNthCall(int, Call) bool
	HasOrderedCalls(...Call) bool
}

type asserter[T any, U any] struct {
	*mocked[T, U]
}

// HasBeenCalled Assert that mock has been called at least once
func (a *asserter[T, U]) HasBeenCalled() bool {
	return a.CallsCount() > 0
}

// HasBeenCalledTimes Assert that mock has been called exact number of times
func (a *asserter[T, U]) HasBeenCalledTimes(times int) bool {
	return a.CallsCount() == times
}

// HasBeenCalledOnce Assert that mock has been called exactly once
func (a *asserter[T, U]) HasBeenCalledOnce() bool {
	return a.HasBeenCalledTimes(1)
}

// HasBeenCalledWith Assert that mock has been called at least once with specified arguments
func (a *asserter[T, U]) HasBeenCalledWith(args ...any) bool {
	expected := common.NewCall(args...)
	filtered := lo.Filter(a.Calls(), func(observed common.Call, _ int) bool {
		return observed.EqualTo(expected)
	})
	return len(filtered) > 0
}

// HasBeenCalledOnceWith Assert that mock has been called exactly once with specified arguments
func (a *asserter[T, U]) HasBeenCalledOnceWith(args ...any) bool {
	expected := common.NewCall(args...)
	filtered := lo.Filter(a.Calls(), func(observed common.Call, _ int) bool {
		return observed.EqualTo(expected)
	})
	return len(filtered) == 1
}

// HasBeenCalledTimesWith Assert that mock has been called with specified arguments exactly number of times
func (a *asserter[T, U]) HasBeenCalledTimesWith(times int, args ...any) bool {
	expected := common.NewCall(args...)
	filtered := lo.Filter(a.Calls(), func(observed common.Call, _ int) bool {
		return observed.EqualTo(expected)
	})
	return len(filtered) == times
}

// HasCalls Assert that mock has specified calls
func (a *asserter[T, U]) HasCalls(calls ...Call) bool {
	candidates := a.Calls()
	for _, expected := range calls {
		_, index, ok := lo.FindIndexOf(candidates, func(observed common.Call) bool {
			return observed.EqualTo(common.NewCall(expected...))
		})
		if ok == false {
			return false
		}
		candidates = append(candidates[:index], candidates[index+1:]...)
	}
	return true
}

// HasNthCall Assert that mock has specified nth call
func (a *asserter[T, U]) HasNthCall(index int, call Call) bool {
	return a.NthCall(index).EqualTo(common.NewCall(call...))
}

// HasOrderedCalls Assert that mock has exactly specified calls in specified order
func (a *asserter[T, U]) HasOrderedCalls(calls ...Call) bool {
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

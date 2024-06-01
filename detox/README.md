# detox

## Caveats

- Cannot pass argument matchers when asserting calls
- No actual stubs due to golang's type system limitations

[//]: # (TODO: matchers)
- HasBeenCalledWith(...any) bool
- HasBeenCalledOnceWith(...any) bool
- HasBeenCalledTimesWith(int, ...any) bool
- HasCalls(...Call) bool
- HasNthCall(int, Call) bool
- HasOrderedCalls(...Call) bool

[//]: # (TODO: test assertion in usage)
[//]: # (TODO: add documentation for gomega matchers)

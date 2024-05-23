package common

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/phi"
	"runtime"
)

func NewMockInfo[T any]() MockInfo {
	_, file, line, _ := runtime.Caller(3)
	return MockInfo{
		interfaceName: phi.BaseTypeName[T](),
		fileReference: fmt.Sprintf("%s [%d]", file, line),
	}
}

type MockInfo struct {
	interfaceName string
	fileReference string
}

func (i MockInfo) Interface() string {
	return i.interfaceName
}

func (i MockInfo) Reference() string {
	return i.fileReference
}

func (i MockInfo) Describe() string {
	return fmt.Sprintf("%s (%s)", i.interfaceName, i.fileReference)
}

func NewMockedMethodInfo[T any](mockInfo MockInfo, method T) MockedMethodInfo {
	return MockedMethodInfo{
		MockInfo:   mockInfo,
		methodName: phi.FunctionName(method),
	}
}

type MockedMethodInfo struct {
	MockInfo
	methodName string
}

func (i MockedMethodInfo) Method() string {
	return i.methodName
}

func (i MockedMethodInfo) Describe() string {
	return fmt.Sprintf("%s.%s (%s)", i.interfaceName, i.methodName, i.fileReference)
}

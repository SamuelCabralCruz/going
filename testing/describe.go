package testing

import (
	"github.com/SamuelCabralCruz/went/phi"
	"runtime"
	"strings"

	"github.com/onsi/ginkgo/v2"
)

func DescribeFunction(function any, args ...any) bool {
	return ginkgo.Describe(phi.FunctionName(function), args...)
}

func DescribeType[T any](args ...any) bool {
	return ginkgo.Describe(phi.BaseTypeName[T](), args...)
}

func getPackageNameForCaller(caller int) string {
	packageName, _ := splitCallerPath(caller)
	return packageName
}

func splitCallerPath(caller int) (packageName string, functionName string) {
	pc, _, _, _ := runtime.Caller(caller)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	pkgName := ""
	funcName := parts[pl-1]
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		pkgName = strings.Join(parts[0:pl-2], ".")
	} else {
		pkgName = strings.Join(parts[0:pl-1], ".")
	}
	return pkgName, funcName
}

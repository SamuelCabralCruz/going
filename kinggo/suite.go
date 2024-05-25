package kinggo

import (
	"fmt"
	"runtime"
	"strings"
)

func CreateTestSuiteName() string {
	return fmt.Sprint(getPackageNameForCaller(3))
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

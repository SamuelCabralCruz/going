package phi

import (
	"regexp"
	"runtime"
	"strings"
)

func FunctionFullPath(value any) string {
	return runtime.FuncForPC(Value(value).Pointer()).Name()
}

func FunctionName(function any) string {
	functionFullPath := strings.ReplaceAll(FunctionFullPath(function), "[...]", "")
	var functionNameSubPathIndex int
	anonymousFuncRegex := regexp.MustCompile(`^.*\.(func\d+(\.\d+)?)$`)
	if anonymousFuncRegex.MatchString(functionFullPath) {
		functionNameSubPathIndex = anonymousFuncRegex.FindStringSubmatchIndex(functionFullPath)[2]
	} else {
		functionNameSubPathIndex = strings.LastIndex(functionFullPath, ".") + 1
	}
	return strings.TrimRight(functionFullPath[functionNameSubPathIndex:], "-fm")

}

func TypeName[T any]() string {
	return Type[T]().Name()
}

func BaseTypeName[T any]() string {
	return regexp.MustCompile("\\[.*]").ReplaceAllString(TypeName[T](), "")
}

package testing

import "fmt"

func CreateTestSuiteName() string {
	return fmt.Sprint(getPackageNameForCaller(3))
}

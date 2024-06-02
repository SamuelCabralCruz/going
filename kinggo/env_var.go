package kinggo

import (
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/samber/lo"
	"os"
)

func SetUpEnvironmentVariable(key string, value string) {
	if err := os.Setenv(key, value); err != nil {
		panic(fmt.Errorf(`could not set environment variable "%s" with value "%s"`, key, value))
	}
	ginkgo.DeferCleanup(func() {
		if err := os.Unsetenv(key); err != nil {
			panic(fmt.Errorf(`could not unset environment variable "%s"`, key))
		}
	})
}

func SetUpEnvironmentVariables(envVars map[string]string) {
	lo.ForEach(lo.Entries(envVars), func(item lo.Entry[string, string], _ int) {
		SetUpEnvironmentVariable(item.Key, item.Value)
	})
}

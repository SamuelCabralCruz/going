package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/samber/lo"
	"strings"
)

func describeCalls(actual detox.Assertable) []string {
	return lo.Map(actual.Calls(), func(call common.Call, index int) string {
		return fmt.Sprintf("[%d]: %s", index, describeArgs(call.Args()...))
	})
}

func describeArgs(args ...any) string {
	return fmt.Sprintf("[%s]", strings.Join(lo.Map(args, func(arg any, _ int) string {
		return fmt.Sprintf("<%T> %+v", arg, arg)
	}), ", "))
}

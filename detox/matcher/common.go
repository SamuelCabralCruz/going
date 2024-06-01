package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/samber/lo"
	"strings"
)

func describeCalls(calls []common.Call) []string {
	return lo.Map(calls, describeCall)
}

func describeCall(call common.Call, index int) string {
	return fmt.Sprintf("[%d]: %s", index, describeArgs(call.Args()))
}

func describeArgs(args []any) string {
	return fmt.Sprintf("[%s]", strings.Join(lo.Map(args, func(arg any, _ int) string {
		return fmt.Sprintf("<%T> %+v", arg, arg)
	}), ", "))
}

func toCommonCalls(calls []detox.Call) []common.Call {
	return lo.Map(calls, func(call detox.Call, _ int) common.Call {
		return toCommonCall(call)
	})
}

func toCommonCall(call detox.Call) common.Call {
	return common.NewCall(call...)
}

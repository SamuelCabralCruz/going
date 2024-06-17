package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/going/detox"
	"github.com/SamuelCabralCruz/going/detox/internal/common"
	"github.com/SamuelCabralCruz/going/thong"
	"github.com/samber/lo"
	"strings"
)

func describeCalls(calls []common.Call) string {
	if len(calls) == 0 {
		return "None"
	}
	return thong.IndentParts("\t\t", lo.Map(calls, describeCall))
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

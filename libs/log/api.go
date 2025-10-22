package log

import (
	"fmt"
	"log/slog"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func logMessage(level string, format string, args ...any) {
	var message string
	if format != "" {
		message = fmt.Sprintf(format, args...)
	} else {
		message = fmt.Sprint(args...)
	}

	switch strings.ToLower(level) {
	case "debug":
		slog.Debug(message)
	case "info":
		slog.Info(message)
	case "warn":
		slog.Warn(message)
	case "error":
		slog.Error(message)
	}
}

func logFunc(level string) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		var v []any
		for i := 1; i <= L.GetTop(); i++ {
			v = append(v, L.Get(i).String())
		}
		logMessage(level, "", v...)
		return 0
	}
}

func logFFunc(level string) func(L *lua.LState) int {
	return func(L *lua.LState) int {
		format := L.CheckString(1)
		top := L.GetTop()
		args := make([]any, 0, top-1)
		for i := 2; i <= top; i++ {
			val := L.CheckAny(i)
			if ud, ok := val.(*lua.LUserData); ok {
				L.Push(ud)
				if L.CallByParam(lua.P{
					Fn:      L.GetGlobal("tostring"),
					NRet:    1,
					Protect: true,
				}, ud) == nil {
					str := L.Get(-1).String()
					L.Pop(1)
					args = append(args, str)
				}
			} else {
				args = append(args, val.String())
			}
		}
		logMessage(level, format, args...)
		return 0
	}
}

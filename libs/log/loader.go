package log

import lua "github.com/yuin/gopher-lua"

// Preload adds log to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//	local log = require("log")
func Preload(L *lua.LState) {
	L.PreloadModule("log", Loader)
}

func Loader(L *lua.LState) int {
	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"debug":  logFunc("debug"),
	"info":   logFunc("info"),
	"warn":   logFunc("warn"),
	"error":  logFunc("error"),
	"debugf": logFFunc("debug"),
	"infof":  logFFunc("info"),
	"warnf":  logFFunc("warn"),
	"errorf": logFFunc("error"),
}

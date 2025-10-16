package golua

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"

	"github.com/wildmap/golua/libs"
)

type Option struct {
	globals map[string]any
	params  map[string]any
}

func WithGlobals(globals map[string]any) Option {
	return Option{
		globals: globals,
	}
}

func WithParams(params map[string]any) Option {
	return Option{
		params: params,
	}
}

func NewLuaState(globals map[string]any) *lua.LState {
	state := lua.NewState(lua.Options{IncludeGoStackTrace: true})
	libs.Preload(state)
	state.SetGlobal("I64", state.NewFunction(I64))

	for k, v := range globals {
		state.SetGlobal(k, luar.New(state, v))
	}

	return state
}

// DoFile 执行lua文件
func DoFile(fpath, fname string, opts ...Option) (outline string, err error) {
	content, err := readScript(fpath, fname)
	if err != nil {
		return "", errors.New("read script file")
	}

	return DoString(content, opts...)
}

// DoString 执行lua字符串
func DoString(script string, opts ...Option) (outline string, err error) {
	// 处理选项参数
	var globals map[string]any
	var params map[string]any
	for _, opt := range opts {
		if opt.globals != nil {
			if globals == nil {
				globals = make(map[string]any)
			}
			for k, v := range opt.globals {
				globals[k] = v
			}
		}
		if opt.params != nil {
			if params == nil {
				params = make(map[string]any)
			}
			for k, v := range opt.params {
				params[k] = v
			}
		}
	}

	L := NewLuaState(globals)
	defer L.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	L.SetContext(ctx)

	// 添加、删除lua脚本参数
	L.SetGlobal("Params", luar.New(L, params))
	// setPrintf
	setPrintf(L, &outline)

	beginTime := time.Now()
	slog.Info("[LuaScript] running lua script start")
	defer func() {
		slog.Info(fmt.Sprintf("[LuaScript] running lua script end, cost=%vus", time.Now().Sub(beginTime).Microseconds()))
	}()

	if err = L.DoString(script); err != nil {
		return "", err
	}
	return
}

func readScript(fpath, fname string) (string, error) {
	scriptPath := path.Join(fpath, fname)
	content, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", err
	}

	return string(content), err
}

func setPrintf(l *lua.LState, outline *string) {
	_printf := func(format string, args ...any) {
		*outline += fmt.Sprintf(format, args...)
	}
	_print := func(args ...any) {
		*outline += fmt.Sprint(args...)
	}
	_println := func(args ...any) {
		*outline += fmt.Sprintln(args...)
	}
	l.SetGlobal("printf", luar.New(l, _printf))
	l.SetGlobal("print", luar.New(l, _print))
	l.SetGlobal("println", luar.New(l, _println))
}

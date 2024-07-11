package code_engine

import (
	"errors"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/issueye/code_magic/backend/pkg/utils"
	"go.uber.org/zap"
)

type JsVM struct {
	// vm 虚拟机
	vm *goja.Runtime
	// 注册
	registry *require.Registry
	// 事件循环
	loop *EventLoop
	// 全局goja加载目录
	globalPath string
	// 日志对象
	log *zap.Logger
	// 输出回调
	ConsoleCallBack ConsoleCallBack
}

type Code struct {
	Path    string
	Program *goja.Program
}

type ModuleFunc = func(vm *goja.Runtime, module *goja.Object)

func NewJsVM(globalPath string, log *zap.Logger, consoleCallBack ConsoleCallBack) *JsVM {
	jsVM := &JsVM{
		vm:              goja.New(),
		globalPath:      globalPath,
		ConsoleCallBack: consoleCallBack,
	}

	// 输出日志
	console := newConsole(log, consoleCallBack)
	o := jsVM.vm.NewObject()
	o.Set("log", console.Log)
	o.Set("debug", console.Debug)
	o.Set("info", console.Info)
	o.Set("error", console.Error)
	o.Set("warn", console.Warn)
	jsVM.vm.Set("console", o)

	ops := []require.Option{}

	if globalPath != "" {
		ops = append(ops, require.WithGlobalFolders(globalPath))
	}

	// source
	ops = append(ops, require.WithLoader(jsVM.sourceLoader(globalPath)))

	jsVM.loop = NewEventLoop(jsVM.vm)
	jsVM.registry = require.NewRegistry(ops...)

	self := jsVM.vm.GlobalObject()
	jsVM.vm.Set("self", self)

	jsVM.vm.Set("atob", func(code string) string {
		raw, err := utils.Base64DecodeString(code)
		if err != nil {
			panic(err)
		}
		return raw
	})

	jsVM.vm.Set("btoa", func(code string) string {
		return utils.Base64EncodeString(code)
	})

	return jsVM
}

func (jv *JsVM) Run(name string, code *Code) error {
	if code.Program != nil {
		var loop *EventLoop
		loop = NewEventLoop(jv.vm)

		var exception error
		loop.Run(func(r *goja.Runtime) {
			_, err := r.RunProgram(code.Program)
			if gojaErr, ok := err.(*goja.Exception); ok {
				exception = errors.New(gojaErr.String())
				return
			}
		})

		if exception != nil {
			return exception
		}
	} else {
		// jv.
	}

	return nil
}
